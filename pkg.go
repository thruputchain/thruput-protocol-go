package proto

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o generated.go -p proto protocol/broadcast.x protocol/config.x protocol/contract.x protocol/coord.x protocol/db_read_interface.x protocol/discovery.x protocol/rpc_types.x protocol/shard.x protocol/storage_item.x protocol/transaction.x protocol/types.x 
