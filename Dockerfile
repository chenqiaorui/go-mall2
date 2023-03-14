FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

RUN go env -w GOPROXY=https://goproxy.cn

RUN go install github.com/cortesi/modd/cmd/modd@latest

COPY . .
#RUN go build -v -o /usr/local/bin/app ./...

ENTRYPOINT ["sh", "-c", "while true;do echo  hello;sleep 10;done"]
#CMD ["app"]
