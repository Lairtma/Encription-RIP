package repository

import (
	"RIP/internal/app/ds"
	"RIP/internal/app/schemas"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
func (r *Repository) GetTextByID(textId string) (ds.TextToEncOrDec, error) {
	text := ds.TextToEncOrDec{}
	err := r.db.Where("id = ?", textId).First(&text).Error
	if err != nil {
		return ds.TextToEncOrDec{}, err
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

func (r *Repository) GetLastOrder() (ds.EncOrDecOrder, error) {
	var order ds.EncOrDecOrder
	err := r.db.Order("date_create DESC").Find(&order).Error
	if err != nil {
		return ds.EncOrDecOrder{}, err
	}
	return order, nil
}

func (r *Repository) CreateOrder(userId int, moderatorId int) (ds.EncOrDecOrder, error) {
	newOrder := ds.EncOrDecOrder{
		Status:      0,
		DateCreate:  time.Now(),
		DateUpdate:  time.Now(),
		CreatorID:   &userId,
		ModeratorID: &moderatorId,
		Priority:    1,
	}
	err := r.db.Create(&newOrder).Error
	if err != nil {
		return ds.EncOrDecOrder{}, err
	}
	order, err := r.GetLastOrder()
	return order, err
}

func (r *Repository) GetOrderByID(id int) (ds.EncOrDecOrder, error) { // ?
	order := ds.EncOrDecOrder{}
	err := r.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return ds.EncOrDecOrder{}, err
	}

	return order, nil
}

func (r *Repository) AddToOrder(orderId int, textId int, position int, encType string) error {
	orderText := ds.OrderText{
		OrderID:  orderId,
		TextID:   textId,
		Position: position,
		EncType:  encType,
	}

	err := r.db.Create(&orderText).Error
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

func (r *Repository) CreateUser(user ds.Users) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateText(text ds.TextToEncOrDec) error {
	if err := r.db.Create(&text).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteTextByID(id string) error {
	var text ds.TextToEncOrDec
	if err := r.db.First(&text, id).Error; err != nil {
		return err
	}
	text.Img = ""
	text.Status = false
	if err := r.db.Save(&text).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTextByID(id string, text ds.TextToEncOrDec) error {
	var existingText ds.TextToEncOrDec
	if err := r.db.First(&existingText, "id = ?", id).Error; err != nil {
		return err
	}

	existingText.Enc = text.Enc
	existingText.Text = text.Text
	existingText.Img = text.Img
	existingText.ByteLen = text.ByteLen
	existingText.Status = text.Status

	err := r.db.Save(&existingText).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ChangePicByID(id string, image string) error {
	// 1. Поиск записи по ID
	text := ds.TextToEncOrDec{}
	result := r.db.First(&text, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("запись с ID %s не найдена", id)
	}
	text.Img = image
	err := r.db.Save(&text).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllOrdersWithFilters(status int, having_status bool) ([]ds.EncOrDecOrder, error) {
	var milkRequests []ds.EncOrDecOrder
	log.Println(status, having_status)
	db := r.db // Инициализируем db без фильтра по дате
	if having_status {
		db = db.Where("status = ?", status) // Фильтр по статусу
	}
	err := db.Find(&milkRequests).Error // Выборка записей из базы данных
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	return milkRequests, nil
}

func (r *Repository) UpdateFieldsOrder(request schemas.UpdateFieldsOrderRequest) error {
	var order ds.EncOrDecOrder
	// Загрузка записи из базы данных по ID
	if err := r.db.First(&order, "id = ?", request.Id).Error; err != nil {
		return err
	}
	if request.Priority != -1 {
		order.Priority = request.Priority
	}
	if err := r.db.Save(&order).Error; err != nil {
		return err
	}
	return nil // Возвращаем nil, если все прошло успешно
}

func (r *Repository) FormOrder(id string) error {
	var order ds.EncOrDecOrder
	if err := r.db.First(&order, "id = ?", id).Error; err != nil {
		return err
	}
	if order.CreatorID == nil {
		err := fmt.Errorf("Unable to finish request. Probably some fields are empty")
		return err
	}
	order.Status = 1
	if err := r.db.Save(&order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) FinishOrder(id string, status int) error {
	var milkRequest ds.EncOrDecOrder
	if err := r.db.First(&milkRequest, "id = ?", id).Error; err != nil {
		return err
	}
	if milkRequest.CreatorID == nil {
		err := fmt.Errorf("Unable to finish request. Probably some fields are empty")
		return err
	}
	mod_id := 2
	milkRequest.Status = status
	milkRequest.DateFinish = time.Now()
	milkRequest.ModeratorID = &mod_id
	if err := r.db.Save(&milkRequest).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteTextFromOrder(id int, text_id int) error {
	var order_text ds.OrderText
	if err := r.db.Where("order_id = ? AND text_id = ?", id, text_id).First(&order_text).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&order_text).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdatePositionTextInOrder(id int, text_id int, position int) error {
	var order_text ds.OrderText
	if err := r.db.Where("order_id = ? AND text_id = ?", id, text_id).First(&order_text).Error; err != nil {
		return err
	}
	order_text.Position = position
	if err := r.db.Save(&order_text).Error; err != nil {
		return err
	}
	return nil
}
