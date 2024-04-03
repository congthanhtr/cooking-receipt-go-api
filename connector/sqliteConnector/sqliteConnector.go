package sqliteConnector

import (
	"cooking-receipt/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Gorm struct {
	db       *gorm.DB
	dbString string
}

type IGorm interface {
	New() error
	Migration() error
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
}

func (s *Gorm) New() error {
	db, err := gorm.Open(sqlite.Open(s.dbString), &gorm.Config{})
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

var sqLite IGorm = &Gorm{}

func GetInstance() *IGorm {
	if sqLite == nil {
		err := sqLite.New()
		if err != nil {
			db, err := gorm.Open(sqlite.Open("cooking-receipt.db"), &gorm.Config{})
			if err != nil {
				return nil
			}
			var g IGorm = &Gorm{
				db: db,
			}
			return &g
		}
	}
	return &sqLite
}

func (s *Gorm) Migration() error {
	err := s.db.AutoMigrate(model.CookingReceipt{})
	err = s.db.AutoMigrate(model.Ingredient{})
	return err
}

func (s *Gorm) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return s.db.Find(dest, conds...)
}

func (s *Gorm) Where(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return s.db.Where(dest, conds)
}

func (s *Gorm) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return s.db.First(dest, &conds)
}

func (s *Gorm) Preload(query string, args ...interface{}) (tx *gorm.DB) {
	return s.db.Preload(query, args)
}

func (s *Gorm) Model(value interface{}) (tx *gorm.DB) {
	return s.db.Model(value)
}
