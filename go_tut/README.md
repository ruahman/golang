
# go run 
`go run ./main.go`

# go test 
`go test -v <package-name>`
- v: is for verbose

# package 

a package is like a **project**, or **workspace**.

every go file has the package keyword on top

`package main` this marks the file as the main point of entry.
- it's important to note that each package can have a main.go.  
  This is by design, so that a project can have multiple points of entries through 
  their packages.  golang package system lends it self to a micro service architecture.
- go projects can have multiple packages in it and multiple entry points if it wants.
  

`package <name>` is making a reusable package that can be imported and used as a dependenciy.

# module #
in order to include third parties you need to define your project as a module
`go mod init <name-of-your-module>`
- this will be the first thing you type when starting a go project

*go.mod*
- is where you define the name of your module, version of go and the third parties you include in your module
- kinda like package.json for node.js

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
