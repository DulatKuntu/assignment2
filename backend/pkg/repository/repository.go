package repository

import (
	"bonus/model"
	"database/sql"

	"gorm.io/gorm"
)

type All interface {
	CreateCountry(country *model.Country) error
	GetAllCountries() ([]model.Country, error)
	GetCountry(cname string) (model.Country, error)
	UpdateCountry(country *model.Country) error
	DeleteCountry(cname string) error
	CreateDisease(disease *model.Disease) error
	GetAllDiseases() ([]model.Disease, error)
	GetDisease(diseaseCode string) (model.Disease, error)
	UpdateDisease(disease *model.Disease) error
	DeleteDisease(diseaseCode string) error
	CreateDiseaseType(diseaseType *model.DiseaseType) error
	GetAllDiseaseTypes() ([]model.DiseaseType, error)
	GetDiseaseType(id int) (model.DiseaseType, error)
	UpdateDiseaseType(diseaseType *model.DiseaseType) error
	DeleteDiseaseType(id int) error
	CreateDiscover(discover *model.Discover) error
	GetAllDiscovers() ([]model.Discover, error)
	GetDiscover(cname string, diseaseCode string) (model.Discover, error)
	UpdateDiscover(discover *model.Discover) error
	DeleteDiscover(cname string, diseaseCode string) error
	CreateDoctor(doctor *model.Doctor) error
	GetAllDoctors() ([]model.Doctor, error)
	GetDoctor(email string) (model.Doctor, error)
	UpdateDoctor(doctor *model.Doctor) error

	DeleteDoctor(email string) error
	CreatePublicServant(publicServant *model.PublicServant) error
	GetAllPublicServants() ([]model.PublicServant, error)
	GetPublicServant(email string) (model.PublicServant, error)
	UpdatePublicServant(publicServant *model.PublicServant) error
	DeletePublicServant(email string) error
	CreateRecord(record *model.Record) error
	GetAllRecords() ([]model.Record, error)
	GetRecord(email string, cname string, diseaseCode string) (model.Record, error)
	UpdateRecord(record *model.Record) error
	DeleteRecord(email string, cname string, diseaseCode string) error
	CreateSpecialize(specialize *model.Specialize) error
	GetAllSpecializes() ([]model.Specialize, error)
	GetSpecialize(id int, email string) (model.Specialize, error)
	UpdateSpecialize(specialize *model.Specialize) error
	DeleteSpecialize(id int, email string) error
	CreateUsers(users *model.Users) error
	GetAllUsers() ([]model.Users, error)
	GetUsers(email string) (model.Users, error)
	UpdateUsers(users *model.Users) error
	DeleteUsers(email string) error
}

type Repository struct {
	All
}

func NewRepository(db *sql.DB, gormDB *gorm.DB) *Repository {
	return &Repository{
		All: NewAllRepository(db, gormDB),
	}
}
