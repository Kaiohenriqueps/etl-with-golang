FROM golang

RUN mkdir /var/etl_with_golang
ADD . /var/etl_with_golang/
WORKDIR /var/etl_with_golang

RUN go get -u github.com/lib/pq
