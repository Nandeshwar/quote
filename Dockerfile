FROM golang:1.15.2 AS builder

WORKDIR /quote
COPY . .

RUN CGO_ENABLED=1 go install -v ./...

#FROM alpine:3.8

# Switched to debian from alpine to enable CGO required by SQLite3. alpine complains at run time when container is built with CGO_ENABLED=1
#FROM amd64/debian:bullseye-20200908

# Switched to centos: Because there is issue with newrelic -
# error was-
# (1) 2020/11/14 19:57:41.334152 {"level":"warn","msg":"application connect failure","context":{"error":"Post \"https://collector.newrelic.com/agent_listener/invoke_raw_method?license_key=96a09863f7b19aa298a8a1dNRAL\u0026marshal_format=json\u0026method=preconnect\u0026protocol_version=17\": x509: certificate signed by unknown authority"}}
FROM centos:latest
COPY --from=builder /go/bin/quote /quote

COPY image /image
COPY image-motivational /image-motivational

COPY db /db
COPY views /views

COPY swagger-ui /swagger-ui

#RUN apk add --no-cache tzdata // alpine
#RUN apt-get install -y tzdata // debian
#RUN yum -y upgrade
RUN yum -y install tzdata
ENV TZ=America/Denver

EXPOSE 1922
EXPOSE 1923
ENTRYPOINT ["/quote"]
