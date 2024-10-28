hello:
	echo "hello, world!"

build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o memo-app-api/main src/web.go

start: build
	cd memo-app-api && docker-compose up --build

deploy: build
	eval "$(ssh-agent -s)"
	ssh-add ~/.ssh/gcp_tamurakeito_key
	rsync -avz memo-app-api/ tamurakeito@34.146.93.87:/home/tamurakeito/memo-app-api
	rm -r memo-app-api/main

ssh:
	ssh -i ~/.ssh/gcp_tamurakeito_key tamurakeito@34.146.93.87
