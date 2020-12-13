package entity

// PurchaseAdd ...
type PurchaseAdd struct {
	PurchaseID               string `json:"purchase_id"`
	AccountID                string `json:"account_id"`
	PsProductCode            string `json:"psProductCode"`
	PsProductName            string `json:"psProductName"`
	CountryCode              string `json:"countryCode"`
	Source                   string `json:"source"`
	CallingSystemName        string `json:"callingSystemName"`
	PreAuthorization         bool   `json:"preAuthorization"`
	IncrementalAuthorization bool   `json:"incrementalAuthorization"`
	Authorization            struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"authorization"`
	Brand string `json:"brand"`
	Card  struct {
		PaysmartID string `json:"paysmart_id"`
		IssuerID   int    `json:"issuer_id"`
		Pan        string `json:"pan"`
		Panseq     string `json:"panseq"`
	} `json:"card"`
	TotalAmount struct {
		Amount       int `json:"amount"`
		CurrencyCode int `json:"currency_code"`
	} `json:"total_amount"`
	DollarAmount   interface{} `json:"dollar_amount"`
	OriginalAmount struct {
		Amount       int `json:"amount"`
		CurrencyCode int `json:"currency_code"`
	} `json:"original_amount"`
	DollarRealRate interface{} `json:"dollar_real_rate"`
	Spread         interface{} `json:"spread"`
	EntryMode      string      `json:"entry_mode"`
	ProcessingCode struct {
		TipoTransacao          string `json:"tipo_transacao"`
		SourceAccountType      string `json:"source_account_type"`
		DestinationAccountType string `json:"destination_account_type"`
	} `json:"processing_code"`
	HolderValidationMode string        `json:"holder_validation_mode"`
	Fees                 []interface{} `json:"fees"`
	Establishment        struct {
		Mcc     int         `json:"mcc"`
		Name    string      `json:"name"`
		City    string      `json:"city"`
		Address string      `json:"address"`
		Zipcode string      `json:"zipcode"`
		Country string      `json:"country"`
		Cnpj    interface{} `json:"cnpj"`
	} `json:"establishment"`
	Internacional   bool `json:"internacional"`
	OriginalIso8583 struct {
		Mti   string      `json:"mti"`
		De002 string      `json:"de002"`
		De003 string      `json:"de003"`
		De004 string      `json:"de004"`
		De005 interface{} `json:"de005"`
		De006 interface{} `json:"de006"`
		De007 string      `json:"de007"`
		De008 interface{} `json:"de008"`
		De009 interface{} `json:"de009"`
		De010 interface{} `json:"de010"`
		De011 string      `json:"de011"`
		De012 string      `json:"de012"`
		De013 string      `json:"de013"`
		De014 string      `json:"de014"`
		De015 interface{} `json:"de015"`
		De016 interface{} `json:"de016"`
		De018 string      `json:"de018"`
		De019 string      `json:"de019"`
		De022 string      `json:"de022"`
		De023 string      `json:"de023"`
		De024 string      `json:"de024"`
		De025 interface{} `json:"de025"`
		De026 interface{} `json:"de026"`
		De028 interface{} `json:"de028"`
		De029 interface{} `json:"de029"`
		De032 string      `json:"de032"`
		De033 interface{} `json:"de033"`
		De035 string      `json:"de035"`
		De036 interface{} `json:"de036"`
		De037 string      `json:"de037"`
		De038 interface{} `json:"de038"`
		De039 interface{} `json:"de039"`
		De041 string      `json:"de041"`
		De042 string      `json:"de042"`
		De043 string      `json:"de043"`
		De045 interface{} `json:"de045"`
		De046 interface{} `json:"de046"`
		De047 interface{} `json:"de047"`
		De048 string      `json:"de048"`
		De049 string      `json:"de049"`
		De050 interface{} `json:"de050"`
		De051 interface{} `json:"de051"`
		De052 string      `json:"de052"`
		De053 interface{} `json:"de053"`
		De054 interface{} `json:"de054"`
		De055 string      `json:"de055"`
		De056 interface{} `json:"de056"`
		De058 string      `json:"de058"`
		De059 interface{} `json:"de059"`
		De060 string      `json:"de060"`
		De062 interface{} `json:"de062"`
		De063 interface{} `json:"de063"`
		De090 interface{} `json:"de090"`
		De095 interface{} `json:"de095"`
		De105 interface{} `json:"de105"`
		De106 interface{} `json:"de106"`
		De107 interface{} `json:"de107"`
		De121 interface{} `json:"de121"`
		De122 interface{} `json:"de122"`
		De123 interface{} `json:"de123"`
		De124 interface{} `json:"de124"`
		De125 interface{} `json:"de125"`
		De126 interface{} `json:"de126"`
		De127 int         `json:"de127"`
	} `json:"original_iso8583"`
	ForceAccept bool `json:"forceAccept"`
	FraudData   struct {
		CreditorFraudScore          int    `json:"creditorFraudScore"`
		EloBrandFraudScore          int    `json:"eloBrandFraudScore"`
		FraudScorePrimaryReason     int    `json:"fraudScorePrimaryReason"`
		FraudScoreSecondaryReason   int    `json:"fraudScoreSecondaryReason"`
		FraudScoreTertiaryReason    int    `json:"fraudScoreTertiaryReason"`
		FraudDecisionRecommendation string `json:"fraudDecisionRecommendation"`
		ScoreOriginIndicator        int    `json:"scoreOriginIndicator"`
	} `json:"fraudData"`
}

// PurchaseAddResponse ...
type PurchaseAddResponse struct {
	Message         string `json:"message"`
	Code            int    `json:"code"`
	AuthorizationID int    `json:"authorization_id"`
	Balance         struct {
		Amount       int `json:"amount"`
		CurrencyCode int `json:"currency_code"`
	} `json:"balance"`
}
