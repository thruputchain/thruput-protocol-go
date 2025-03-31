package proto

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o generated.go -p proto protocol/coord.x protocol/discovery.x protocol/shard.x protocol/storage_item.x protocol/types.x 
