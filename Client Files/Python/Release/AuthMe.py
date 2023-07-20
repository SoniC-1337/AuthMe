import requests as r


class AuthMe:
    def __init__(self, url, username, password):
        self.url = url
        self.username = username
        self.password = password

    def login(self):
        payload = {
            'username': self.username,
            'password': self.password
        }
        r.post(self.url, data=payload)
        return r.get(self.url)
