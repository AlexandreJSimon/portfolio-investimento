deps:
	@ echo
	@ echo "Starting downloading dependencies..."
	@ echo
	@ go get -u ./...

run:
	@ echo
	@ echo "Running Portifolio Investimento..."
	@ echo
	@ go run ./ -fileName=${fileName}