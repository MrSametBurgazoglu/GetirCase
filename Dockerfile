FROM golang:1.20.0

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/webapp
ENV MONGO_CONNECTION_STRING="mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY go.mod .
COPY go.sum .
COPY local.env .

RUN go mod download
COPY . .
RUN go build cmd/app/main.go


EXPOSE 8080
CMD ["./main"]