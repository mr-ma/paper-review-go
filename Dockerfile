FROM golang:1.8

#RUN apt-get update && apt-get install -y git sudo wget golang-go
#ENV GOROOT=
#ENV GOPATH=$HOME/work
#ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

RUN go get github.com/rcrowley/go-tigertonic
RUN go get github.com/alexedwards/scs
RUN go get github.com/Jeffail/gabs
RUN go get github.com/stretchr/testify/assert
RUN go get github.com/go-sql-driver/mysql
#RUN go get github.com/mr-ma/paper-review-go


# clone paper review
#WORKDIR /go/src/app
WORKDIR /go/src/github
RUN mkdir -p mr-ma/paper-review-go/
WORKDIR /go/src/github/mr-ma/paper-review-go
COPY . .
#WORKDIR /$HOME/work/src/mr-ma
#RUN git clone https://github.com/mr-ma/paper-review-go.git
#WORKDIR /$HOME/work/src/mr-ma/paper-review-go
RUN go build taxonomyserver.go
#EXPOSE 8080
#CMD "./taxonomyserver -mysqluser root -mysqlpass kB*6jd<KPa<M2x -listen 0.0.0.0:8080"
