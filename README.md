# quote
Display random quotes evertimes it is run:

## install golang
### install golang
```
https://golang.org/doc/install
```

### set path: vi ~/.bash_profile
```
export GOPATH=/Users/nandeshwar.sah/go
export GOROOT=/usr/local/go

export PATH=$PATH:/Users/nandeshwar.sah/go/bin
export PATH=$PATH:/usr/local/go/bin
export GOBIN=/usr/local/go/bin

alias dc='docker-compose'
alias dcup='docker-compose up -d'
alias dcdown='docker-compose down'
```

## How to run? 
```
go run cmd/quote/main.go
```

## To get test coverage
```
go test $(go list ./pkg/*) -coverprofile r.txt
go tool cover -func r.txt
go tool cover -html r.txt
```

## Run using docker container (default server run duration 5 minutes)
```
docker run -p 1922:1922 -t nandeshwar/quote
or
docker run -p 1922:1922 -e SERVER_RUN_DURATION_MIN=10 -e SERVER_RUN_DURATION_HOUR=10 -t nandeshwar/quote
```

## Run using docker compose 
```
docker-compose up -d
```

### Default time zone is set to America/Denver, can be changed with env variable  TZ=America/Denver
```
docker run -p 1922:1922 -e SERVER_RUN_DURATION_MIN=10 -e SERVER_RUN_DURATION_HOUR=10 -e TZ=America/Denver -t nandeshwar/quote
```

### build image and push Image to docker container
```
docker build -t nandeshwar/quote .
docker push nandeshwar/quote

or
 
docker-compose up -d --build
docker push nandeshwar/quote
docker-compose down
docker images -q -f "dangling=true" | xargs docker rmi
```

### Create image compatible for Raspberry pi
```
Step1: Create image specific to Raspberry pi 

Install buildx with below commands,
git clone git://github.com/docker/buildx && cd buildx
make install
Create image compatible with raspberry pi with below command,
docker buildx create --name testbuilder
docker buildx ls

docker buildx use testbuilder

docker buildx build --platform linux/amd64,linux/arm64 --tag nandeshwar/quote-r:latest --push .

docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 --tag nandeshwar/quote-r:latest --push .

```

### Enter to MySqlite3
```
sqlite3 ./db/quote.db 
```

### List tables
```
.tables
```

## Swagger setup in local
### Install go-swagger in mac
```
brew tap go-swagger/go-swagger
brew install go-swagger
```

### Setup code for swagger
```
1. copy swagger-ui folder as it is

2. change title in index.html

3. api.go has following lines at top of file and should be commented as given below
// Package Quote QuoteAPI.
//
//     Consumes:
//		- application/xml
//     Produces:
//      - application/json
//
// swagger:meta

4. Add line below in api.go
//go:generate swagger generate spec -m -o ../../swagger-ui/swagger.json

5. Link these 2 lines given below api.go
sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/")))
mux.PathPrefix("/swagger-ui/").Handler(sh)

6. Add swagger tag in struct like given below
// swagger:model infoResponse
type Info struct {

7. Add swagger tag before function similar to below
// swagger:operation GET /api/quote/v1/info/{id} INFO info
// ---
// description: get INFO by id
// consumes:
// - "application/json"
// parameters:
// - name: id
//   description: id to get info
//   in: path
//   required: true
//   default: 1
//   type: string
// Responses:
//   '200':
//     description: Ok
//     schema:
//        '$ref': '#/definitions/infoResponse'
//   '400':
//     description: Bad request
//   '404':
//     description: Not found
//   '500':
//     description: Internal server error
```

### Generate Swagger documentation
```
go generate ./...
```

### Local Swagger UI URL
```
http://localhost:1922/swagger-ui/
```

### Docker compose Swagger UI URL
```
http://localhost:1922/swagger-ui/
```

### Install grpc related utility
```
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

go get -u google.golang.org/grpc

```
### go generate pb file grpc - EventDetail
```
protoc -I=./pkg/grpcquote/ --go_out=./pkg/grpcquote/ --go-grpc_out=./pkg/grpcquote/ ./pkg/grpcquote/event-detail.proto
```

# Development setup for grpc
https://grpc.io/docs/quickstart/go.html

Look "Before you begin" section
## For endpoints using grpc
https://github.com/phuongdo/go-grpc-tutorial
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

### run grpc client to get & update record for event-detail
#### default get request
```
go run cmd/grpc_client/eventDetailClient.go
```

#### get event detail
```
go run cmd/grpc_client/eventDetailClient.go get 41
```

#### update event detail
```
go run cmd/grpc_client/eventDetailClient.go put event.json
```

#### get multiple event detail - grpc stream
```
go run cmd/grpc_client/eventDetailClient.go gets "1, 2, 3, 4, 5"
```

#### Kafka UI
```
http://localhost:8000/
```
#### Run Kafka producer
```
go run cmd/binaryclient/kafka/producer/kafkaproducer.go
```

#### Run Kafka Consumer
```
go run cmd/binaryclient/kafka/consumer/kafkaconsumer.go
```
#### RabbitMQ UI - user: rabbitmq, password: rabbitmq
```
http://localhost:15672/
```
#### RabbitMQ Producer
```
sudo go run cmd/binaryclient/rabbitmq/producer/producer.go
```
#### RabbitMQ Consumer
```
sudo go run cmd/binaryclient/rabbitmq/consumer/consumer.go
```

#### Email Server setting
```
export EMAIL_SERVER="smtp.gmail.com"
export EMAIL_SERVER_PORT=587
export EMAIL_FROM="abc@gmail.com"
export EMAIL_FROM_PWD="****"
export EMAIL_TO_FOR_EVENTS="abc@gmail.com, xyz@gmail.com"
export EMAIL_TO_FOR_QUOTE_IMAGE="abc@gmail.com, xyz@gmail.com"

```
Note: for smtp gmail. go to gmail acccount -> Search -> Less secure app access -> turn it on


#### To setup git repo
```
1. Add ssh public key in github
2. git remote set-url origin git@github.com:Nandeshwar/quote.git
3. git push origin <branch>
```

