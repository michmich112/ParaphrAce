FROM node:14-alpine

WORKDIR /usr/local/client

COPY . .

ARG PARAPHRACE_API_URL
ENV PARAPHRACE_API_URL=${PARAPHRACE_API_URL}

RUN npm install
RUN npm run build

EXPOSE 5000

CMD ["npm","run", "docker:start"]
