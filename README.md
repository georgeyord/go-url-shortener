# go-scrumpoker-api

A scrumpoker api implementation in Go lang #agile #scrum #poker #api #golang

## Run locally

First download dependencies:
```
make test_deps
make deps
```

Verify test are running successfully:
```
make test
```

Run `helloworld` command:
```
make run-helloworld-cmd
```


Build `helloworld` binary:
```
make build-helloworld-cmd
```

Then, you can run the binary:
```
./bin/helloworld
```

## Run with docker

First build:
```
docker-compose build helloworld
```

then, run:
```
docker-compose run helloworld
```