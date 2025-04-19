## To run the Spring-Boot App in Container
### which means run our microservice
```bash
git clone https://github.com/sadnemous/java-docker.git
cd java-server
./gradlew clean build
docker stop docker-demo .
docker rm docker-demo
docker rmi docker-demo
docker rmi $(docker images -f "dangling=true" -q)

docker build -t docker-demo .
docker run -d --name docker-demo -p 8080:8080 docker-demo
```

## Client
```bash
cd ../golang-client
sudo apt install gojq
sudo apt install python3-demjson

./run.sh 
or
go run client.go|cut -d: -f2-|gojq

```
