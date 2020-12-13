package transaction

import "github.com/julioc98/robin/internal/app/entity"

type service struct {
	repo Repository
}

//NewService create new service factory
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

// Create a Transaction
func (s service) Create(t *Transaction) (int, error) {
	if t.OperationID == 4 {
		t.Amount = t.Amount * -1
	}
	return s.repo.Create(t)
}

// Get a Account
func (s service) Get(id int) (*Transaction, error) {
	return s.repo.Get(id)
}

// AddFunds a Transaction
func (s service) AddFunds(t *Transaction) (int, error) {
	t.Category = "ADD_FOUNDS"
	id, err := s.Create(t)
	if err != nil {
		return 0, err
	}

	amountFloat := float64(t.Amount / 100)

	essentialT := &Transaction{}
	essentialT.Category = "ESSENTIAL"
	essentialT.Amount = int((amountFloat * 0.55) * 100)
	essentialT.OperationID = 1
	essentialT.AccountID = t.AccountID
	_, err = s.Create(essentialT)
	if err != nil {
		return 0, err
	}

	objectiveT := &Transaction{}
	objectiveT.Category = "OBJECTIVE"
	objectiveT.Amount = int((amountFloat * 0.35) * 100)
	objectiveT.OperationID = 1
	objectiveT.AccountID = t.AccountID
	_, err = s.Create(objectiveT)
	if err != nil {
		return 0, err
	}

	freeT := &Transaction{}
	freeT.Category = "FREE"
	freeT.Amount = int((amountFloat * 0.10) * 100)
	freeT.OperationID = 1
	freeT.AccountID = t.AccountID
	_, err = s.Create(freeT)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (s service) GetFounds(accountID string) ([]entity.Found, error) {
	categories := []string{"ESSENTIAL", "OBJECTIVE", "FREE"}

	var founds []entity.Found

	for _, v := range categories {
		amount, err := s.repo.GetFounds(accountID, v)
		if err != nil {
			return nil, err
		}
		founds = append(founds, entity.Found{Amount: amount, Category: v})
	}

	return founds, nil
}

func (s service) GetFoundByCategory(accountID, category string) (int, error) {
	return s.repo.GetFounds(accountID, category)
}
