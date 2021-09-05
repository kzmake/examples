import json
from typing import List

from dapr.clients import DaprClient
from dapr.clients.grpc._state import StateItem

from .domain import AggregateRootUser, UserID


class UserRepository:
    def add(self, user: AggregateRootUser) -> None:
        keys = self.__keys()
        keys.append(user.user_id.value)
        print(f"keys: {keys}")
        with DaprClient() as d:
            keys = StateItem(key="__keys__", value=json.dumps(keys))
            user_data = StateItem(key=user.user_id.value, value=user.to_json())
            d.save_bulk_state(store_name="state-user", states=[keys, user_data])

    def list(self) -> List[AggregateRootUser]:
        keys = self.__keys()
        users = self.__find(keys)

        return users

    def get(self, user_id: UserID) -> AggregateRootUser:
        users = self.__find([str(user_id.value)])
        assert len(users) == 1

        return users[0]

    def remove(self, user: AggregateRootUser) -> None:
        keys = self.__keys()

        new_keys = [x for x in keys if x != user.user_id.value]
        with DaprClient() as d:
            d.save_state(
                store_name="state-user", key="__keys__", value=json.dumps(new_keys)
            )
            d.delete_state(store_name="state-user", key=user.user_id.value)

    def update(self, user: AggregateRootUser) -> None:
        with DaprClient() as d:
            d.save_state(
                store_name="state-user",
                key=user.user_id.value,
                value=user.to_json(),
            )

    def __keys(self) -> List[str]:
        with DaprClient() as d:
            data = d.get_state(store_name="state-user", key="__keys__").data
            print(f"keys: {data}")
            if not data:
                return []

            return json.loads(data)

    def __find(self, keys: List[str]) -> List[AggregateRootUser]:
        if not keys:
            return []

        with DaprClient() as d:
            items = d.get_bulk_state(store_name="state-user", keys=keys).items
            print(f"find: {[i.data for i in items]}")

            users = [AggregateRootUser.from_json(i.data) for i in items]
            print(f"find user: {users}")

            return users
