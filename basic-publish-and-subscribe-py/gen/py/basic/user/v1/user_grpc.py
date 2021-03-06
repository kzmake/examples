# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: basic/user/v1/user.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import google.api.annotations_pb2
import basic.user.v1.user_pb2


class UserServiceBase(abc.ABC):

    @abc.abstractmethod
    async def Create(self, stream: 'grpclib.server.Stream[basic.user.v1.user_pb2.CreateRequest, basic.user.v1.user_pb2.CreateResponse]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/basic.user.v1.UserService/Create': grpclib.const.Handler(
                self.Create,
                grpclib.const.Cardinality.UNARY_UNARY,
                basic.user.v1.user_pb2.CreateRequest,
                basic.user.v1.user_pb2.CreateResponse,
            ),
        }


class UserServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.Create = grpclib.client.UnaryUnaryMethod(
            channel,
            '/basic.user.v1.UserService/Create',
            basic.user.v1.user_pb2.CreateRequest,
            basic.user.v1.user_pb2.CreateResponse,
        )
