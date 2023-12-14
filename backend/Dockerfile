# --platform=linux/amd64 required for render
FROM  golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -tags netgo -ldflags '-s -w' -o app

# Install redis  
RUN apk --update add redis

EXPOSE 6379:6379
EXPOSE 3000/tcp
CMD redis-server --daemonize yes && ./app -env=prod

