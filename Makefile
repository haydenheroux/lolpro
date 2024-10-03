build:
	go build -o bin/ ./cmd/lolpro

view:
	sqlitebrowser ./test.db
