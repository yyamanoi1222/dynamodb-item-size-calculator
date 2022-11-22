dynamodb-item-size-calculator: cmd/*/*.go pkg/*/*.go
	go build -o dynamodb-item-size-calculator cmd/dynamodb-item-size-calculator/main.go
