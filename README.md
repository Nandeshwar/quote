# quote
Display random quotes evertimes it is run:

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

### Default time zone is set to America/Denver, can be changed with env variable  TZ=America/Denver
```
docker run -p 1922:1922 -e SERVER_RUN_DURATION_MIN=10 -e SERVER_RUN_DURATION_HOUR=10 -e TZ=America/Denver -t nandeshwar/quote
```

### Push Image to docker container
```
docker build -t nandeshwar/quote .
docker push nandeshwar/quote
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
