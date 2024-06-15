.PHONY: build
build: 
	@wire ./src/cmd && \
		go mod tidy && \
		go generate ./src/cmd && \
		go build -o ./build/app ./src/cmd

.PHONY: build-cgo
build-cgo: 
	@wire ./src/cmd && \
		go mod tidy && \
		go generate ./src/cmd && \
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/app ./src/cmd/

.PHONY: run
run: 
	@air --build.cmd "go build -o ./build/app ./src/cmd" --build.bin "./build/app"

.PHONY: db-start
db-start: 
	@docker start mysql-docker postgres-docker redis-docker

.PHONY: db-stop
db-stop: 
	@docker stop mysql-docker postgres-docker redis-docker

.PHONY: queue-start
queue-start: 
	@docker start rabbitmq-docker && \
		docker compose -f docker/kafka.yaml up

.PHONY: queue-stop
queue-stop: 
	@docker stop rabbitmq-docker && \
		docker compose -f docker/kafka.yaml down

.PHONY: rabbit-start
rabbit-start: 
	@docker start rabbitmq-docker

.PHONY: rabbit-stop
rabbit-stop: 
	@docker stop rabbitmq-docker

.PHONY: kafka-start
kafka-start: 
	@docker compose -f docker/kafka.yaml up

.PHONY: kafka-stop
kafka-stop: 
	@docker compose -f docker/kafka.yaml down
	
# .PHONY: docker-compose-build
# docker-compose-build:
# 	@docker compose -f docker-compose.yml build 

# .PHONY: docker-compose-up
# docker-compose-up:
# 	@docker compose -f docker-compose.yml up 	

# .PHONY: docker-compose-down
# docker-compose-down:
# 	@docker compose -f docker-compose.yml down 

# .PHONY: docker-build
# docker-build:
# 	# @docker build --tag linggaaskaedo/go-playground-wire-v3-docker --build-arg PORT=6666 -f docker/app/Dockerfile .
# 	@docker build --tag linggaaskaedo/go-playground-wire-v3-docker --build-arg PORT=6666 -f Dockerfile .

# .PHONY: docker-build-multistage
# docker-build-multistage:
# 	@docker build --tag linggaaskaedo/go-playground-wire-v3-docker:multistage --build-arg PORT=6666 -f docker/app/Dockerfile.multistage .

# .PHONY: docker-build-scratch
# docker-build-scratch:
# 	@docker build --tag linggaaskaedo/go-playground-wire-v3-docker:scratch --build-arg PORT=6666 -f docker/app/Dockerfile.scratch .	