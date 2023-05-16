// To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path.
// For example, if you’re developing in a stringtools directory, your temporary module path might be <company-name>/stringtools
// https://go.dev/doc/modules/gomod-ref
module getting-started/example.io/hello-world

go 1.20

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
