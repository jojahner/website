ODIR=.build
DEVBINARY=website
DOCKERBINARY=website.docker

GC=go build

.DEFAULT_GOAL: build-dev

$(ODIR)/$(DOCKERBINARY):
	GOOS=linux GOARCH=amd64 $(GC) -o $@

$(ODIR)/$(DEVBINARY):
	$(GC) -o $@

get-deps:
	go get -u github.com/labstack/echo
	go get -u github.com/valyala/quicktemplate

build-dev: $(ODIR)/$(DEVBINARY)

docker: $(ODIR)/$(DOCKERBINARY)
	docker build -t website .

clean:
	rm -f $(ODIR)/*

.PHONY: get-deps clean-dev build-dev docker $(ODIR)/$(DEVBINARY) $(ODIR)/$(DOCKERBINARY)
