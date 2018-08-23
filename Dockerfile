FROM golang:1.10.3

# set directory to work
WORKDIR /go

# copy source code 
COPY src/holiday /go/src/holiday
COPY configs /go/configs
COPY atdd /go/atdd

# install dependency
RUN go get -v -d github.com/gin-gonic/gin
RUN go get -v -d gopkg.in/mgo.v2
# go test 
RUN go test holiday/...

RUN go build -o ./bin/holiday holiday/main

ENV ENV development
ENV PORT 3000
ENV MONGO_URL mongodb://mongo:27017

EXPOSE 3000
CMD ./bin/holiday