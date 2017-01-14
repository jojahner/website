BINARY=.build/web_server

.DEFAULT_GOAL: build-dev

get-deps:
	go get -u github.com/labstack/echo
	go get -u github.com/valyala/quicktemplate

build-dev:
	go build -o ${BINARY}

clean-dev:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: get-deps clean-dev build-dev
