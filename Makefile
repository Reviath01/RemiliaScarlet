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
	@echo "Building to build/remiliascarlet."
	@go build -tags=jsoniter -ldflags="$(LDFLAGS) $(STATIC_LDFLAGS) $(WINDOW_LDFLAGS)" -o build/remiliascarlet$(shell go env GOEXE) main.go
	@echo "Successfully builded, ready to start."

.PHONY: start
start:
	@./build/remiliascarlet
