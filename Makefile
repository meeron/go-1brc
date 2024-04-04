build:
	@go build -o ./bin/go-1brc .

run: build
	@./bin/go-1brc ./data/input.txt