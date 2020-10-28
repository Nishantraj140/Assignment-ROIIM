package user

import (
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email              string
	FirstName          string
	MiddleName         string
	LastName           string
	Phone              string
	BillingDetailsId   string
	MerchantCustomerId string
	CustomerId         string
	Password           string
}

func (u *User) Create() (err error) {
	err = sql.DB.Model(u).Create(u).Error
	return
}

func (u *User) Get() (err error) {
	return sql.DB.Model(u).Where("email = ?", u.Email).First(u).Error
}

func (u User) Update(pid string, aid string) (err error) {
	err = sql.DB.Model(u).Updates(User{CustomerId: pid, BillingDetailsId: aid}).Error
	return
}

type UserCard struct {
	gorm.Model
	CustomerId         string
	PaymentHandleId    string
	PaymentHandleToken string
}

func (uc *UserCard) GetAll() ([]UserCard, error) {
	var ucs []UserCard
	err := sql.DB.Model(uc).Where("customer_id=?", uc.CustomerId).Find(&ucs).Error
	return ucs, err
}

func (uc *UserCard) Create() (err error) {
	err = sql.DB.Model(uc).Create(uc).Error
	return
}