from grpclib.server import Server
from grpclib.utils import graceful_exit
from handler import WalletService


async def run(host: str = "0.0.0.0", port: int = 5050) -> None:
    server = Server([WalletService()])

    with graceful_exit([server]):
        await server.start(host, port)
        print(f"Serving on {host}:{port}")
        await server.wait_closed()
