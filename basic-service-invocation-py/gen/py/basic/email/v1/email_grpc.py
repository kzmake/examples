# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: basic/email/v1/email.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import google.api.annotations_pb2
import basic.email.v1.email_pb2


class EmailServiceBase(abc.ABC):

    @abc.abstractmethod
    async def Create(self, stream: 'grpclib.server.Stream[basic.email.v1.email_pb2.CreateRequest, basic.email.v1.email_pb2.CreateResponse]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/basic.email.v1.EmailService/Create': grpclib.const.Handler(
                self.Create,
                grpclib.const.Cardinality.UNARY_UNARY,
                basic.email.v1.email_pb2.CreateRequest,
                basic.email.v1.email_pb2.CreateResponse,
            ),
        }


class EmailServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.Create = grpclib.client.UnaryUnaryMethod(
            channel,
            '/basic.email.v1.EmailService/Create',
            basic.email.v1.email_pb2.CreateRequest,
            basic.email.v1.email_pb2.CreateResponse,
        )
