# refers https://docs.locust.io/en/stable/testing-other-systems.html

import sys
sys.path.append('../../generated/grpc/todo/v1')
import grpc
import todo_pb2_grpc
import todo_pb2
from locust import User, task
import time

# patch grpc so that it uses gevent instead of asyncio
import grpc.experimental.gevent as grpc_gevent

grpc_gevent.init_gevent()


class GrpcClient:
    def __init__(self, environment, stub):
        self.env = environment
        self._stub_class = stub.__class__
        self._stub = stub

    def __getattr__(self, name):
        func = self._stub_class.__getattribute__(self._stub, name)

        def wrapper(*args, **kwargs):
            request_meta = {
                "request_type": "grpc",
                "name": name,
                "start_time": time.time(),
                "response_length": 1,
                "exception": None,
                "context": None,
                "response": None,
            }
            start_perf_counter = time.perf_counter()
            try:
                request_meta["response"] = func(*args, **kwargs)
            except grpc.RpcError as e:
                request_meta["exception"] = e
            request_meta["response_time"] = (time.perf_counter() - start_perf_counter) * 1000
            self.env.events.request.fire(**request_meta)
            return request_meta["response"]

        return wrapper


class GrpcUser(User):
    abstract = True
    stub_class = todo_pb2_grpc.TodoManagerStub

    def __init__(self, environment):
        super().__init__(environment)
        self._channel = grpc.insecure_channel('localhost:50051')
        self._channel_closed = False
        stub = self.stub_class(self._channel)
        self.client = GrpcClient(environment, stub)


class GrpcFetcher(GrpcUser):
    host = 'localhost:50051'
    stub_class = todo_pb2_grpc.TodoManagerStub

    @task
    def fetchTodo(self):
        if not self._channel_closed:
            self.client.FetchTodo(todo_pb2.FetchTodoRequest(id='1'))
