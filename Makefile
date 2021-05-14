.NOTPARALLEL: static-arg nowindowsgui build start
SHELL = sh

all: build
static: static-arg build
run: nowindowsgui build start

.PHONY: static-arg
static-arg:
	$(eval STATIC_LDFLAGS = -extldflags "-static")

.PHONY: nowindowsgui
nowindowsgui:
	$(eval WINDOW_LDFLAGS = )

.PHONY: build
build:
	@echo "Building ..."
	@go build -tags=jsoniter -ldflags="$(LDFLAGS) $(STATIC_LDFLAGS) $(WINDOW_LDFLAGS)" -o build/remiliascarlet$(shell go env GOEXE) $$PWD/main.go

.PHONY: start
start:
	@./build/remiliascarlet
