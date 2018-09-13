# Finance Processor
An REST server built with Go to provide a POC with common finance functions

## Installation & Run

### Install Go.

To set up the GOPATH directory. first set up the following hierarchy

```bash
mkdir -p $HOME/go/bin $HOME/go/src
```

Then add these lines to your .profile
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
Then clone the repo to `$GOPATH/src`
```bash
$ cd $GOPATH/src
$ git clone <git-url>
$ cd finproc
```
### Install Dep

If you have Hombrew installed, run the following:

```
$ brew install dep
$ brew upgrade dep
```
otherwise
```bash
$ cd $GOPATH/bin
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

To get all the external packages, run

```bash
# Download the public dependencies
GOPATH=`pwd` go get -v -d ./...
```

### Run

Navigate to the project's root directory cd `$GOPATH/src/finprod` then, compile and run

```bash
# Build and Run
go build && ./finproc

# API Endpoint : http://127.0.0.1:8080
```
