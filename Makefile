.PHONY: build clean deploy

STACK_NAME ?= serverless-aurora
FUNCTIONS := getUsers getUserById postUser deleteUser putUser

# To try different version of Go
GO := go

# Make sure to install aarch64 GCC compilers if you want to compile with GCC.
CC := aarch64-linux-gnu-gcc
GCCGO := aarch64-linux-gnu-gccgo-10

build:
		${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-${function})

build-%:
		cd functions/$* && GOOS=linux GOARCH=arm64 CGO_ENABLED=0 ${GO} build -o bootstrap


clean:
	@rm $(foreach function,${FUNCTIONS}, functions/${function}/bootstrap)
