# quote
Display random quotes evertimes it is run:

## How to run? 
```
go run cmd/quote/main.go
```

## Run using docker container (default server run duration 5 minutes)
```
docker run -p 1922:1922 -t nandeshwar/quote
or
docker run -p 1922:1922 -e SERVER_RUN_DURATION_MIN=10 SERVER_RUN_DURATION_HOUR=10 -t nandeshwar/quote
```

### Push Image to docker container
```
docker build -t nandeshwar/quote .
docker push nandeshwar/quote
```
