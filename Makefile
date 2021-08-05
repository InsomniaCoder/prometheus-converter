BINARY=prometheus-converter
test: 
	go test -v -cover -covermode=atomic ./...

local:
	go run app/main.go

lint:
	go fmt ./...

prometheus-converter:
	go build -o ${BINARY} app/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t ${BINARY} .

run:
	docker-compose up --build -d

stop:
	docker-compose down