package entity

import "time"

// Account ...
type Account struct {
	ID             string      `gorm:"primary_key" json:"accountId"`
	Name           string      `json:"name"`
	Email          string      `gorm:"index" json:"email"`
	DocumentNumber string      `json:"documentNumber,omitempty"`
	Password       string      `gorm:"index" json:"password,omitempty"`
	ExternalID     string      `gorm:"index" json:"accountExternalId,omitempty"`
	PhoneNumber    string      `json:"phoneNumber,omitempty"`
	VirtualCard    VirtualCard `gorm:"embedded" json:"virtualCard,omitempty"`
}

// CreateAccountRequest ...
type CreateAccountRequest struct {
	PsProductCode string `json:"psProductCode"`
	AccountOwner  struct {
		FullName               string `json:"fullName"`
		IdentityDocumentNumber string `json:"identityDocumentNumber"`
		ContactInformation     struct {
			PersonalPhoneNumber1 string `json:"personalPhoneNumber1"`
			Email                string `json:"email"`
		} `json:"contactInformation"`
	} `json:"accountOwner"`
	BillingAddress struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		State        string `json:"state"`
		Neighborhood string `json:"neighborhood"`
		Zipcode      string `json:"zipcode"`
	} `json:"billingAddress"`
	CardDeliveryAddress struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		State        string `json:"state"`
		Neighborhood string `json:"neighborhood"`
		Zipcode      string `json:"zipcode"`
	} `json:"cardDeliveryAddress"`
}

// CreateAccountResponse ...
type CreateAccountResponse struct {
	ResultData struct {
		ResultCode        int    `json:"resultCode"`
		ResultDescription string `json:"resultDescription"`
		PsResponseID      string `json:"psResponseId"`
	} `json:"resultData"`
	Account struct {
		AccountID       string `json:"accountId"`
		PsProductCode   string `json:"psProductCode"`
		PsProductName   string `json:"psProductName"`
		IssuerAccountID string `json:"issuerAccountId"`
		AccountOwner    struct {
			FullName                    string      `json:"fullName"`
			IdentityDocumentNumber      string      `json:"identityDocumentNumber"`
			OtherIdentityDocumentNumber interface{} `json:"otherIdentityDocumentNumber"`
			MotherName                  string      `json:"motherName"`
			ContactInformation          struct {
				PersonalPhoneNumber1  string      `json:"personalPhoneNumber1"`
				PersonalPhoneNumber2  interface{} `json:"personalPhoneNumber2"`
				ComercialPhoneNumber1 interface{} `json:"comercialPhoneNumber1"`
				Email                 string      `json:"email"`
			} `json:"contactInformation"`
			OccupationInfo interface{} `json:"occupationInfo"`
		} `json:"accountOwner"`
		Cards         []interface{} `json:"cards"`
		Status        string        `json:"status"`
		InclusionDate time.Time     `json:"inclusion_date"`
	} `json:"account"`
}
