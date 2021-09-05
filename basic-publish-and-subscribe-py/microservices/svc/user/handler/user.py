import ulid
from basic.user.v1.event_pb2 import CreatedEvent, CreatedEventUser
from basic.user.v1.user_grpc import UserServiceBase
from basic.user.v1.user_pb2 import CreateRequest, CreateResponse, User
from dapr.clients import DaprClient
from google.protobuf.timestamp_pb2 import Timestamp
from grpclib.server import Stream


class UserService(UserServiceBase):
    async def Create(self, stream: Stream[CreateRequest, CreateResponse]) -> None:
        req = await stream.recv_message()
        assert req is not None

        user = User(id=str(ulid.new()), name=req.name)

        # do something
        pass

        event = CreatedEvent(
            user=CreatedEventUser(
                id=user.id, name=user.name, created_at=Timestamp().GetCurrentTime()
            ),
        )

        with DaprClient() as d:
            d.publish_event(
                pubsub_name="pubsub",
                topic_name="user.v1.CreatedEvent",
                data=event.SerializeToString(),
            )

        res = CreateResponse(user=user)
        await stream.send_message(res)
