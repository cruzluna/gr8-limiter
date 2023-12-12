# --platform=linux/amd64 
FROM  --platform=linux/amd64 golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -tags netgo -ldflags '-s -w' -o app

EXPOSE 10000
CMD [ "./app", "-env=local" ]

