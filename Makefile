.PHONY: clean build deploy

STACK_NAME ?= invoice-dashboard
FUNCTIONS := getInvoiceById getInvoices createInvoice modifyInvoice removeInvoice

# To try different version of Go
GO := go

# Make sure to install aarch64 GCC compilers if you want to compile with GCC.
CC := aarch64-linux-gnu-gcc
GCCGO := aarch64-linux-gnu-gccgo-10

run-containers:
		docker compose --env-file ./config/dev.yml up
build-cmd:
		go build -o bin/main -v cmd/main.go
run-cmd:
		bin/main

build:
		${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-${function})

build-%:
		cd functions/$* && GOOS=linux GOARCH=arm64 CGO_ENABLED=0 ${GO} build -o bootstrap

build-gcc:
		${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-gcc-${function})

build-gcc-%:
		cd functions/$* && GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=${CC} ${GO} build -o bootstrap

build-gcc-optimized:
		${MAKE} ${MAKEOPTS} $(foreach function,${FUNCTIONS}, build-gcc-optimized-${function})

build-gcc-optimized-%:
		cd functions/$* && GOOS=linux GOARCH=arm64 GCCGO=${GCCGO} ${GO} build -compiler gccgo -gccgoflags '-static -Ofast -march=armv8.2-a+fp16+rcpc+dotprod+crypto -mtune=neoverse-n1 -moutline-atomics' -o bootstrap

clean:
	@rm $(foreach function,${FUNCTIONS}, functions/${function}/bootstrap)


invoke-get-all:
	@sam local invoke --env-vars env-vars.json --event functions/getInvoices/event.json GetAllInvoicesFunction

invoke-get:
	@sam local invoke --env-vars env-vars.json --event functions/getInvoiceById/event.json GetInvoiceByIdFunction

invoke-create:
	@sam local invoke --env-vars env-vars.json --event functions/createInvoice/event.json CreateInvoiceFunction
	
invoke-put:
	@sam local invoke --env-vars env-vars.json --event functions/modifyInvoice/event.json ModifyInvoiceFunction

invoke-delete:
	@sam local invoke --env-vars env-vars.json --event functions/removeInvoice/event.json RemoveInvoiceFunction


deploy:
	if [ -f samconfig.toml ]; \
		then sam deploy --stack-name ${STACK_NAME}; \
		else sam deploy -g --stack-name ${STACK_NAME}; \
  fi
