import basic.blog.v1.blog_pb2 as blogpb
import basic.email.v1.email_pb2 as emailpb
import basic.wallet.v1.wallet_pb2 as walletpb
import ulid
from basic.blog.v1.blog_grpc import BlogServiceStub
from basic.email.v1.email_grpc import EmailServiceStub
from basic.user.v1.user_grpc import UserServiceBase
from basic.user.v1.user_pb2 import CreateRequest, CreateResponse, User
from basic.wallet.v1.wallet_grpc import WalletServiceStub
from grpclib.client import Channel
from grpclib.server import Stream


class UserService(UserServiceBase):
    async def Create(self, stream: Stream[CreateRequest, CreateResponse]) -> None:
        req = await stream.recv_message()
        assert req is not None

        user = User(id=str(ulid.new()), name=req.name)

        async with Channel(host="localhost", port=50001) as channel:
            blog = BlogServiceStub(channel)
            blogres = await blog.Create(
                blogpb.CreateRequest(username=req.name),
                metadata={"dapr-app-id": "svc-blog"},
            )
            user.blog_url = blogres.blog.url

        async with Channel(host="localhost", port=50001) as channel:
            email = EmailServiceStub(channel)
            email_res = await email.Create(
                emailpb.CreateRequest(username=req.name),
                metadata={"dapr-app-id": "svc-email"},
            )
            user.email_address = email_res.email.address

        async with Channel(host="localhost", port=50001) as channel:
            wallet = WalletServiceStub(channel)
            wallet_res = await wallet.Create(
                walletpb.CreateRequest(),
                metadata={"dapr-app-id": "svc-wallet"},
            )
            user.wallet_id = wallet_res.wallet.id

        # do something

        res = CreateResponse(user=user)
        await stream.send_message(res)
