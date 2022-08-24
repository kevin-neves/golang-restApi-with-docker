build: 
	cd cmd && go build -o ../out/golang-restApi-with-docker .

run:
	./out/golang-restApi-with-docker

docker:
	docker-compose up -d --build