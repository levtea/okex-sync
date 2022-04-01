package okexsync

import (
	"github.com/okex/exchain/app/rpc/namespaces/eth"
	"github.com/okex/exchain/app/rpc/types"
)

type SyncRPCInfo struct {
	publicEthereumAPI *eth.PublicEthereumAPI
}

func NewSyncRPCInfo(publicEthereumAPI *eth.PublicEthereumAPI) *SyncRPCInfo {
	return &SyncRPCInfo{
		publicEthereumAPI: publicEthereumAPI,
	}
}

func (s *SyncRPCInfo) GetRPCBlockInfo(blockNum int64) (interface{}, error) {
	return s.publicEthereumAPI.GetBlockByNumber(types.BlockNumber(blockNum), false)
}
