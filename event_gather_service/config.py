import os
from dotenv import load_dotenv


class Config(object):
    def __init__(self):
        load_dotenv()
        self.LEADER_ID_API = os.getenv("LEADER_ID_API")
        self.CLIENT_ID = os.getenv("CLIENT_ID")
        self.CLIENT_SECRET = os.getenv("CLIENT_SECRET")

        for var in vars(self):
            if self.__getattribute__(var) is None:
                raise ValueError(f"tg_bot: переменная '{var}' должна быть установлена в окружении")

    def __new__(cls):
        if not hasattr(cls, 'instance'):
            cls.instance = super(Config, cls).__new__(cls)
        return cls.instance
