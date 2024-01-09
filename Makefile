python-test:
	python3 python/test_die.py

python-die-run:
	python3 python/die.py 6d6

go-test:
	cd go && go test -v ./...

go-build:
	go build -o build/go/die ./go/main.go

go-die-run:
	go run go/main.go 6d6

run: python-die-run go-die-run
test: python-test go-build go-test