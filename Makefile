build:
	go build -o build/prepend cmd/prepend/main.go

install: build
	install build/prepend /usr/local/bin
clean:
	rm build/*

