gen:
	swagger generate spec -o ./spec/swagger.json

cli:
	@echo "=== downloading ==="
	git clone https://github.com/swagger-api/swagger-ui /tmp/swagger-ui
	mv /tmp/swagger-ui/dist/* ./spec/
	rm -rf /tmp/swagger-ui

	find ./spec -type f -name "*.map" -o -name "*.js" -o -name "*.html" | xargs sed -i "" 's,https://petstore.swagger.io/v2,.,g'
