FROM python:3.12-slim

WORKDIR /event_gather_service

COPY . .

RUN pip install -r requirements.txt

ENV PYTHONPATH="${PYTHONPATH}:/event_gather_service"

ENTRYPOINT ["uvicorn", "main:app", "--host=0.0.0.0", "--port=8000"]