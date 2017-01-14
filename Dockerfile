# FROM golang:latest
#
# RUN mkdir -p /go/src/app
# WORKDIR /go/src/app
#
# CMD ["go-wrapper", "run"]
#
# COPY . /go/src/app
# RUN go-wrapper download
# RUN go-wrapper install
#
# EXPOSE 8080
#

FROM alpine

RUN mkdir -p /website/bin
ADD .build/website.docker /website/bin/website

ENTRYPOINT ["/website/bin/website"]
