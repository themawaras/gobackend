APP=goback
APP_EXE="./build/$(APP)"

build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=darwin go build -o ${APP_EXE}

test:
	go test -cover -v ./...