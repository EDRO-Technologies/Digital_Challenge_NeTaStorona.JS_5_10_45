import os


class Config(object):
    def __init__(self):
        self.LEADER_ID_API = "https://apps.leader-id.ru/api/v1"
        self.CLIENT_ID = "ccc2a1b4-0729-48ef-9b39-dbd480bac841"
        self.CLIENT_SECRET = "k1JtDqYSt9cxWEMLcYuPDEoBeOHv7PaO"
        # self.FRONTEND_URL = os.getenv("FRONTEND_URL")

        for var in vars(self):
            if self.__getattribute__(var) is None:
                raise ValueError(f"tg_bot: переменная '{var}' должна быть установлена в окружении")

    def __new__(cls):
        if not hasattr(cls, 'instance'):
            cls.instance = super(Config, cls).__new__(cls)
        return cls.instance
