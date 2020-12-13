package entity

// VirtualCard ...
type VirtualCard struct {
	VCardID     string `json:"vCardId"`
	VPan        string `json:"vPan"`
	VCvv        string `json:"vCvv"`
	VDateExp    string `json:"vDateExp"`
	VCardholder string `json:"vCardholder"`
}

// VirtualCardRequest ...
type VirtualCardRequest struct {
	BirthDate   string `json:"birthDate"`
	Constraints struct {
		CurrencyCode        int    `json:"currency_code"`
		MaxAmount           string `json:"max_amount"`
		UsageLimit          string `json:"usage_limit"`
		ExpirationTimestamp string `json:"expiration_timestamp"`
	} `json:"constraints"`
}

// VirtualCardResponse ...
type VirtualCardResponse struct {
	ResultData struct {
		ResultCode        int    `json:"resultCode"`
		ResultDescription string `json:"resultDescription"`
		PsResponseID      string `json:"psResponseId"`
	} `json:"resultData"`
	VirtualCard *VirtualCard `json:"virtualCard"`
}
