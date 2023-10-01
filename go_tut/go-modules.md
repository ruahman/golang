a go module is a colleciton of packages in a folder with a go.mod file.
- go.mod
  + define a module path 
  + define dependencies 

first you need to define your directory as a go module,
which means give you module a name and set it's dependencies.
`go mod init <module-name>`

to bring in dependencies, just import them in your code
when you do go build or run it will add the imported lib to go.mod

go.sum is like a lock file

list module and version
`go list -m -versions github.com/pkg/errors`

get a versiono package
`go get <package>@<version>`

clean up the go.sum 
`go mod tidy`
- tidy the go.sum by removing old packages 
- also will update go.mod dependencies
