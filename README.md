# API-SOAP-JSON-with-GO
## ก่อน run project
```
go get gopkg.in/mgo.v2
go get github.com/gin-gonic/gin
```
## Download Docker image mongo version 3.2.20
```
docker pull mongo:3.2.20
```
## run docker image post 27017
```
 docker container run -p 27017:27017 mongo:3.2.20
```

### Step run project 
- Set GOPATH
```
export GOPATH=`pwd`
```
- Run Project
```
go run src/holiday/main/main.go
```
#### !!! ก่อน run project กำหนด environment ก่อน !!!
```
ENV={ชื่อ environment}
ENV=development go run src/holiday/main/main.go
```
