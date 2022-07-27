from locust import HttpUser, task


class Fetcher(HttpUser):
    @task
    def fetch_todo(self):
        self.client.get("/todos/1")
