from locust import HttpUser, task


class Fetcher(HttpUser):
    @task
    def fetch_todo(self):
        query = '''
        query{
            todo(id: 1) {
                id
                text
                done
                user {
                    id
                    name
                }
            }
        }
        '''
        self.client.post(
            "/query",
            headers={
                "Accept": "application/graphql",
            },
            json={"query": query}
        )
