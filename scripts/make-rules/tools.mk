.PHONY: tools.install.%
tools.install.%:
	@echo "===========> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.addlicense
install.addlicense:
	@$(GO) install github.com/marmotedu/addlicense@latest

.PHONY: install.golines
install.golines:
	@$(GO) install github.com/segmentio/golines@latest

.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest


.PHONY: install.go-gitlint
install.go-gitlint:
	@$(GO) install github.com/marmotedu/go-gitlint/cmd/go-gitlint@latest