from basic.email.v1.email_grpc import EmailServiceBase
from basic.email.v1.email_pb2 import CreateRequest, CreateResponse, Email
from grpclib.server import Stream


class EmailService(EmailServiceBase):
    async def Create(self, stream: Stream[CreateRequest, CreateResponse]) -> None:
        req = await stream.recv_message()
        assert req is not None

        email = Email(address=f"{req.username}@example.com")

        # do something

        res = CreateResponse(email=email)
        await stream.send_message(res)
