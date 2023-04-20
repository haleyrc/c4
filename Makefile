all: build diagrams clean

examples = basic components containers systems theming

.PHONY: $(examples)
$(examples): tmp
	go run ./examples/$@ > ./tmp/$@.txt

.PHONY: gallery
gallery: tmp
	go run ./examples/gallery > ./tmp/gallery.txt
	go run ./examples/gallery --sketch > ./tmp/sketch.txt

.PHONY: png
png:
	java -jar ./plantuml/plantuml.jar -o ./out ./tmp/*.txt
	cp ./tmp/out/*.png ./examples/

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean:
	rm -rf tmp

.PHONY: diagrams
diagrams: $(examples) gallery png

.PHONY: tmp
tmp:
	mkdir -p tmp
