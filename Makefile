.PHONY: build
build: solvesample

.PHONY: solvesample
solvesample:
	go build solvesample


.PHONY: clean
clean:
	- rm solvesample
	- *.exe
