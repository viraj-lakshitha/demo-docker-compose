FROM node:16
WORKDIR /app

COPY package*.json ./
RUN yarn install --ignore-engines

COPY . .
EXPOSE 3000
ENTRYPOINT ["yarn","start"]
