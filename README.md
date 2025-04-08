## Server
```bash
git clone https://github.com/sadnemous/java-docker.git
cd java-server
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
