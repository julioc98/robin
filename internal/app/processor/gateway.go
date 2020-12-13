package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julioc98/robin/internal/app/entity"
	"github.com/julioc98/robin/pkg/client"
)

const (
	createAccountPath     = "/accounts"
	psProductCodePF       = "083001"
	createVirtualCardPath = "/virtualcards"
)

type paytGateway struct {
	baseURL string
	client  client.Requester
}

//NewPayGateway create new payt repository
func NewPayGateway(baseURL string, client client.Requester) Gateway {
	return &paytGateway{
		baseURL: baseURL,
		client:  client,
	}
}

// CreateAccount ...
func (g *paytGateway) CreateAccount(account *entity.Account) (string, error) {
	uri := fmt.Sprint(g.baseURL, createAccountPath)

	createAccountRequest := entity.CreateAccountRequest{
		PsProductCode: psProductCodePF,
	}

	createAccountRequest.AccountOwner.FullName = account.Name
	createAccountRequest.AccountOwner.IdentityDocumentNumber = account.DocumentNumber
	createAccountRequest.AccountOwner.ContactInformation.Email = account.Email
	createAccountRequest.AccountOwner.ContactInformation.PersonalPhoneNumber1 = "+5551912345678"

	createAccountRequest.BillingAddress.AddressLine1 = "R. Manoelito de Ornellas"
	createAccountRequest.BillingAddress.AddressLine2 = "55"
	createAccountRequest.BillingAddress.City = "Porto Alegre"
	createAccountRequest.BillingAddress.State = "RS"
	createAccountRequest.BillingAddress.Neighborhood = "Praia de Belas"
	createAccountRequest.BillingAddress.Zipcode = "990000"

	createAccountRequest.CardDeliveryAddress.AddressLine1 = "R. Manoelito de Ornellas"
	createAccountRequest.CardDeliveryAddress.AddressLine2 = "55"
	createAccountRequest.CardDeliveryAddress.City = "Porto Alegre"
	createAccountRequest.CardDeliveryAddress.State = "RS"
	createAccountRequest.CardDeliveryAddress.Neighborhood = "Praia de Belas"
	createAccountRequest.CardDeliveryAddress.Zipcode = "990000"

	jsonBody, err := json.Marshal(createAccountRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(jsonBody))

	resp, err := g.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status response:%s", http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res entity.CreateAccountResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	return res.Account.AccountID, nil
}

// CreateVirtualCard Account
func (g *paytGateway) CreateVirtualCard(accountID string) (*entity.VirtualCard, error) {
	uri := fmt.Sprint(g.baseURL, createAccountPath, "/", accountID, createVirtualCardPath)

	createRequest := entity.VirtualCardRequest{
		BirthDate: "15/04/1996",
	}

	createRequest.Constraints.CurrencyCode = 986
	createRequest.Constraints.MaxAmount = "*"
	createRequest.Constraints.UsageLimit = "*"
	createRequest.Constraints.ExpirationTimestamp = "*"

	jsonBody, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(jsonBody))

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status response:%s", http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res entity.VirtualCardResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.VirtualCard, nil
}
