.PHONY: build
build: solve distcompile getbodyinfo


.PHONY: solve
solve:
	go build solve.go

.PHONY: distcompile
distcompile:
	go build distcompile.go

.PHONY: getbodyinfo
getbodyinfo:
	go build getbodyinfo.go


.PHONY: clean
clean:
	- rm solve
	- *.exe
