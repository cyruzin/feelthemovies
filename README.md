<p align="center"><img src="./assets/img/logo.png" width="350"></p>

[![Build Status](https://travis-ci.org/cyruzin/feelthemovies.svg?branch=master)](https://travis-ci.org/cyruzin/feelthemovies) [![Coverage Status](https://coveralls.io/repos/github/cyruzin/feelthemovies/badge.svg?branch=master&service=github)](https://coveralls.io/github/cyruzin/feelthemovies?branch=master&service=github) [![Go Report Card](https://goreportcard.com/badge/github.com/cyruzin/feelthemovies)](https://goreportcard.com/report/github.com/cyruzin/feelthemovies) [![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

This is the new API repository for Feel the Movies. Written in Go, totally open source.

## App

Currently available for Android only. I have plans for an iOS version, since the application was written in React Native, it should be no problem to release a version in the future.

[Download it on Play Store](https://play.google.com/store/apps/details?id=br.com.feelthemovies&hl=en_ZA)

## Installation

Change the values of environment variables in the .env file if desired. 

Make sure that you have Docker and Docker Compose installed and then run the following command:

```sh
$ docker-compose up -d
```

That's it! Wait for the API container to be ready and check url: **[localhost:8000/v1/recommendations](http://localhost:8000/v1/recommendations)**.

To recompile, run:

```sh
$ docker-compose up -d --build
```

## Build

Run the command below and check the binary in the dist folder.

```sh
$ make build
```

## Contributing

To start contributing, please check [CONTRIBUTING](https://github.com/cyruzin/feelthemovies/blob/master/CONTRIBUTING.md).

## License

MIT
