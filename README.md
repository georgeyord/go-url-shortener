# go-url-shortener

A url shortener includes 3 components:

- a redirection module that takes a short url as input, retrieves long url and redirects to that
- an API to manage pairs of short/long urls
- a cli to use the API from the command line

> The redirection module and the API co-exist under one single web server

# Development

## Prepare local running

First download dependencies:
```
make test-deps
make deps
```

Verify test are running successfully:
```
make test
```

## Run the api server locally


Run `url-shortener-api` command:
```
make run-url-shortener-api
```


Build `url-shortener-api` binary:
```
make build-url-shortener-api
```

Then, you can run the binary:
```
./bin/url-shortener-api
```


## Run the cli locally

Run `url-shortener` command:
```
make run-url-shortener-cli
```


Build `url-shortener` binary:
```
make build-url-shortener-cli
```

Then, you can run the binary:
```
./bin/url-shortener-cli
```

## Run the web server with docker

First build:
```
docker-compose build url-shortener-api
```

then, run:
```
docker-compose run url-shortener-api
```

# Business requirements

## Redirection module

- Can read the short url - Done
- Can check if the short url exists
- Can return a 404 (Not Found) HTTP code if short url does not exist - Done
- Can retrieve the long url from the short one - Done
- Can redirect to long url using a 301 (Moved Permanently) HTTP code - Done
- Can redirect to long url using a custom HTTP code
- Can redirect to long url honoring the incoming query params
- Can keep stats of the short url usage (using Kafka)

## API for admininstartion

- Can create a pair of short/long urls - Done
- Can check if a short url exists
- Can modify a pair of short/long urls - Done
- Can delete a pair of short/long urls
- Can generate a random short url (generation process should be abstracted) - Done
- Can have multiple generation options and params (valid characters, length)
- Can store a pair of short/long urls (storage process should be abstracted) - Done
- Can have multiple storage options (file, key-value db, traditional db etc.)
- Can run behind an authentication wall

## CLI for admininstartion

- Can create a pair of short/long urls - Done
- Can check if a short url exists
- Can modify a pair of short/long urls - Done
- Can delete a pair of short/long urls
- Can use authentication to access the API

# Technical requirements

- Can use web and cli as separate docker images

# Implementation tracking

Here is the long list...

First some technical requirements:

- Add a configuration management system - Done - Sep 2019
- Use [gin](https://github.com/gin-gonic/gin) as a web framework or [chi](https://github.com/go-chi/chi)/[gorilla/mux](https://github.com/gorilla/mux) as router - Done - Sep 2019
- Use [Gin Binding](https://github.com/gin-gonic/gin#bind-query-string-or-post-data) to fill model - Done - Feb 2020
- Add [HTTP tests](https://github.com/gin-gonic/gin#testing) - Done - Mar 2020
- Use Godoc to document usage
- Http Error handling - Done - Mar 2020
- Use custom logger
- Use [Corba](https://github.com/spf13/cobra) for cli - Done - Feb 2020
- Use Enum for valid application environments

# Thank you...

> API is heavily influenced by [informatics-lab/url-shortener](https://github.com/informatics-lab/url-shortener)
