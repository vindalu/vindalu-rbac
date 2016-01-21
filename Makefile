
.clean:
	rm -f coverage.test
	go clean -i ./...

.test: .clean
	go test -coverprofile=coverage.test ./...
