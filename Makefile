include .env
BOT=bot
SERVER=server

init:
	cd ${BOT} && go get .
	cd ${SERVER} && go get .

build.bot:
	cd ${BOT} && docker build -t go-bot .

build.server:
	cd ${SERVER} && docker build -t go-server .

build:
	make build.bot
	make build.server

update:
	-docker-compose down
	make build
	docker-compose up -d
