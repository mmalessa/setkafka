# Kafka topic configurator
Allows you to create and delete topics in Kafka and copy messages from one topic to another.

**I share as is. Use at your own risk!**

## How to...

### Build binary
```sh
make up
make binary
```

### Develop
```sh
make up
make shell
```
#### Visual Studio Code
- Install (Ctrl + Shift + X): Go (Go Team at Google)
- On left bottom corner click >< icon and select Attach to running container... and select setkafka container
- Run command (Ctrl + Shift + P) Go: Install/Update tools, select all and click OK

### Use
Customize the config.yaml file to suit your needs

```sh
./setkafka -h
./setkafka topic -h
./setkafka topic list -h
./setkafka topic create -h
./setkafka topic delete -h
./setkafka topic copy -h
```

## DockerHub notes 
Just for me, so I don't forget
```sh
make build-prod
docker image ls | grep setkafka
docker tag setkafka:latest mmalessa/setkafka:0.0.1
docker push mmalessa/setkafka:0.0.1
```