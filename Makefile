watch:
	find . | grep -v .git | entr -c make test

test:
	go test ./...
