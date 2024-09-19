# Kafka topic configurator
Allows you to create and delete topics in Kafka and copy messages from one topic to another.

**I share as is. Use at your own risk!**

## How to...

### Build binary
```sh
make binary
```

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


## Dev notes

### Environment
```sh
. envexport
./bin/setkafka -c ./setkafka.yaml
```

### Visual Studio Code

    On left bottom corner click >< icon and select Attach to running container... and select container $(APP_NAME)
    Install (Ctrl + Shift + X):
        Go (Go Team at Google)
    Run command (Ctrl + Shift + P) Go: Install/Update tools, select all and click OK
