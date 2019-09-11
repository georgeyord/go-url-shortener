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

# Next steps

Here is the long list...

First some technical requirements:

- Add a configuration management system - Done - Sep 2019
- Http Router
- Http Error handling
- Use custom logger
- Use [Corba](https://github.com/spf13/cobra) for cli

And now some business  requirements:

- Define the scrumpoker entities
- Create a Room as Administrator with a specific (optional) name
- Room should create a custom identifier, Room Id
- Join a Room as someone
- Set a Story for evaluation
- Start voting
- Do NOT show results until everybody votes
- Show how many Votersremian to vote
- Find the min/max voters and let them fight
- Re-run a voting
- Have a customizable Voting system (Fibinacci, custom...)
