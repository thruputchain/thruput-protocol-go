package proto

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o generated.go -p proto protocol/batch.x protocol/broadcast.x protocol/catchup.x protocol/config.x protocol/contract.x protocol/coord.x protocol/db_coord.x protocol/discovery.x protocol/exec_coord.x protocol/executor_to_db_rpc.x protocol/ffi_enums.x protocol/multiplex.x protocol/proposal.x protocol/retrocausal_queue.x protocol/rpc_types.x protocol/shard.x protocol/solo.x protocol/spectra.x protocol/storage_item.x protocol/storage_lock.x protocol/transaction.x protocol/types.x protocol/wallet.x 
