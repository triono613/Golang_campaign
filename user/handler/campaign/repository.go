package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() []Campaign
	FindByUserID(UserID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) FindByUserID(UserID int) ([]Campaign, error) {
	var campaignsx []Campaign

	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = 1 ").Find(&campaignsx).Error
	if err != nil {
		return campaignsx, err
	}

	return campaignsx, nil
}
