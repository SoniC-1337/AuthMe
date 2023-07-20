import requests as r


class AuthMe:
    """Main AuthMe class. Create an instance of this class with the UID and URL/IP of the server to configure.

    Args:
        url (str): The URL/IP of the server to configure.
        uid (str): The UID of the user to authenticate.
    """
    def __init__(self, url, uid):
        self.url = url
        self.uid = uid

    def authenticate(self) -> bool:
        """Passes the UID to the login endpoint and validates product access.

        Returns:
            bool: True if the user is authenticated, False otherwise.
        """

        payload = {
            'UID': self.uid,
        }
        try:
            if r.post(self.url, data=payload).status_code == 200:
                return True
        except Exception as e:
            print(e)
            return False




# Authentication test
"""
test = AuthMe('http://localhost:8080/login', 'Xoro')
print(test.authenticate())
"""
