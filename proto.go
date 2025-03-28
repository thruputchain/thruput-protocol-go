package proto

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o types.go -p proto types.x
//go:generate goxdr -o coord.go -p proto coord.x
