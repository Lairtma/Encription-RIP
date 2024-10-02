package repository

import (
	"RIP/internal/app/ds"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

//cart status: 0 - черновик, 1 - сформирован, 2 - завершен, 3 - удалён, 4 - отклонён

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

func (r *Repository) GetAllTexts() ([]ds.TextToEncOrDec, error) {
	var prods []ds.TextToEncOrDec
	err := r.db.Where("status=true").Order("id").Find(&prods).Error
	if err != nil {
		return nil, err
	}
	return prods, nil
}
func (r *Repository) GetTextByID(textId int) (*ds.TextToEncOrDec, error) {
	text := &ds.TextToEncOrDec{}
	err := r.db.Where("id = ?", textId).First(text).Error
	if err != nil {
		return nil, err
	}
	return text, nil
}

func (r *Repository) GetTextByType(encType bool) ([]ds.TextToEncOrDec, error) {
	var text []ds.TextToEncOrDec
	err := r.db.Where("enc = ?", encType).Find(&text).Error
	if err != nil {
		return nil, err
	}
	return text, nil
}

func (r *Repository) GetWorkingOrderByUserId(userId int) ([]ds.EncOrDecOrder, error) {
	var order []ds.EncOrDecOrder
	err := r.db.Where("status=0").Where("creator_id = ?", userId).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) GetLastOrder() ([]ds.EncOrDecOrder, error) {
	var order []ds.EncOrDecOrder
	err := r.db.Order("date_create DESC").Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) CreateOrder(userId int, moderatorId int) ([]ds.EncOrDecOrder, error) {
	newOrder := ds.EncOrDecOrder{
		Status:      0,
		DateCreate:  time.Now(),
		DateUpdate:  time.Now(),
		CreatorID:   userId,
		ModeratorID: moderatorId,
	}
	err := r.db.Create(&newOrder).Error
	if err != nil {
		return nil, err
	}
	order, err := r.GetLastOrder()
	return order, err
}

func (r *Repository) GetOrderByID(id int) (*ds.EncOrDecOrder, error) { // ?
	order := &ds.EncOrDecOrder{}
	err := r.db.Where("id = ?", id).First(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) AddToOrder(orderId int, textId int, position int, encType string) error {
	query := "INSERT INTO order_texts (order_id, text_id, position, enc_type) VALUES (?, ?, ?, ?)"
	err := r.db.Exec(query, orderId, textId, position, encType).Error
	if err != nil {
		return fmt.Errorf("failed to add to order: %w", err)
	}
	return nil
}

func (r *Repository) GetTextsByOrderId(orderID int) ([]ds.OrderText, error) {
	var Texts []ds.OrderText

	err := r.db.
		Where("order_texts.order_id = ?", orderID).Order("position ASC").Find(&Texts).Error

	if err != nil {
		return nil, err
	}

	return Texts, nil
}

func (r *Repository) GetTextIdsByOrderId(orderID int) ([]int, error) {
	var textIds []int

	err := r.db.
		Model(&ds.OrderText{}).
		Where("order_id = ?", orderID).
		Pluck("text_id", &textIds).
		Error

	if err != nil {
		return nil, err
	}

	return textIds, nil
}

func (r *Repository) DeleteOrder(id int) error {
	err := r.db.Exec("UPDATE enc_or_dec_orders SET status = ? WHERE id = ?", 3, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetOrderStatusByID(id int) (int, error) {
	order := &ds.EncOrDecOrder{}
	err := r.db.Where("id = ?", id).First(order).Error
	if err != nil {
		return -1, err
	}
	return order.Status, nil
}

func (r *Repository) GetUserFioById(id int) (string, error) {
	user := &ds.Users{}
	err := r.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return "", err
	}
	return user.FIO, nil
}
