##
## eat Makefile
##

# -----------------------------------
#  # Makefile自体の設定
# -----------------------------------

RAND=`date "+%m%H%M%d%S"`

# ====== misc

help: ## ヘルプを表示
	@echo ""
	@grep "^##" $(MAKEFILE_LIST)
	@echo ""
	@grep "^[0-9a-zA-Z\-_]*: #" $(MAKEFILE_LIST) | sed -e 's/^[^:]*://' | sed -e 's/^/make /' | sed -e 's/: *#\([^#]\)/:\1/' | sed -e 's/://' | sed -e 's/##/#/'
	@echo ""

generate_api_codes: # APIコードを自動生成します
	rm -rf ./openapi/*
	openapi-generator generate -i openapi.yaml -g go-server -o ./openapi --additional-properties packageName=viewmodel --global-property apiTests=false,modelTests=false
	cp openapi/go/model_* domain/viewmodel/
