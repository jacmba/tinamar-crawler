FROM golang:alpine

RUN mkdir -p /src
RUN mkdir -p /app

WORKDIR /src

COPY . .

RUN apk add --no-cache git

RUN go get go.mongodb.org/mongo-driver/mongo
RUN go build -o /app/tinamar-crawler ./src

WORKDIR /app

CMD ./tinamar-crawler