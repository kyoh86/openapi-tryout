# generate all code
generate:
	@$(MAKE) -C simple generate

# generator calls openapi-generator-cli with any arguments: call it with ARGS like "make generator ARGS='help list'"
generator:
	@docker run --rm -u `id -u`:`id -g` -v $(CURDIR):/app openapitools/openapi-generator-cli $(ARGS)
.PHONY: generator-help

# show help message for the Makefile
help:
	@perl -nle'$$c=/^# *(.*)/?$$1:"";/^(?!\.PHONY)(([^ #:](?!\.mk))+):/&&print"$$1|$$p";$$p=$$c' $(MAKEFILE_LIST) | column -ts\|
.PHONY: help
.DEFAULT_GOAL := help
