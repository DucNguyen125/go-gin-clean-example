.PHONY: tools
tools:
	go install github.com/automation-co/husky@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/conventionalcommit/commitlint@latest 

.PHONY: hook
hook:
	husky install