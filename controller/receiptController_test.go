package controller

import (
	"cooking-receipt/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"testing"
)

type mockReceiptWrapper struct {
	find     func() ([]*model.CookingReceipt, error)
	findById func(id string) (*model.CookingReceipt, error)
	create   func(recipe model.CookingReceipt) error
	save     func(id string, recipe model.CookingReceipt) (*model.CookingReceipt, error)
	delete   func(id string) error
	search   func(search string) (*model.CookingReceipt, error)
}

func (m *mockReceiptWrapper) Find() ([]*model.CookingReceipt, error) {
	return m.find()
}

func (m *mockReceiptWrapper) FindById(id string) (*model.CookingReceipt, error) {
	return m.findById(id)
}

func (m *mockReceiptWrapper) Create(recipe model.CookingReceipt) error {
	return m.create(recipe)
}

func (m *mockReceiptWrapper) Save(id string, recipe model.CookingReceipt) (*model.CookingReceipt, error) {
	return m.save(id, recipe)
}

func (m *mockReceiptWrapper) Delete(id string) error {
	return m.delete(id)
}

func (m *mockReceiptWrapper) Search(search string) (*model.CookingReceipt, error) {
	return m.search(search)
}

func TestReceiptController_FindAll(t *testing.T) {
	controller := &ReceiptController{
		Receipt: &mockReceiptWrapper{
			find: func() ([]*model.CookingReceipt, error) {
				return []*model.CookingReceipt{
					{Name: "test name"},
				}, nil
			},
		},
	}
	var ctx = &gin.Context{}
	controller.FindAll(ctx)
	assert.Equal(t, ctx.Status, 200)

	// find error
	controller = &ReceiptController{
		Receipt: &mockReceiptWrapper{
			find: func() ([]*model.CookingReceipt, error) {
				return nil, errors.New("find error")
			},
		},
	}
	ctx = &gin.Context{}
	controller.FindAll(ctx)
	assert.Equal(t, ctx.Status, 500)
}

func TestFindAll(t *testing.T) {
	controller := &ReceiptController{
		Receipt: &mockReceiptWrapper{
			findById: func(id string) (*model.CookingReceipt, error) {
				return &model.CookingReceipt{}, nil
			},
		},
	}
	ctx := &gin.Context{
		Params: gin.Params{
			{Key: "id", Value: "test id"},
		},
	}
	controller.FindRecipe(ctx)
	assert.Equal(t, ctx.Status, 200)

	// find error
	controller = &ReceiptController{
		Receipt: &mockReceiptWrapper{
			findById: func(id string) (*model.CookingReceipt, error) {
				return nil, errors.New("error when find by id")
			},
		},
	}
	ctx = &gin.Context{
		Params: gin.Params{
			{Key: "id", Value: "test id"},
		},
	}
	controller.FindRecipe(ctx)
	assert.Equal(t, ctx.Status, 500)
}

func TestSearch(t *testing.T) {
	controller := &ReceiptController{
		Receipt: &mockReceiptWrapper{
			search: func(search string) (*model.CookingReceipt, error) {
				return &model.CookingReceipt{}, nil
			},
		},
	}
	ctx := &gin.Context{
		Params: gin.Params{
			{Key: "search", Value: "test id"},
		},
	}
	controller.Seach(ctx)
	assert.Equal(t, ctx.Status, 200)

	// find error
	controller = &ReceiptController{
		Receipt: &mockReceiptWrapper{
			search: func(id string) (*model.CookingReceipt, error) {
				return nil, errors.New("error when search")
			},
		},
	}
	ctx = &gin.Context{
		Params: gin.Params{
			{Key: "search", Value: "test id"},
		},
	}
	controller.Seach(ctx)
	assert.Equal(t, ctx.Status, 500)
}

// TODO: implement other tests
