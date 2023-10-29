NAME=myCut

${NAME}: build

clean:
	rm ${NAME}

build:
	go build -o ${NAME}

rebuild: clean build

test: ${NAME}
	go test -v -short -race -count=1 ./...