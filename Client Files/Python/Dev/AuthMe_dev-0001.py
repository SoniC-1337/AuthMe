import hashlib
import hmac
import time

import requests


class AuthMe:
    endpoint = ''

    def __init__(self, endpoint):
        self.endpoint = endpoint  # URL/IP of the AuthMe server

    def init(self):
        if self.sessionid != '':
            return 'Already initialized'

        post_data = {
            'req': 'init',
        }

    def __do_request(self, post_data):
        try:
            response = requests.post(
                self.endpoint, data=post_data, timeout=10
            )

            key = self.secret if post_data["type"] == "init" else self.enckey

            client_computed = hmac.new(key.encode('utf-8'), response.text.encode('utf-8'), hashlib.sha256).hexdigest()

            signature = response.headers["signature"]

            if not hmac.compare_digest(client_computed, signature):
                print("Signature checksum failed. Request was tampered with or session ended most likely.")
                print("Response: " + response.text)
                time.sleep(3)
                exit(1)

            return response.text
        except requests.exceptions.Timeout:
            print("Request timed out. Server is probably down/slow at the moment")
