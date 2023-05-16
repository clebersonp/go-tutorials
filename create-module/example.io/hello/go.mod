module create-module/example.io/hello

go 1.20

// For production use, you’d publish the example.io/greetings module from its repository (with a module path that reflected its published location),
// where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.io/hello module
// so it can find the example.io/greetings code on your local file system.
// From the command prompt in the hello directory, run the following command: $ go mod edit -replace example.io/greetings=../greetings
replace create-module/example.io/greetings => ../greetings/

// From the command prompt in the hello directory, run the go mod tidy command to synchronize the example.io/hello module's dependencies,
// adding those required by the code, but not yet tracked in the module: $ go mod tidy

require create-module/example.io/greetings v0.0.0-00010101000000-000000000000
