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

    @abc.abstractmethod
    async def List(self, stream: 'grpclib.server.Stream[basic.user.v1.user_pb2.ListRequest, basic.user.v1.user_pb2.ListResponse]') -> None:
        pass

    @abc.abstractmethod
    async def Get(self, stream: 'grpclib.server.Stream[basic.user.v1.user_pb2.GetRequest, basic.user.v1.user_pb2.GetResponse]') -> None:
        pass

    @abc.abstractmethod
    async def Rename(self, stream: 'grpclib.server.Stream[basic.user.v1.user_pb2.RenameRequest, basic.user.v1.user_pb2.RenameResponse]') -> None:
        pass

    @abc.abstractmethod
    async def Delete(self, stream: 'grpclib.server.Stream[basic.user.v1.user_pb2.DeleteRequest, basic.user.v1.user_pb2.DeleteResponse]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/basic.user.v1.UserService/Create': grpclib.const.Handler(
                self.Create,
                grpclib.const.Cardinality.UNARY_UNARY,
                basic.user.v1.user_pb2.CreateRequest,
                basic.user.v1.user_pb2.CreateResponse,
            ),
            '/basic.user.v1.UserService/List': grpclib.const.Handler(
                self.List,
                grpclib.const.Cardinality.UNARY_UNARY,
                basic.user.v1.user_pb2.ListRequest,
                basic.user.v1.user_pb2.ListResponse,
            ),
            '/basic.user.v1.UserService/Get': grpclib.const.Handler(
                self.Get,
                grpclib.const.Cardinality.UNARY_UNARY,
                basic.user.v1.user_pb2.GetRequest,
                basic.user.v1.user_pb2.GetResponse,
            ),
            '/basic.user.v1.UserService/Rename': grpclib.const.Handler(
                self.Rename,
                grpclib.const.Cardinality.UNARY_UNARY,
                basic.user.v1.user_pb2.RenameRequest,
                basic.user.v1.user_pb2.RenameResponse,
            ),
            '/basic.user.v1.UserService/Delete': grpclib.const.Handler(
                self.Delete,
                grpclib.const.Cardinality.UNARY_UNARY,
                basic.user.v1.user_pb2.DeleteRequest,
                basic.user.v1.user_pb2.DeleteResponse,
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
        self.List = grpclib.client.UnaryUnaryMethod(
            channel,
            '/basic.user.v1.UserService/List',
            basic.user.v1.user_pb2.ListRequest,
            basic.user.v1.user_pb2.ListResponse,
        )
        self.Get = grpclib.client.UnaryUnaryMethod(
            channel,
            '/basic.user.v1.UserService/Get',
            basic.user.v1.user_pb2.GetRequest,
            basic.user.v1.user_pb2.GetResponse,
        )
        self.Rename = grpclib.client.UnaryUnaryMethod(
            channel,
            '/basic.user.v1.UserService/Rename',
            basic.user.v1.user_pb2.RenameRequest,
            basic.user.v1.user_pb2.RenameResponse,
        )
        self.Delete = grpclib.client.UnaryUnaryMethod(
            channel,
            '/basic.user.v1.UserService/Delete',
            basic.user.v1.user_pb2.DeleteRequest,
            basic.user.v1.user_pb2.DeleteResponse,
        )
