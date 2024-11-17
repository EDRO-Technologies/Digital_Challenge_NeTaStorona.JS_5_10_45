FROM node:18-alpine AS build

WORKDIR /app

COPY package.json ./
COPY . .

RUN npm install
CMD ["npm", "run", "dev"]
