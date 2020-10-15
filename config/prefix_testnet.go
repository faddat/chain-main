// +build testnet

package config

const (
	CoinType       = 1
	FundraiserPath = "44'/1'/0'/0/1"
)

var (
	AccountAddressPrefix   = "tcro"
	AccountPubKeyPrefix    = "tcropub"
	ValidatorAddressPrefix = "tcrocncl"
	ValidatorPubKeyPrefix  = "tcrocnclpub"
	ConsNodeAddressPrefix  = "tcrocnclcons"
	ConsNodePubKeyPrefix   = "tcrocnclconspub"
	HumanCoinUnit          = "tcro"
	BaseCoinUnit           = "basetcro" // 10^-8 AKA "carson"
	CroExponent            = 8
	CoinToBaseUnitMuls     = map[string]uint64{
		"tcro": 1_0000_0000,
	}
)