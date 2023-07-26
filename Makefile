.PHONY: tools
tools:
	go install github.com/automation-co/husky@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/conventionalcommit/commitlint@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: hook
hook:
	husky install

.PHONY: update-dependencies
update-dependencies:
	go get -d -u -t ./...

.PHONY: install
install:
	go mod download
