FROM ubuntu:16.04

RUN apt-get update && apt-get install -y git sudo wget golang-go

ENV GOROOT=
ENV GOPATH=$HOME/work
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

RUN go get github.com/rcrowley/go-tigertonic
RUN github.com/go-sql-driver/mysql

# clone paper review
WORKDIR /$HOME/work/src/github
RUN mkdir mr-ma
WORKDIR /$HOME/work/src/mr-ma
RUN git clone https://github.com/mr-ma/paper-review-go.git
WORKDIR /$HOME/work/src/mr-ma/paper-review-go

RUN go build
