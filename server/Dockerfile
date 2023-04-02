FROM golang:1.20.2-bullseye as builder

WORKDIR /go/src/app

RUN go install github.com/cespare/reflex@latest

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# ENV PATH=$PATH:$HOME/go/bin
RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]
