package raw

import (
	"context"
	"fmt"

	neutrontypes "github.com/neutron-org/neutron/x/interchainqueries/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client/http"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

const websocketPath = "/websocket"

// Subscribe subscribes to a blockchain by the given rpcAddress using websocket and returns a chan
// of events.
//
// WARNING: rpcclient.Subscribe from tendermint can fail to work with some blockchain versions of
// tendermint.
func Subscribe(ctx context.Context, subscriberName string, rpcAddress string, query string) (<-chan coretypes.ResultEvent, error) {
	httpclient, err := rpcclient.New(rpcAddress, websocketPath)
	if err != nil {
		return nil, fmt.Errorf("could not create new rpcclient: %w", err)
	}
	err = httpclient.Start()
	if err != nil {
		return nil, fmt.Errorf("could not start httpclient when subscribing to target chain: %w", err)
	}

	response, err := httpclient.Subscribe(ctx, subscriberName, query)
	if err != nil {
		return nil, fmt.Errorf("could not subscribe to target chain: %w", err)
	}
	return response, nil
}

// SubscribeQuery describes query to filter out events by module and action.
func SubscribeQuery(zoneId string) string {
	// TODO: add zoneId to the query when zone id passing is implemented
	return fmt.Sprintf("message.module='%s' AND message.action='%s'", neutrontypes.ModuleName, neutrontypes.AttributeValueQuery)
}
