<p align="center"><img src="./assets/img/logo.png" width="350"></p>

[![Build Status](https://travis-ci.org/cyruzin/feelthemovies.svg?branch=master)](https://travis-ci.org/cyruzin/feelthemovies) [![Coverage Status](https://coveralls.io/repos/github/cyruzin/feelthemovies/badge.svg?branch=master&service=github)](https://coveralls.io/github/cyruzin/feelthemovies?branch=master&service=github) [![Go Report Card](https://goreportcard.com/badge/github.com/cyruzin/feelthemovies)](https://goreportcard.com/report/github.com/cyruzin/feelthemovies) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is the new API repository for Feel the Movies. Written in Go, totally open source.

This is my first Golang project, so probably there's a lot to improve. Any kind of help is welcome. I did my best researching the best practices, the best folder structuration and all that I could find. In the future, I pretend to bump this project to the version 2.0 using advanced concepts, like design patterns.

## Installation

Set the environment variables on the .env file.

Make sure that you have Docker installed and then run the following commands:

```sh
$ docker-compose up -d
$ go mod download 
```

That's it! Go to folder "cmd/feelthemovies" and run the command:

```sh
$ go run main.go
```

## Build

Run the command below and check the binary in the dist folder.

```sh
$ make build
```

## Packages

These are the packages that helped me build this project:

* [go-chi/chi](https://github.com/go-chi/chi) - chi is a lightweight, idiomatic and composable router for building Go HTTP services.
* [uber-go/zap](https://github.com/uber-go/zap) - Blazing fast, structured, leveled logging in Go.
* [InVisionApp/go-health](https://github.com/InVisionApp/go-health) - Library for enabling asynchronous health checks in your service.
* [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT).
* [go-playground/validator](https://github.com/go-playground/validator) - 
Package validator implements value validations for structs and individual fields based on tags.
* [google/uuid](https://github.com/google/uuid) - 
The uuid package generates and inspects UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
* [tome](https://github.com/cyruzin/tome) - Package tome was designed to paginate simple RESTful APIs.
* [envconfig](https://github.com/kelseyhightower/envconfig) - Golang library for managing configuration data from environment variables.
* [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - A MySQL-Driver for Go's database/sql package.
* [jmoiron/sqlx](https://github.com/jmoiron/sqlx) - sqlx is a library which provides a set of extensions on go's standard database/sql library.
* [go-redis](https://github.com/go-redis/redis) - Type-safe Redis client for Golang.

## Useful links

* [golang-standards / project-layout](https://github.com/golang-standards/project-layout) - Standard Go Project Layout.
* [avelino / awesome-go](https://github.com/avelino/awesome-go) - 
A curated list of awesome Go frameworks, libraries and software.
* [insomnia](https://insomnia.rest/download/) - Rest client.

## Contributing

To start contributing, please check [CONTRIBUTING](https://github.com/cyruzin/feelthemovies/blob/master/CONTRIBUTING.md).

## License

MIT
