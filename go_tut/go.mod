// set your module with `go mod init <name-of-your-module>`

// a module is just a go project

// workspace -> modules -> packages

// define name of our module and the go version we are using in this module
module go_tut

go 1.18

// go mod tidy,  to tidy this up
require github.com/mattn/go-sqlite3 v1.14.22
