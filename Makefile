.PHONY: run clean build

run: build
	./tinamar-crawler

build:
	go build -o tinamar-crawler

clean:
	rm tinamar-crawler