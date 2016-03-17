# util
Example GOLang library. As with other examples in this GOLang primer, the target platform is `linux_amd64`.

## Build
The `util.go` package contains an example function called `AsUpper`. The function is exported because it begins with an uppercase letter. This is one of GOLang's **programming by convention** idioms, where variables and functions are scoped based on the case of the first letter of their identifier. E.g., `AsUpper` is globally scoped whereas `asUpper` is package scoped (inaccessible from outside of the local package).

To build the example library, execute:

```
go install github.com/foosmoo/util
```

The above produces `$GOPATH/pkg/linux_amd64/github.com/foosmoo/util.a` in 

## Test

The GOLang compiler supports test constructs, again by **convention**. The test execution system looks for the test harness matching the `<name>_test`, where `<name>` is the function name. In our example, the `AsUpper` function has a corresponding test harness, which is `AsUpper_test`.

The test harnesses are executed using the `test` command of the `go` compiler.

```
go test
```

## Run

To execute the main driver, run the `hello` binary, produced by the build step above:

```
make run
```


*Stuart Moorfoot* © **17 March 2016** foosmoo@gmail.com
