# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: greeter/v1/greeter.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import google.api.annotations_pb2
import greeter.v1.greeter_pb2


class GreeterBase(abc.ABC):

    @abc.abstractmethod
    async def Hello(self, stream: 'grpclib.server.Stream[greeter.v1.greeter_pb2.HelloRequest, greeter.v1.greeter_pb2.HelloResponse]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/basic.greeter.v1.Greeter/Hello': grpclib.const.Handler(
                self.Hello,
                grpclib.const.Cardinality.UNARY_UNARY,
                greeter.v1.greeter_pb2.HelloRequest,
                greeter.v1.greeter_pb2.HelloResponse,
            ),
        }


class GreeterStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.Hello = grpclib.client.UnaryUnaryMethod(
            channel,
            '/basic.greeter.v1.Greeter/Hello',
            greeter.v1.greeter_pb2.HelloRequest,
            greeter.v1.greeter_pb2.HelloResponse,
        )
