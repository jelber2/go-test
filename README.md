# go-test

testing go with python on Ubuntu linux 20.04

## first install go

    $ go --version
    go version go1.15.8 linux/amd64

## install python3

    $ python3 --version
    Python 3.8.10

## install gopy (https://github.com/go-python/gopy)

    $ python3 -m pip install pybindgen
    $ go get golang.org/x/tools/cmd/goimports
    $ go get github.com/go-python/gopy
    
## compiling example to shared object

example is https://github.com/jelber2/go-test/blob/main/bktree/bktree.go

### get go modules for the example

    $ go get github.com/creasty/go-levenshtein
    $ go get github.com/agatan/bktree
    
    # this part is a little weird, but apparently
    # you need to have the module you want to compile
    # in your GOPATH locally on your system first?
    $ go get github.com/jelber2/go-test/bktree
    
    # I think you could also just copy the .go file to a directory
    # such as 
    $ mkdir -p /home/jelbers/go/src/github.com/jelber2/go-test/bktree2/ 
    $ cp bktree.go /home/jelbers/go/src/github.com/jelber2/go-test/bktree2/.
    
### run gopy

    $ gopy build -vm="python3" -output test github.com/jelber2/go-test/bktree/

    --- Processing package: github.com/jelber2/go-test/bktree ---

    --- building package ---
    gopy build -vm=python3 -output test github.com/jelber2/go-test/bktree/
    goimports -w bktree.go
    go build -buildmode=c-shared -o bktree_go.so .
    /usr/bin/python3 build.py
    CGO_CFLAGS=-I/usr/include/python3.8 -fPIC -Ofast
    CGO_LDFLAGS=-L/usr/lib -lpython3.8 -lcrypt -lpthread -ldl -lutil -lm -lm
    go build -buildmode=c-shared -o _bktree.cpython-38-x86_64-linux-gnu.so .
    
    $ ls -1 test
    bktree.c
    _bktree.cpython-38-x86_64-linux-gnu.h
    _bktree.cpython-38-x86_64-linux-gnu.so
    bktree.go
    bktree_go.h
    bktree.py
    build.py
    go.py
    __init__.py
    Makefile
    
### import module then function of interest with python3

    $ cd test/
    $ python3
    $ import _bktree
    $ _bktree.bktree_Add("yes")
    Input is AAAA. 
    Did you mean:
    AAAA (distance: 0)
    AACC (distance: 2)
    AACT (distance: 2)
    AAGC (distance: 2)
    AACG (distance: 2)
    AAAG (distance: 1)
    AAAT (distance: 1)
    AAAC (distance: 1)
