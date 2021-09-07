import basic.user.v1.user_pb2 as userpb
import ulid
from basic.user.v1.user_grpc import UserServiceBase
from grpclib.server import Stream

from .domain import User, UserID, UserName
from .repository import UserRepository

user_repository = UserRepository()


class UserService(UserServiceBase):
    async def Create(
        self, stream: Stream[userpb.CreateRequest, userpb.CreateResponse]
    ) -> None:
        req = await stream.recv_message()
        assert req is not None

        user = User(user_id=UserID(str(ulid.new())), name=UserName(req.name))
        user_repository.add(user)

        u = userpb.User(id=user.user_id.value, name=user.name.value)
        res = userpb.CreateResponse(user=u)
        await stream.send_message(res)

    async def List(
        self, stream: Stream[userpb.ListRequest, userpb.ListResponse]
    ) -> None:
        req = await stream.recv_message()
        assert req is not None

        users = user_repository.list()

        us = [
            userpb.User(id=user.user_id.value, name=user.name.value) for user in users
        ]
        res = userpb.ListResponse(users=us)
        await stream.send_message(res)

    async def Get(self, stream: Stream[userpb.GetRequest, userpb.GetResponse]) -> None:
        req = await stream.recv_message()
        assert req is not None

        user_id = UserID(req.user_id)
        user = user_repository.get(user_id)

        u = userpb.User(id=user.user_id.value, name=user.name.value)
        res = userpb.GetResponse(user=u)
        await stream.send_message(res)

    async def Rename(
        self, stream: Stream[userpb.RenameRequest, userpb.RenameResponse]
    ) -> None:
        req = await stream.recv_message()
        assert req is not None

        user_id = UserID(req.user_id)
        user = user_repository.get(user_id)

        user.rename(UserName(req.name))
        user_repository.update(user)

        user = user_repository.get(user_id)

        u = userpb.User(id=user.user_id.value, name=user.name.value)
        res = userpb.RenameResponse(user=u)
        await stream.send_message(res)

    async def Delete(
        self, stream: Stream[userpb.DeleteRequest, userpb.DeleteResponse]
    ) -> None:
        req = await stream.recv_message()
        assert req is not None

        user_id = UserID(req.user_id)
        user = user_repository.get(user_id)
        user_repository.remove(user)

        res = userpb.DeleteResponse()
        await stream.send_message(res)
