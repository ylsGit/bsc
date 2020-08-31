package utils

const (
	Passwd            = "12345678"
	NetworkType       = "network-type"
	KeystorePath      = "keystore-path"
	ConfigPath        = "config-path"
	Operation         = "operation"
	BEP20ContractAddr = "bep20-contract-addr"
	LedgerAccount     = "ledger-account"

	InitKey        = "initKey"
	RefundRestBNB  = "refundRestBNB"
	DeployContract = "deployContract"
	ApproveBind    = "approveBindAndTransferOwnership"

	Mainnet = "mainnet"
	TestNet = "testnet"

	BindKeystore = "bind_keystore"

	TestnetRPC     = "https://data-seed-prebsc-1-s1.binance.org:8545"
	TestnetChainID = 97

	MainnnetRPC    = "https://bsc-dataseed1.binance.org:443"
	MainnetChainID = 56

	OneBNB          = 1e18
	DefaultGasPrice = 20000000000
	DefaultGasLimit = 4700000

	MainnetExplorerTxUrl = "%s: https://explorer.binance.org/smart/tx/%s"
	TestnetExplorerTxUrl = "%s: https://explorer.binance.org/smart-testnet/tx/%s"

	MainnetExplorerAddressUrl = "%s: https://explorer.binance.org/smart/address/%s"
	TestnetExplorerAddressUrl = "%s: https://explorer.binance.org/smart-testnet/address/%s"
)
