package repository

import (
	"bonus/model"
	"database/sql"

	"gorm.io/gorm"
)

type AllDB struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func NewAllRepository(db *sql.DB, gormDB *gorm.DB) *AllDB {
	return &AllDB{db: db, gormDB: gormDB}
}

func (r *AllDB) CreateCountry(country *model.Country) error {
	if err := r.gormDB.Create(country).Error; err != nil {
		return err
	}
	return nil
}

func (r *AllDB) GetAllCountries() ([]model.Country, error) {
	var countries []model.Country
	if err := r.gormDB.Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}
func (r *AllDB) GetCountry(cname string) (model.Country, error) {
	var country model.Country
	if err := r.gormDB.Where("cname = ?", cname).First(&country).Error; err != nil {
		return model.Country{}, err
	}
	return country, nil
}
func (r *AllDB) UpdateCountry(country *model.Country) error {
	if err := r.gormDB.Model(&country).Updates(country).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeleteCountry(cname string) error {
	if err := r.gormDB.Where("cname = ?", cname).Delete(&model.Country{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreateDisease(disease *model.Disease) error {
	if err := r.gormDB.Create(disease).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllDiseases() ([]model.Disease, error) {
	var diseases []model.Disease
	if err := r.gormDB.Find(&diseases).Error; err != nil {
		return nil, err
	}
	return diseases, nil
}
func (r *AllDB) GetDisease(diseaseCode string) (model.Disease, error) {
	var disease model.Disease
	if err := r.gormDB.Where("disease_code = ?", diseaseCode).First(&disease).Error; err != nil {
		return model.Disease{}, err
	}
	return disease, nil
}
func (r *AllDB) UpdateDisease(disease *model.Disease) error {
	if err := r.gormDB.Model(&disease).Updates(disease).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeleteDisease(diseaseCode string) error {
	if err := r.gormDB.Where("disease_code = ?", diseaseCode).Delete(&model.Disease{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreateDiseaseType(diseaseType *model.DiseaseType) error {
	if err := r.gormDB.Create(diseaseType).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllDiseaseTypes() ([]model.DiseaseType, error) {
	var diseaseTypes []model.DiseaseType
	if err := r.gormDB.Find(&diseaseTypes).Error; err != nil {
		return nil, err
	}
	return diseaseTypes, nil
}
func (r *AllDB) GetDiseaseType(id int) (model.DiseaseType, error) {
	var diseaseType model.DiseaseType
	if err := r.gormDB.Where("id = ?", id).First(&diseaseType).Error; err != nil {
		return model.DiseaseType{}, err
	}
	return diseaseType, nil
}
func (r *AllDB) UpdateDiseaseType(diseaseType *model.DiseaseType) error {
	if err := r.gormDB.Model(&diseaseType).Updates(diseaseType).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeleteDiseaseType(id int) error {
	if err := r.gormDB.Where("id = ?", id).Delete(&model.DiseaseType{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreateDiscover(discover *model.Discover) error {
	if err := r.gormDB.Create(discover).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllDiscovers() ([]model.Discover, error) {
	var discovers []model.Discover
	if err := r.gormDB.Find(&discovers).Error; err != nil {
		return nil, err
	}
	return discovers, nil
}
func (r *AllDB) GetDiscover(cname string, diseaseCode string) (model.Discover, error) {
	var discover model.Discover
	if err := r.gormDB.Where("cname = ? AND disease_code = ?", cname, diseaseCode).First(&discover).Error; err != nil {
		return model.Discover{}, err
	}
	return discover, nil
}
func (r *AllDB) UpdateDiscover(discover *model.Discover) error {
	if err := r.gormDB.Model(&discover).Updates(discover).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeleteDiscover(cname string, diseaseCode string) error {
	if err := r.gormDB.Where("cname = ? AND disease_code = ?", cname, diseaseCode).Delete(&model.Discover{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreateDoctor(doctor *model.Doctor) error {
	if err := r.gormDB.Create(doctor).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllDoctors() ([]model.Doctor, error) {
	var doctors []model.Doctor
	if err := r.gormDB.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}
func (r *AllDB) GetDoctor(email string) (model.Doctor, error) {
	var doctor model.Doctor
	if err := r.gormDB.Where("email = ?", email).First(&doctor).Error; err != nil {
		return model.Doctor{}, err
	}
	return doctor, nil
}
func (r *AllDB) UpdateDoctor(doctor *model.Doctor) error {
	if err := r.gormDB.Model(&doctor).Updates(doctor).Error; err != nil {
		return err
	}
	return nil
}

func (r *AllDB) DeleteDoctor(email string) error {
	if err := r.gormDB.Where("email = ?", email).Delete(&model.Doctor{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreatePublicServant(publicServant *model.PublicServant) error {
	if err := r.gormDB.Create(publicServant).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllPublicServants() ([]model.PublicServant, error) {
	var publicServants []model.PublicServant
	if err := r.gormDB.Find(&publicServants).Error; err != nil {
		return nil, err
	}
	return publicServants, nil
}
func (r *AllDB) GetPublicServant(email string) (model.PublicServant, error) {
	var publicServant model.PublicServant
	if err := r.gormDB.Where("email = ?", email).First(&publicServant).Error; err != nil {
		return model.PublicServant{}, err
	}
	return publicServant, nil
}
func (r *AllDB) UpdatePublicServant(publicServant *model.PublicServant) error {
	if err := r.gormDB.Model(&publicServant).Updates(publicServant).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeletePublicServant(email string) error {
	if err := r.gormDB.Where("email = ?", email).Delete(&model.PublicServant{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreateRecord(record *model.Record) error {
	if err := r.gormDB.Create(record).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllRecords() ([]model.Record, error) {
	var records []model.Record
	if err := r.gormDB.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
func (r *AllDB) GetRecord(email string, cname string, diseaseCode string) (model.Record, error) {
	var record model.Record
	if err := r.gormDB.Where("email = ? AND cname = ? AND disease_code = ?", email, cname, diseaseCode).First(&record).Error; err != nil {
		return model.Record{}, err
	}
	return record, nil
}
func (r *AllDB) UpdateRecord(record *model.Record) error {
	if err := r.gormDB.Model(&record).Updates(record).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeleteRecord(email string, cname string, diseaseCode string) error {
	if err := r.gormDB.Where("email = ? AND cname = ? AND disease_code = ?", email, cname, diseaseCode).Delete(&model.Record{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreateSpecialize(specialize *model.Specialize) error {
	if err := r.gormDB.Create(specialize).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllSpecializes() ([]model.Specialize, error) {
	var specializes []model.Specialize
	if err := r.gormDB.Find(&specializes).Error; err != nil {
		return nil, err
	}
	return specializes, nil
}
func (r *AllDB) GetSpecialize(id int, email string) (model.Specialize, error) {
	var specialize model.Specialize
	if err := r.gormDB.Where("id = ? AND email = ?", id, email).First(&specialize).Error; err != nil {
		return model.Specialize{}, err
	}
	return specialize, nil
}
func (r *AllDB) UpdateSpecialize(specialize *model.Specialize) error {
	if err := r.gormDB.Model(&specialize).Updates(specialize).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeleteSpecialize(id int, email string) error {
	if err := r.gormDB.Where("id = ? AND email = ?", id, email).Delete(&model.Specialize{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) CreateUsers(users *model.Users) error {
	if err := r.gormDB.Create(users).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) GetAllUsers() ([]model.Users, error) {
	var users []model.Users
	if err := r.gormDB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *AllDB) GetUsers(email string) (model.Users, error) {
	var users model.Users
	if err := r.gormDB.Where("email = ?", email).First(&users).Error; err != nil {
		return model.Users{}, err
	}
	return users, nil
}
func (r *AllDB) UpdateUsers(users *model.Users) error {
	if err := r.gormDB.Model(&users).Updates(users).Error; err != nil {
		return err
	}
	return nil
}
func (r *AllDB) DeleteUsers(email string) error {
	if err := r.gormDB.Where("email = ?", email).Delete(&model.Users{}).Error; err != nil {
		return err
	}
	return nil
}
