build:
	@go build -o ./bin/go-1brc .

run: build
	@./bin/go-1brc ./data/input-100m.txt