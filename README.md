Error message:
```
$ gomobile bind -v -target ios .
type-checking package "." failed (/Users/username/Dev/st/example/main.go:6:2: could not import github.com/google/netstack/tcpip (type-checking package "github.com/google/netstack/tcpip" failed (/Users/username/.gvm/pkgsets/go1.11/global/pkg/mod/github.com/petethepig/netstack@v0.0.0-20171026205909-b4b77f7e31f6/tcpip/tcpip.go:25:2: could not import github.com/google/netstack/tcpip/buffer (cannot find package "github.com/google/netstack/tcpip/buffer" in any of:
	/Users/username/.gvm/gos/go1.11/src/github.com/google/netstack/tcpip/buffer (from $GOROOT)
	/Users/username/.gvm/pkgsets/go1.11/global/src/github.com/google/netstack/tcpip/buffer (from $GOPATH)))))

gomobile: /Users/username/.gvm/pkgsets/go1.11/global/bin/gobind -lang=go,objc -outdir=/var/folders/zg/ssp6j1213dg2nd1l8wmc1xs40000gn/T/gomobile-work-114613772 -tags=ios . failed: exit status 1
```
