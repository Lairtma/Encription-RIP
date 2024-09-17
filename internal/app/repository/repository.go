package repository

import (
	"RIP/internal/app/ds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetCardByType(encryption bool) ([]ds.Card, error) {
	var card []ds.Card

	err := r.db.
		Where("cards.encrypting = ?", encryption).Where("used = ?", true).Order("by cards.id").
		Find(&card).Error

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (r *Repository) GetAllCards() ([]ds.Card, error) { //FIO ?
	var cards []ds.Card

	err := r.db.Find(&cards).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *Repository) CreateProduct(product ds.Card) error {
	return r.db.Create(product).Error
}
