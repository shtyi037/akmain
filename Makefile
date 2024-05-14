# 會執行config.local.toml (須自行由config.base.toml複製過去)
local_test:
	go run -v -ldflags "-X main.VERSION=$(RELEASE_TAG)" \
	./cmd/main.go \
	-b ./deploy/config/config.base.toml \
  	-c ./deploy/config/config.local.toml
