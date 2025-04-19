docker stop docker-demo
docker rm docker-demo
docker rmi docker-demo
docker rmi $(docker images -f "dangling=true" -q)
