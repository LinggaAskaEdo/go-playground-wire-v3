.PHONY: build
build: 
	@wire ./src/cmd && \
		go mod tidy && \
		go generate ./src/cmd && \
		go build -o ./build/app ./src/cmd

.PHONY: vendor
vendor: generate 
		go mod vendor

.PHONY: run
run: 
	@air --build.cmd "go build -o ./build/app ./src/cmd" --build.bin "./build/app"

.PHONY: kafka-start
kafka-start: 
	@docker compose -f docker/kafka.yaml up

.PHONY: kafka-stop
kafka-stop: 
	@docker compose -f docker/kafka.yaml down
	