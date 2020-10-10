FROM golang:1.15.2 AS builder

WORKDIR /quote
COPY . .

RUN CGO_ENABLED=1 go install -v ./...

#FROM alpine:3.8
# Switched to debian from alpine to enable CGO required by SQLite3. alpine complains at run time when container is built with CGO_ENABLED=1
FROM amd64/debian:bullseye-20200908
COPY --from=builder /go/bin/quote /quote

COPY image /image
COPY image-motivational /image-motivational

COPY db /db
COPY views /views

#RUN apk add --no-cache tzdata
RUN apt-get install -y tzdata
ENV TZ=America/Denver

EXPOSE 1922
ENTRYPOINT ["/quote"]
