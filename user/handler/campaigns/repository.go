package campaigns

import "gorm.io/gorm"

type Repository interface {
	FindAll() []Campaigns
	FindByUserID(UserID int) ([]Campaigns, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaigns, error) {
	var campaignsx []Campaigns

	err := r.db.Find(&campaignsx).Error
	if err != nil {
		return campaignsx, err
	}

	return campaignsx, nil
}

func (r *repository) FindByUserID(UserID int) ([]Campaigns, error) {
	var campaignsx []Campaigns

	err := r.db.Where("user_id = ?", UserID).Find(&campaignsx).Error
	if err != nil {
		return campaignsx, err
	}

	return campaignsx, nil
}
