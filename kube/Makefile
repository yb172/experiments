pwd := $(shell pwd)

gen-proto:
	# Gateway
	rm -rf ./gateway/proto
	mkdir -p ./gateway/proto
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I ./proto ./proto/wordgen/*.proto --go_out=plugins=grpc:./gateway/proto/

	# Internets
	rm -rf ./internets/proto
	mkdir -p ./internets/proto
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I ./proto ./proto/wordgen/*.proto --go_out=plugins=grpc:./internets/proto/

	# Number
	rm -rf ./number/proto
	mkdir -p ./number/proto
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I ./proto ./proto/wordgen/*.proto --go_out=plugins=grpc:./number/proto/

	# Word
	rm -rf ./word/proto
	mkdir -p ./word/proto
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I ./proto ./proto/wordgen/*.proto --go_out=plugins=grpc:./word/proto/
