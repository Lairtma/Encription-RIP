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

func (r *Repository) GetCardByID(id int) (*ds.Card, error) { // ?
	card := &ds.Card{}
	err := r.db.Where("id = ?", id).First(card).Error
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (r *Repository) GetCardByType(encryption bool) ([]ds.Card, error) {
	var card []ds.Card

	err := r.db.
		Where("cards.encrypting = ?", encryption).Where("used = ?", true).Order("id asc").
		Find(&card).Error

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (r *Repository) GetLastCartByCreatorId(creatorId int) (*ds.Cart, error) {
	cart := &ds.Cart{}

	err := r.db.
		Where("carts.creator_id = ?", creatorId).Where("status = ?", 1).Order("id desc").
		First(cart).Error

	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *Repository) GetCartCardsByCartId(cartId int) ([]ds.CartCard, error) {
	var cartCard []ds.CartCard

	err := r.db.
		Where("cart_cards.cart_id = ?", cartId).Order("position asc").
		Find(&cartCard).Error

	if err != nil {
		return nil, err
	}

	return cartCard, nil
}

func (r *Repository) GetAllCards() ([]ds.Card, error) { //FIO ?
	var cards []ds.Card

	err := r.db.Where("used = ?", true).Order("id asc").Find(&cards).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *Repository) CreateProduct(product ds.Card) error {
	return r.db.Create(product).Error
}
