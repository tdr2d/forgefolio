.PHONY: core redis tests

help: 
	@fgrep -h "##" Makefile | fgrep -v "fgrep" | gsed -r 's/(.*):.*##(.*)/\1:\2/' - | column -s: -t | sed -e 's/##//'

installjs:  ## Install javascript dependencies
	cd core/assets/js/blog && npm install

blogjs: ## Compile js for the blog-posts screen
	cd core/assets/js/blog && npm run build

core: blogjs ## Run core backend
	cd core && go run .

redis:
	redis-server


