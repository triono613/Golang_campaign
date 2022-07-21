package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByName(name string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) Save(user User) (User, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	fmt.Println("repo data err : ", err)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByName(name string) (User, error) {
	var user User
	err := r.db.Where("name = ?", name).Find(&user).Error
	r.db.Debug().Where("name = ?", name).Find(&User{})

	//fmt.Println("repo data name : ", err)
	//log.Println("repo data name : ", user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}
