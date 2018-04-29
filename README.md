# GO-IMGUR-CLI



Go-IMGUR-CLI is a CLI based uploader to imgur.



### Installation:

Go-IMGUR-CLI requires [Golang Runtime](https://golang.org/) to run.

Assuming you have the correct path specified in your $GOPATH.


Clone the repo: `git@github.com:whaangbuu/go-imgur-cli.git`

```sh
$ cd go-imgur-cli
$ go install
$ go-imgur-cli <file/to/upload>
```
NOTE: By running `go install`, our app is now a global command that you can access anywhere in your workspace.

### Todos:
 - [ ] Create a more reusable `uploader` struct.
 - [ ] Add a loader when uploading.
 - [ ] Write test
 

License
----

MIT
