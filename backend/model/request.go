package model

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ReqID struct
type ReqID struct {
	ID int `json:"id"`
}

func (p *ReqID) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.ID == 0 {
		return errors.New("bad request | id is required")
	}

	return nil
}

type ReqCountry struct {
	Cname string `json:"cname"`
}

func (p *ReqCountry) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqDiseaseCode struct {
	DiseaseCode string `json:"diseaseCode"`
}

func (p *ReqDiseaseCode) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqUser struct {
	Email string `json:"email"`
}

func (p *ReqUser) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqDiseaseType struct {
	ID int `json:"id"`
}

func (p *ReqDiseaseType) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqDiscover struct {
	Cname       string `json:"cname"`
	DiseaseCode string `json:"disease_code"`
}

func (p *ReqDiscover) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqSpecialize struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (p *ReqSpecialize) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqRecord struct {
	Email       string `json:"email" gorm:"primaryKey"`
	Cname       string `json:"cname" gorm:"primaryKey"`
	DiseaseCode string `json:"disease_code" gorm:"primaryKey"`
}

func (p *ReqRecord) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}
