run:
	@go run cmd/main.go

GetGoPackages:
	@cd cmd && go get && go mod tidy