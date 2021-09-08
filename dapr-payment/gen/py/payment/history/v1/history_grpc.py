# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: payment/history/v1/history.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import google.api.annotations_pb2
import google.protobuf.empty_pb2
import payment.history.v1.history_pb2


class HistoryServiceBase(abc.ABC):

    @abc.abstractmethod
    async def RegisterHistory(self, stream: 'grpclib.server.Stream[payment.history.v1.history_pb2.RegisterHistoryRequest, payment.history.v1.history_pb2.History]') -> None:
        pass

    @abc.abstractmethod
    async def GetHistory(self, stream: 'grpclib.server.Stream[payment.history.v1.history_pb2.GetHistoryRequest, payment.history.v1.history_pb2.History]') -> None:
        pass

    @abc.abstractmethod
    async def ListHistories(self, stream: 'grpclib.server.Stream[google.protobuf.empty_pb2.Empty, payment.history.v1.history_pb2.History]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/payment.history.v1.HistoryService/RegisterHistory': grpclib.const.Handler(
                self.RegisterHistory,
                grpclib.const.Cardinality.UNARY_UNARY,
                payment.history.v1.history_pb2.RegisterHistoryRequest,
                payment.history.v1.history_pb2.History,
            ),
            '/payment.history.v1.HistoryService/GetHistory': grpclib.const.Handler(
                self.GetHistory,
                grpclib.const.Cardinality.UNARY_UNARY,
                payment.history.v1.history_pb2.GetHistoryRequest,
                payment.history.v1.history_pb2.History,
            ),
            '/payment.history.v1.HistoryService/ListHistories': grpclib.const.Handler(
                self.ListHistories,
                grpclib.const.Cardinality.UNARY_UNARY,
                google.protobuf.empty_pb2.Empty,
                payment.history.v1.history_pb2.History,
            ),
        }


class HistoryServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.RegisterHistory = grpclib.client.UnaryUnaryMethod(
            channel,
            '/payment.history.v1.HistoryService/RegisterHistory',
            payment.history.v1.history_pb2.RegisterHistoryRequest,
            payment.history.v1.history_pb2.History,
        )
        self.GetHistory = grpclib.client.UnaryUnaryMethod(
            channel,
            '/payment.history.v1.HistoryService/GetHistory',
            payment.history.v1.history_pb2.GetHistoryRequest,
            payment.history.v1.history_pb2.History,
        )
        self.ListHistories = grpclib.client.UnaryUnaryMethod(
            channel,
            '/payment.history.v1.HistoryService/ListHistories',
            google.protobuf.empty_pb2.Empty,
            payment.history.v1.history_pb2.History,
        )