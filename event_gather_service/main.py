from typing import Annotated

from fastapi import Depends, FastAPI, HTTPException, status
from schemas.leader_id_schemas import Event
from config import Config
from typing import List
from fastapi.middleware.cors import CORSMiddleware

from apiclient import apiclient

config = Config()

app = FastAPI()

origins = ["*"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

leader_id_client = apiclient.ApiClient(address=config.LEADER_ID_API, client_id=config.CLIENT_ID, client_secret=config.CLIENT_SECRET)


@app.get('/events/leader_id', response_model=List[Event])
async def get_leader_id_events():
    resp = await leader_id_client.get_events()
    return resp.get('items')
