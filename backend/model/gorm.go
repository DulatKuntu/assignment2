package model

import "time"

type DiseaseType struct {
	ID          int          `json:"id" gorm:"primaryKey"`
	Description string       `json:"description"`
	Diseases    []Disease    `json:"diseases" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;"`
	Specialize  []Specialize `json:"specialize" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;"`
}

func (d *DiseaseType) TableName() string {
	return "diseasetype"
}

type Disease struct {
	DiseaseCode string   `json:"disease_code" gorm:"primaryKey"`
	Pathogen    string   `json:"pathogen"`
	Description string   `json:"description"`
	ID          int      `json:"id"`
	Record      []Record `json:"record" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DiseaseCode;"`
}

func (d *Disease) TableName() string {
	return "disease"
}

type Country struct {
	Cname      string     `json:"cname" gorm:"primaryKey"`
	Population int64      `json:"population"`
	Discover   []Discover `json:"discover" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Cname;"`
	Users      []Users    `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Cname;"`
	Record     []Record   `json:"record" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Cname;"`
}

func (d *Country) TableName() string {
	return "country"
}

type Discover struct {
	Cname        string    `json:"cname" gorm:"primaryKey"`
	DiseaseCode  string    `json:"disease_code" gorm:"primaryKey"`
	FirstEncDate time.Time `json:"first_enc_date"`
}

func (d *Discover) TableName() string {
	return "discover"
}

type Users struct {
	Email         string          `json:"email" gorm:"primaryKey"`
	Name          string          `json:"name"`
	Surname       string          `json:"surname"`
	Salary        int             `json:"salary"`
	Phone         string          `json:"phone"`
	Cname         string          `json:"cname"`
	Doctor        []Doctor        `json:"doctor" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Email;"`
	PublicServant []PublicServant `json:"public_servant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Email;"`
}

func (d *Users) TableName() string {
	return "users"
}

type PublicServant struct {
	Email      string   `json:"email" gorm:"primaryKey"`
	Department string   `json:"department"`
	Record     []Record `json:"record" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Email;"`
}

func (d *PublicServant) TableName() string {
	return "publicservant"
}

type Doctor struct {
	Email      string       `json:"email" gorm:"primaryKey"`
	Degree     string       `json:"degree"`
	Specialize []Specialize `json:"specialize" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Email;"`
}

func (d *Doctor) TableName() string {
	return "doctor"
}

type Specialize struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"primaryKey"`
}

func (d *Specialize) TableName() string {
	return "specialize"
}

type Record struct {
	Email         string `json:"email" gorm:"primaryKey"`
	Cname         string `json:"cname" gorm:"primaryKey"`
	DiseaseCode   string `json:"disease_code" gorm:"primaryKey"`
	TotalDeaths   int    `json:"total_deaths"`
	TotalPatients int    `json:"total_patients"`
}

func (d *Record) TableName() string {
	return "record"
}
