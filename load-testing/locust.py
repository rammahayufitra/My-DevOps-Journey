from locust import HttpUser, task, between
import time

class MyUser(HttpUser):
    wait_time = between(1, 2)
    counting = 0

    @task
    def submit_data(self):
        headers = {'Content-type': 'application/json'}
        self.counting += 1
        data = {
            'counting': str(self.counting),
            'timestamp': str(time.time())
        }
        self.client.post('/ai_data', json=data, headers=headers)
