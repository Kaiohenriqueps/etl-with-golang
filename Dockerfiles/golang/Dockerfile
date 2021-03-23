FROM golang

RUN go get -u github.com/lib/pq
RUN go get -u github.com/joho/godotenv

RUN mkdir /var/etl
ADD . /var/etl/
WORKDIR /var/etl