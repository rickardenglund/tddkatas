GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test

.PHONY: test/cover test

test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

testa:
	$(GOTEST) -v  ./...
