package enums

const
(
	None = 0
	InvalidOrExpiredGameProviderToken = 10
	InvalidOrExpiredPlayerToken = 11
	InvalidOrExpiredBrandToken = 12
	InvalidOrExpiredResellerToken = 16
	InvalidOrExpiredJwtToken = 17
	InvalidTokenScope = 19
	UserBlocked = 20
	CountryRestricted = 21
	InvalidCredentials = 30
	CurrencyMismatch = 50
	InvalidAmount = 60
	InsufficientFunds = 100
	TransactionIdNotFound = 200
	DuplicateTransaction = 210
	InvalidGameAction = 250
	InvalidArguments = 300
	MissingToken = 400
	IncorrectFormat = 500
	PlayerNotFound = 600
	CreateCampaignFail = 610
	UnderMaintenance = 700
	ConfiguredTimeoutExceeded = 800
	ConfiguredGameProviderTimeoutExceeded = 801
	SystemError = 900
	ConnectionError = 901
	OperationNotAllowed = 902
	TestPlayerStakeAmountCapExceeded = 904
	BrowserIncompatibility = 905
	PlatformMismatched = 906
)