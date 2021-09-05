import ulid
from basic.wallet.v1.wallet_grpc import WalletServiceBase
from basic.wallet.v1.wallet_pb2 import CreateRequest, CreateResponse, Wallet
from grpclib.server import Stream


class WalletService(WalletServiceBase):
    async def Create(self, stream: Stream[CreateRequest, CreateResponse]) -> None:
        req = await stream.recv_message()
        assert req is not None

        wallet = Wallet(id=str(ulid.new()))

        # do something

        res = CreateResponse(wallet=wallet)
        await stream.send_message(res)
