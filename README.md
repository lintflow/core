# core
grpc core

# Install

 ```
 $ go get github.com/lintflow/core
 $ cd $GOPATH/src/github.com/lintflow/core
 $ make install
 $ export PATH=$GOPATH/bin:$PATH
 $ go get github.com/ddollar/forego
 ```

# Run

## run services

```
 $ forego start

```

# Remote call from CLI

 ```
 $ lintflow services
services:
        id:"validator-1" address:"localhost:4545" name:"ctv" description:"core test validator" tags:"test" tags:"start" type:LINTER task_config:"{\"config\":\"json\"}"
        id:"resourser-1" address:"localhost:4535" name:"resourser" description:"core test resourser" tags:"test" tags:"start" task_config:"{\"config\":\"json\"}"
        id:"reporter-1" address:"localhost:4525" name:"reporter" description:"core test reporter" tags:"test" tags:"start" type:REPORTER task_config:"{\"config\":\"json\"}"

$ lintflow -repo https://github.com/grpc/grpc-go.git inspect
 73 / 72 [=====================================================================================================================================================] 101.39% 0
 see your report here - /tmp/report-2ff07704-bb38-4620-bd07-4c1c94949dcb.txt
 was finded 72 problems in 73 files.
 Finish!

# see report
$ cat /tmp/report-2ff07704-bb38-4620-bd07-4c1c94949dcb.txt

......

http://golang.org/doc/effective_go.html#mixed-caps
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/trace.go:34:1 - comments:should have a package comment, unless it's in another file for this package

https://golang.org/wiki/CodeReviewComments#package-comments
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/trace.go:104:9 - indent:if block ends with a return statement, so drop this else and outdent its block

https://golang.org/wiki/CodeReviewComments#indent-error-flow
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/transport/control.go:34:1 - comments:should have a package comment, unless it's in another file for this package

https://golang.org/wiki/CodeReviewComments#package-comments
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/transport/handler_server.go:38:1 - comments:should have a package comment, unless it's in another file for this package

https://golang.org/wiki/CodeReviewComments#package-comments
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/transport/http2_client.go:34:1 - comments:should have a package comment, unless it's in another file for this package

https://golang.org/wiki/CodeReviewComments#package-comments
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/transport/http2_server.go:34:1 - comments:should have a package comment, unless it's in another file for this package

https://golang.org/wiki/CodeReviewComments#package-comments
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/transport/http_util.go:34:1 - comments:should have a package comment, unless it's in another file for this package

https://golang.org/wiki/CodeReviewComments#package-comments
/tmp/86456488-66db-41db-b1b1-f7aacc1f5ffe/transport/transport.go:471:1 - comments:comment on exported var ErrConnClosing should be of the form "ErrConnClosing ..."

.....

 ```
