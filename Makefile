.PHONY: core redis tests

help: 
	@fgrep -h "##" Makefile | fgrep -v "fgrep" | gsed -r 's/(.*):.*##(.*)/\1:\2/' - | column -s: -t | sed -e 's/##//'

core: blogjs ## Run core
	cd core && go run .

redis: ## Run redis-server
	redis-server

blogjs: ## Compile js for the blog-posts screen
	cd core/assets/js/blog && npm run build
