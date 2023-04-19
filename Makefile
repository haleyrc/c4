all: build diagrams clean

examples = basic components containers systems theming

$(examples): tmp
	go run ./examples/$@ > ./tmp/$@.txt
	java -jar ./plantuml/plantuml.jar -o ./out ./tmp/*.txt
	cp ./tmp/out/*.png ./examples/

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean:
	rm -rf tmp

.PHONY: diagrams
diagrams: $(examples)

.PHONY: tmp
tmp:
	mkdir -p tmp
