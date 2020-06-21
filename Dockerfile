FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/blog-example
COPY . $GOPATH/blog-example
RUN go build .

EXPOSE 8000
ENTRYPOINT["./blog-example"]