package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain/x/common"
)

// nolint
const (
	FeeTypeOrderNew     = "new"
	FeeTypeOrderCancel  = "cancel"
	FeeTypeOrderExpire  = "expire"
	FeeTypeOrderDeal    = "deal"
	FeeTypeOrderReceive = "receive"

	BuyOrder            = "BUY"
	SellOrder           = "SELL"
)

var TestTokenPair       = common.TestToken + "_" + sdk.DefaultBondDenom