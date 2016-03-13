# gohello
Hello, world example in GO. This basic primer assumes AMD 64 installation.

## 1. Installation ##

The *go* distribution is packaged as a tarball, which can be downloaded from https://golang.org/dl.  Extract the tarball into `/usr/local`. The content of the tarball typically contains `go` as the root level pathname, so in this primer, the extracted pathname is renamed and a symlink is created:

```
sudo tar -C /usr/local -xzf go1.6.linux-amd64.tar.gz
sudo mv /usr/local/go /usr/local/go16
sudo ln -s /usr/local/go16 /usr/local/go.current
```

After the tarball has been extracted, set the `GOROOT` environment variable and add `$GOROOT/bin` to the path

```
export GOROOT=/usr/local/go.current
export PATH=$GOROOT/bin:$PATH
```

## 2. Workspaces ##
 
The [GO Code Style](https://golang.org/doc/code.html) recommends that go repos are contained within a workspace container. Each of the workspaces' repositories map to the concept of a source code repo (e.g., a `github` repo). The structure of a repository is quite simple, containing `src`, `pkg` and `bin` pathnames. The go workspace is controlled by the `GOPATH` environment variable.

The following illustration creates the top level workapce, and then creates a github repository called `gohello`. 

```
mkdir -p $HOME/development/goworkspace
export GOPATH=$HOME/development/goworkspace
```


In this primer, a repository called `gohello` is created and added to github.com:

```
mkdir -p $GOPATH/src/github.com/foosmoo/gohello
cd $GOPATH/src/github.com/foosmoo/gohello
git init
```

Now create a simple //hello,world// example program:

```
cat << EOF > hello.go
package main
import "fmt"

func main() {
    fmt.Printf("Hello, world.\n")
}
EOF
```

Compile and run (go's''install'' command will locate ''hello.go'' with a fall-through search path that includes ''$GOROOT/src'' and ''$GOPATH/src''. The source code for the ''hello'' package is found because our package is located in ''$GOPATH/src/github.com/foosmoo/gohello'' (and ''GOPATH'' is set to ''$HOME/development/goworkspace''). This creates a bindary in ''$GOPATH/bin''

```
go install github.com/foosmoo/gohello
ls -laF $GOPATH/bin/hello
$GOPATH/bin/hello
Hello, world.
```

Add the file to the ''hello'' repo and commit:

```
$ git add hello.go
$ git commit -m "initial commit of hello.go"
```


## See Also ##

This README is extracted from:
   * [Google Go Install Guide](https://golang.org/doc/install)
   * [How to Write Go Code](https://golang.org/doc/code.html)

*Stuart Moorfoot* © **13 March 2016** foosmoo@gmail.com
