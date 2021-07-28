FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/Wenchuan-Zhao/goBlogs/src
COPY . $GOPATH/src/github.com/Wenchuan-Zhao/goBlogs/src
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./goBlogs"]