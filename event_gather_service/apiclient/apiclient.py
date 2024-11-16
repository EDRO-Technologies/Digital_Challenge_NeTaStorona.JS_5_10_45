import asyncio
import functools

import httpx
from config import Config


def auth_decorator(func):
    @functools.wraps(func)
    async def wrapper(self, *args, **kwargs):
        headers = {"Authorization": "Bearer " + self.access_token}
        try:
            return await func(self, headers=headers, *args, **kwargs)
        except Exception as e:
            await self.refresh_access_token()
            return await func(self, headers=headers, *args, **kwargs)

    return wrapper


class ApiClient:
    def __init__(self, address: str, client_id: str, client_secret: str) -> None:
        self.user_validated = None
        self.user_id = None
        self.refresh_token = None
        self.access_token = None
        self.address = address
        self.client_id = client_id
        self.client_secret = client_secret
        self.init_auth()

    def authorize(self) -> dict:
        request_data = {
            "client_id": self.client_id,
            "client_secret": self.client_secret,
            "grant_type": "client_credentials"
        }
        with httpx.Client() as client:
            r = client.post(self.address + "/oauth/token", json=request_data)
            if r.status_code == 200:
                return r.json()
            else:
                return {}

    def init_auth(self) -> None:
        auth_response = self.authorize()
        if auth_response:
            self.access_token = auth_response.get("access_token")
            self.refresh_token = auth_response.get("refresh_token")
            self.user_id = auth_response.get("user_id")
            self.user_validated = auth_response.get("user_validated")

    async def refresh_access_token(self) -> None:
        if self.refresh_token is None:
            self.init_auth()
        else:
            request_data = {
                "client_id": self.client_id,
                "client_secret": self.client_secret,
                "grant_type": "refresh_token",
                "refresh_token": self.refresh_token
            }
            async with httpx.AsyncClient() as client:
                r = await client.post(self.address + "/oauth/token", json=request_data)
                if r.status_code == 200:
                    self.access_token = r.json().get("access_token")

    @auth_decorator
    async def get_events(self, headers: dict) -> dict:
        async with httpx.AsyncClient() as client:
            response = await client.get(self.address + "/events", headers=headers)
            print(response.json())
            return response.json()


if __name__ == "__main__":
    config = Config()
    leader_id_client = ApiClient(
        address=config.LEADER_ID_API,
        client_id=config.CLIENT_ID,
        client_secret=config.CLIENT_SECRET,
    )
    asyncio.run(leader_id_client.get_events())
