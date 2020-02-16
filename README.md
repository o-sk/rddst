# rddst
`rddst` is a command to get a redirect destination.

## Installation

Use `go get` to install.

```console
$ go get -u github.com/o-sk/rddst/cmd/rddst
```
## Usage

```console
$ rddst <url>
```

### Example
```console
$ rddst http://example.com
The url is not redirect

$ rddst https://google.com
https://www.google.com/
```
