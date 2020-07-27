FROM golang:1.14.5
WORKDIR /usr/app

COPY go.mod .
COPY go.sum .
RUN go get -v -t -d ./...

COPY . .
RUN make build

CMD PORT=$PORT ./build/app
