FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmalu_urli.go ./

RUN go build -o ./cmalu_urli

EXPOSE 8080

CMD [ "./cmalu_urli" ]