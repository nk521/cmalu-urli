FROM golang:1.19-bullseye

WORKDIR /app

COPY init.sh ./

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY cmalu_urli.go ./

RUN go build -o ./cmalu_urli

EXPOSE 8080

CMD ["bash",  "./init.sh" ]