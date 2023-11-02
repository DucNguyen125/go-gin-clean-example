.PHONY: tools
tools:
	go install github.com/automation-co/husky@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/conventionalcommit/commitlint@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/segmentio/golines@latest
	go install mvdan.cc/gofumpt@latest

.PHONY: hook
hook: tools
	husky install

.PHONY: update-dependencies
update-dependencies:
	go get -d -u -t ./...

.PHONY: install
install:
	go mod download

.PHONY: test
test:
	# go test -v -coverpkg=./... -coverprofile=profile.cov ./... -parallel 1 -failfast
	sh scripts/coverage.sh
	go tool cover -func profile.cov
