package entity

// PurchaseAdd ...
type PurchaseAdd struct {
	PurchaseID    string `json:"purchase_id"`
	AccountID     string `json:"account_id"`
	PsProductCode string `json:"psProductCode"`
	PsProductName string `json:"psProductName"`
	TotalAmount   struct {
		Amount       int `json:"amount"`
		CurrencyCode int `json:"currency_code"`
	} `json:"total_amount"`
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
