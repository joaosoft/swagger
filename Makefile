install:
	@echo "=== installing go-swagger ==="
	brew tap go-swagger/go-swagger
	brew install go-swagger

	@echo "=== installing go-swag ==="
	go get -u github.com/swaggo/swag/cmd/swag

gen-go-swagger:
	@echo "=== creating folder ./spec ==="
	#rm -rf ./spec/swagger.json -f

	@echo "=== generating swagger ==="
	pushd "./go-swagger" && mkdir -p ./docs && swagger generate spec -o ./docs/swagger.json && popd

gen-go-swag:
	@echo "=== creating folder ./spec ==="
	#rm -rf ./spec/swagger.json -f

	@echo "=== generating swagger ==="
	pushd "./go-swag" && mkdir -p ./docs && swag init && popd

cli:
	@echo "=== downloading ==="
	rm -rf ./spec
	#find ./spec/* \! -name 'swagger.json' -delete
	git clone https://github.com/swagger-api/swagger-ui /tmp/swagger-ui

	mkdir -p ./spec
	mv /tmp/swagger-ui/dist/* ./spec/
	rm -rf /tmp/swagger-ui

	find ./spec -type f -name "*.map" -o -name "*.js" -o -name "*.html" | xargs sed -i "" 's,https://petstore.swagger.io/v2,.,g'
