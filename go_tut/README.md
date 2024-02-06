
# go run 
`go run ./main.go`

# go test 
`go test -v <package-name>`
- v: is for verbose

# package 

a package is like a __project__, or __workspace__.

`package main` is making an executable package.

`package <name>` is making a reusable package that can be imported and used as a dependenciy.

# module #
in order to include third parties you need to define your project as a module
`go mod init <name-of-your-module`

*go.mod*
- is where you define the name of your module and the third parties you include in your module

to include a third party 
`go get <repo-path>`
- default is github
ex: `go get rsc.io/quote`

*go.sum* 
- checksum of all third parties and dependencies

you can now import the package in your code

`go mod tidy`
this will match the go.mod file with the dependencies you are using in you.
- it will download the dependencies that are missing 
- and will remove the ones that are no longer used in your code
