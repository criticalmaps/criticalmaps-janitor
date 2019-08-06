FROM golang:1.12-alpine

RUN apk add --no-cache git

RUN mkdir -p /gopath
ENV GOPATH=/gopath
WORKDIR /gopath

RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/lib/pq
RUN go get -u github.com/revel/modules/static

COPY . /gopath/src/github.com/criticalmaps/criticalmaps-janitor

RUN ls /gopath/src/github.com/criticalmaps/criticalmaps-janitor

CMD ["/gopath/bin/revel", "run", "github.com/criticalmaps/criticalmaps-janitor", "prod"]
