# generate all code
#
# NOTE: openapi-generator-cliのgenerateコマンドを使ってコードを生成する
# - docker runでvolumes(-v)でカレントディレクトリをマウントし、
#   CLIのイメージを経由して呼び出しているため、ファイルの権限を整合するため
#   docker上で現在のユーザーID(id -u)とグループID(id -g)で実行している
# - clientコードserverコードともに、global-propertyでAPIとモデル定義、
#   そのテストコードだけを生成するようにしている
# - 各generatorは専用のConfig(config/**.yaml)を渡すことで少し挙動をカスタマイズしている
generate: generate-go-server generate-go-client
.PHONY: generate

# generate go server code to gen/server
generate-go-server:
	@docker run --rm -u `id -u`:`id -g` -v $(CURDIR):/app openapitools/openapi-generator-cli generate --global-property models,modelTests,apis,apiTests -i /app/openapi.yaml -o /app/gen -c /app/config/go-server.yaml -g go-server
.PHONY: generate-go-server

# generate go server code to gen/client
generate-go-client:
	@docker run --rm -u `id -u`:`id -g` -v $(CURDIR):/app openapitools/openapi-generator-cli generate --global-property models,modelTests,apis,apiTests -i /app/openapi.yaml -o /app/gen/client -c /app/config/go-client.yaml -g go
.PHONY: generate-go-client

serve-document:
	@docker run --rm -p 8080:8080 -v $(CURDIR):/app -e SWAGGER_JSON=/app/openapi.yaml swaggerapi/swagger-ui
.PHONY: serve-document

# show help message for the Makefile
help:
	@perl -nle'$$c=/^# *(.*)/?$$1:"";/^(?!\.PHONY)(([^ #:](?!\.mk))+):/&&print"$$1|$$p";$$p=$$c' $(MAKEFILE_LIST) | column -ts\|
.PHONY: help
.DEFAULT_GOAL := help
