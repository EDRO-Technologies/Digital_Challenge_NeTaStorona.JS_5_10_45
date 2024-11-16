from pydantic import BaseModel
from datetime import datetime
from typing import List, Optional


class Address(BaseModel):
    city: str
    title: str


class Space(BaseModel):
    name: str
    address: Address


class Theme(BaseModel):
    name: str


class Type(BaseModel):
    name: str


class Event(BaseModel):
    id: int
    full_name: str
    full_info: str
    date_start: datetime
    date_end: datetime
    themes: List[Theme]
    type: Type
    format: str
    space: Optional[Space]
