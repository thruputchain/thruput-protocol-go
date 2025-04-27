package proto

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o generated.go -p proto protocol/broadcast.x protocol/config.x protocol/contract.x protocol/coord.x protocol/discovery.x protocol/executor_to_db_rpc.x protocol/ffi_enums.x protocol/mempool.x protocol/rpc_types.x protocol/shard.x protocol/solo.x protocol/storage_item.x protocol/transaction.x protocol/types.x 
