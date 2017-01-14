FROM golang:latest

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install
#
# RUN mkdir /app
# ADD web_server /app/web_server
#
# WORKDIR /app
#
# EXPOSE 8080
#
# ENTRYPOINT /app/web_server
