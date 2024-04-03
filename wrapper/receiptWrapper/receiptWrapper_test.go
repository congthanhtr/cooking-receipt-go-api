package receiptWrapper

import (
	"cooking-receipt/model"
	"errors"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
)

type mockGorm struct {
	new       func() error
	migration func() error
	find      func(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	where     func(query interface{}, args ...interface{}) (tx *gorm.DB)
	first     func(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	preload   func(query string, args ...interface{}) (tx *gorm.DB)
	model     func(value interface{}) (tx *gorm.DB)
}

func (m *mockGorm) New() error {
	return m.new()
}

func (m *mockGorm) Migration() error {
	return m.migration()
}

func (m *mockGorm) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return m.find(dest, conds...)
}

func (m *mockGorm) Where(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return m.where(dest, conds)
}

func (m *mockGorm) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return m.first(dest, &conds)
}

func (m *mockGorm) Preload(query string, args ...interface{}) (tx *gorm.DB) {
	return m.preload(query, args)
}

func (m *mockGorm) Model(value interface{}) (tx *gorm.DB) {
	return m.model(value)
}

func TestGetInstance(t *testing.T) {
	wrapper := GetInstance()
	assert.NotEqual(t, wrapper.db, nil)
}

func TestReceiptWrapper_Find(t *testing.T) {
	// happy case
	wrapper := ReceiptWrapper{
		db: &mockGorm{
			find: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: nil,
				}
			},
		},
	}
	_, err := wrapper.Find()
	assert.NotEqual(t, err, nil)

	// error when find
	wrapper = ReceiptWrapper{
		db: &mockGorm{
			find: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: errors.New("error when find"),
				}
			},
		},
	}

	_, err = wrapper.Find()
	assert.Equal(t, err, errors.New("error when find"))
}

func TestReceiptWrapper_FindById(t *testing.T) {
	// happy case
	wrapper := ReceiptWrapper{
		db: &mockGorm{
			find: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				dest = &model.CookingReceipt{
					Name: "test name",
				}
				return &gorm.DB{
					Error: nil,
				}
			},
			where: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: nil,
				}
			},
			preload: func(query string, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: nil,
				}
			},
		},
	}
	receipt, err := wrapper.FindById("temp id")
	assert.NotEqual(t, err, nil)
	assert.NotEqual(t, receipt.Name, "test name")

	// error case
	wrapper = ReceiptWrapper{
		db: &mockGorm{
			find: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: errors.New("error when find"),
				}
			},
			where: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: nil,
				}
			},
			preload: func(query string, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: nil,
				}
			},
		},
	}
	receipt, err = wrapper.FindById("temp id")
	assert.Equal(t, receipt, nil)
	assert.NotEqual(t, err, nil)
}

func TestReceiptWrapper_Search(t *testing.T) {
	// happy case
	wrapper := ReceiptWrapper{
		db: &mockGorm{
			find: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				dest = &model.CookingReceipt{
					Name: "test name",
				}
				return &gorm.DB{
					Error: nil,
				}
			},
			where: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: nil,
				}
			},
		},
	}
	receipt, err := wrapper.Search("temp id")
	assert.NotEqual(t, err, nil)
	assert.NotEqual(t, receipt.Name, "test name")

	// error case
	wrapper = ReceiptWrapper{
		db: &mockGorm{
			find: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: errors.New("error when find"),
				}
			},
			where: func(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
				return &gorm.DB{
					Error: nil,
				}
			},
		},
	}
	receipt, err = wrapper.Search("temp id")
	assert.Equal(t, receipt, nil)
	assert.NotEqual(t, err, nil)
}
