package repository

import (
	"completeUser/features/user/domain"
	"time"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(newUser)
	if err := rq.db.Exec("INSERT INTO users (id, created_at, updated_at, deleted_at, email, name, phone, address, password, status) values (?,?,?,?,?,?,?,?,?,?)",
		nil, time.Now(), time.Now(), nil, newUser.Email, newUser.Name, newUser.Phone, newUser.Address, "qwerty", 0).Error; err != nil {
		return domain.Core{}, err
	}
	newUser = ToDomain(cnv)
	return newUser, nil
}

func (rq *repoQuery) Update(updatedData domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(updatedData)
	if err := rq.db.Exec("UPDATE users SET updated_at = ?, email = ?, name = ?, phone = ?, address = ? WHERE id = ?",
		time.Now(), cnv.Email, cnv.Name, cnv.Phone, cnv.Address, cnv.ID).Error; err != nil {
		return domain.Core{}, err
	}
	updatedData = ToDomain(cnv)
	return updatedData, nil
}

func (rq *repoQuery) Delete(ID uint) error {
	var resQry User
	if err := rq.db.Where("ID = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}

func (rq *repoQuery) Get(ID uint) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var resQry []User
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}
