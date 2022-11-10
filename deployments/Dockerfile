FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o cmd/app-search/main cmd/app-search/main.go

EXPOSE 8080

CMD [ "/cmd/app-search/main" ]