BINARY=.build/web_server

.DEFAULT_GOAL: build

build:
	go build -o ${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean build
