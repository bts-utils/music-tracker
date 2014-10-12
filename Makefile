TEST?=./...

bin:
	@sh -c "$(CURDIR)/scripts/build.sh"

test:
	go test $(TEST) $(TESTARGS) -timeout=10s

.PHONY: bin test