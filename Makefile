.PHONY: core redis tests

core:
	cd core && go run .

corejs:
	cd core/assets/js/ && npm run build
