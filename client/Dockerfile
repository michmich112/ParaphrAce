FROM node:14-alpine

WORKDIR /usr/local/client

COPY . .

RUN npm install
RUN npm run build

EXPOSE 5000

CMD ["npm","run", "docker:start"]
