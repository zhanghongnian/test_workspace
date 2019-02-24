export GO111MODULE=on

tidy:
	go mod tidy -v
	
resize:
	go run cmd/resize/resize.go
