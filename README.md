# Feel the Movies

This is the new API repository for Feel the Movies. Written in Go, totally open source, with test coverage.

This is my first Golang project, so probably there's a lot to improve. Any kind of help is welcome. I did my best researching the best practices, the best folder structuration and all that I could find. In the future, I pretend to bump this project to the version 2.0 using advanced concepts, like design patterns.

## Installation

- Create a database on your MySQL and import the .sql file that are inside of "db" folder.
- Rename the file ".env.example" to ".env" and set your database configuration.
- Inside the folder "api" there's a Insomnia's schema that I created to help with the tests. Download it and import it.

That's it! Go to folder "cmd/feelthemovies" and run the command:
```sh
$ go run main.go
```

### Packages

These are the packages that helped me build this project:

* [gorilla/mux](https://github.com/gorilla/mux) - Package mux implements a request router and dispatcher.
* [go-playground/validator](https://github.com/go-playground/validator) - 
Package validator implements value validations for structs and individual fields based on tags.
* [rs/cors](https://github.com/rs/cors) - 
CORS is a net/http handler implementing Cross Origin Resource Sharing W3 specification in Golang.
* [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - A MySQL-Driver for Go's database/sql package.

## Useful links

* [golang-standards / project-layout](https://github.com/golang-standards/project-layout) - Standard Go Project Layout.
* [avelino / awesome-go](https://github.com/avelino/awesome-go) - 
A curated list of awesome Go frameworks, libraries and software.
* [insomnia](https://insomnia.rest/download/) - Rest client.

### Contributing

To start contributing, please check [CONTRIBUTING](https://github.com/cyruzin/feelthemovies/blob/master/CONTRIBUTING.md).

### License

MIT