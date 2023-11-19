FROM golang:1.20
  
RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

MAINTAINER "公众号: 玩转测试开发"

WORKDIR /home/workspace

ADD . /home/workspace

RUN go mod init momo

RUN go mod tidy

RUN go get -u github.com/gin-gonic/gin

RUN go build main.go

EXPOSE 9090

ENTRYPOINT ["./main"]