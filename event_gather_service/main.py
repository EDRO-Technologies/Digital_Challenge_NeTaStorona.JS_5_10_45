from typing import Annotated

from fastapi import Depends, FastAPI, HTTPException, status
from pydantic import BaseModel
from config import Config

from apiclient import apiclient

config = Config()

app = FastAPI()

leader_id_client = apiclient.ApiClient(address=config.LEADER_ID_API, client_id=config.CLIENT_ID, client_secret=config.CLIENT_SECRET)


@app.get('/events/leader_id')
async def get_leader_id_events():
    resp = await leader_id_client.get_events()
    return {"data": resp}
