mongoUp:
	cd ./build && docker-compose up -d mongo mongo-express

mongoDown:
	cd ./build && docker-compose down mongo mongo-express
