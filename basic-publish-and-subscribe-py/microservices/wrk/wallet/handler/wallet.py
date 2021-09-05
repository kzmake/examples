from basic.user.v1.event_pb2 import CreatedEvent
from cloudevents.sdk.event import v1


class WalletService:
    def Create(self, event: v1.Event) -> None:
        print(f"handle Create: {event.EventID()}")
        e = CreatedEvent()
        e.ParseFromString(event.Data())

        print(f"recv {e.user.id}")

        # do something
        pass
