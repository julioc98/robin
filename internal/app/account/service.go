package account

import (
	"github.com/julioc98/robin/internal/app/entity"
	"github.com/julioc98/robin/internal/app/processor"
)

type service struct {
	gateway processor.Gateway
	repo    Repository
}

//NewService create new service factory
func NewService(r Repository, gateway processor.Gateway) Service {
	return &service{
		repo:    r,
		gateway: gateway,
	}
}

// Create a Account
func (s service) Create(a *entity.Account) (string, error) {
	externalID, err := s.gateway.CreateAccount(a)
	if err != nil {
		return "", err
	}
	a.ExternalID = externalID
	a.ID = externalID
	virtualCard, err := s.gateway.CreateVirtualCard(externalID)
	if err != nil {
		return "", err
	}
	a.VirtualCard = *virtualCard

	return s.repo.Create(a)
}

// Get a Account
func (s service) Get(id string) (*entity.Account, error) {
	return s.repo.Get(id)
}

// Get a Account
func (s service) GetByEmailAndPassword(email, password string) (*entity.Account, error) {
	return s.repo.GetByEmailAndPassword(email, password)
}
