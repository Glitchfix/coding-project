swag:
	# This line generates the files in docs folder for swagger
	swag init --parseDependency --parseInternal --parseDepth 1
run: swag
	go run main.go