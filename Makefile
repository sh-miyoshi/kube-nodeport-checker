APP_NAME=kube-nodeport-checker
APP_VERSION=latest

go_build:
	go build -o $(APP_NAME) main.go

docker_build:
	docker build -t $(APP_NAME):$(APP_VERSION) --build-arg http_proxy="$(http_proxy)" --build-arg https_proxy="$(https_proxy)" .
