from cloudevents.sdk.event import v1
from dapr.ext.grpc import App
from handler import BlogService

app = App()
handler = BlogService()


@app.subscribe(pubsub_name="pubsub", topic="user.v1.CreatedEvent")
def handle_created_event(event: v1.Event) -> None:
    handler.Create(event)


def run(port: int = 5050) -> None:
    app.run(port)
