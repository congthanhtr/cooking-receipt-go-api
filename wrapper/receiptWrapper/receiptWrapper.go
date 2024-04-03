package receiptWrapper

import (
	"cooking-receipt/connector/sqliteConnector"
	"cooking-receipt/model"
)

type IReceiptWrapper interface {
	Find() ([]*model.CookingReceipt, error)
	FindById(id string) (*model.CookingReceipt, error)
	Create(recipe model.CookingReceipt) error
	Save(id string, recipe model.CookingReceipt) (*model.CookingReceipt, error)
	Delete(id string) error
	Search(search string) (*model.CookingReceipt, error)
}
type ReceiptWrapper struct {
	db sqliteConnector.IGorm
}

var receiptWrapper = &ReceiptWrapper{}

func GetInstance() *ReceiptWrapper {
	if receiptWrapper.db == nil {
		receiptWrapper = &ReceiptWrapper{
			db: *sqliteConnector.GetInstance(),
		}
	}
	return receiptWrapper
}

func (r *ReceiptWrapper) Find() ([]*model.CookingReceipt, error) {
	var receipts []*model.CookingReceipt
	var query = r.db.Find(&receipts)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return receipts, nil
}

func (r *ReceiptWrapper) FindById(id string) (*model.CookingReceipt, error) {
	var receipt *model.CookingReceipt
	var receiptQuery = r.db.Preload("Ingredients")
	receiptQuery.Where("id = ?", id)
	receiptQuery.Find(&receipt)
	err := receiptQuery.Error
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (r *ReceiptWrapper) Create(recipe model.CookingReceipt) error {
	// TODO: implement create here
	return nil
}

func (r *ReceiptWrapper) Save(id string, recipe model.CookingReceipt) (*model.CookingReceipt, error) {
	// TODO: implement update here
	return nil, nil
}

func (r *ReceiptWrapper) Delete(id string) error {
	// TODO: implement update here
	return nil
}

func (r *ReceiptWrapper) Search(search string) (*model.CookingReceipt, error) {
	var receipt model.CookingReceipt
	var query = r.db.Where("name LIKE ? OR Description LIKE ?", search, search)
	query.Find(&receipt)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return &receipt, nil
}
