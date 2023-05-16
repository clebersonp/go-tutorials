// the following command create this file: $ go mod init example.io/greetings
// The go mod init command creates a go.mod file to track your code's dependencies.
// So far, the file includes only the name of your module and the Go version your code supports.
// But as you add dependencies, the go.mod file will list the versions your code depends on.
// This keeps builds reproducible and gives you direct control over which module versions to use.
module create-module/example.io/greetings

go 1.20
