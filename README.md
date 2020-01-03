# quote
Display random quotes evertimes it is run:

## How to run? 
go run cmd/quote/main.go

## Run using docker container
docker run -t nandeshwar/quote

### Push Image to docker container
docker build -t nandeshwar/quote .
docker push nandeshwar/quote
