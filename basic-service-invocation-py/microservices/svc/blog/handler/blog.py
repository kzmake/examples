from basic.blog.v1.blog_grpc import BlogServiceBase
from basic.blog.v1.blog_pb2 import Blog, CreateRequest, CreateResponse
from grpclib.server import Stream


class BlogService(BlogServiceBase):
    async def Create(self, stream: Stream[CreateRequest, CreateResponse]) -> None:
        req = await stream.recv_message()
        assert req is not None

        blog = Blog(url=f"https://example.com/{req.username}")

        # do something

        res = CreateResponse(blog=blog)
        await stream.send_message(res)
