.PHONY: h push

h:
	@echo 'push PUSHを実行'

push:
	$(GOPATH)/bin/air -c .air.push.toml
