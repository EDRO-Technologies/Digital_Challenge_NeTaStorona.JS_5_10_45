FROM node:18-alpine3.17 as build

WORKDIR /app
COPY . /app

EXPOSE 5173

RUN npm install

RUN npm run dev
