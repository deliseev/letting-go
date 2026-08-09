.PHONY: test test-examples test-site

test: test-examples test-site

test-examples:
	cd examples && go test ./...

test-site:
	hugo --minify --gc
