package repository

import (
	"errors"
	"log"

	"github.com/dunky-star/ecommerce-app-golang/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
	CreateBankAccount(e domain.BankAccount) error

	// Other functions later on
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}
	return user, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {

	var user domain.User

	err := r.db.First(&user, id).Error

	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {

	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error

	if err != nil {
		log.Printf("error on update %v", err)
		return domain.User{}, errors.New("failed update user")
	}

	return user, nil
}
