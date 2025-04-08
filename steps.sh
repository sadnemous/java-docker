./gradlew clean build
docker stop docker-demo .
docker rm docker-demo
docker rmi docker-demo
docker rmi $(docker images -f "dangling=true" -q)

docker build -t docker-demo .
docker run -d --name docker-demo -p 8080:8080 docker-demo


