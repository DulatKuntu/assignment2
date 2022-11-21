package handler

import (
	"bonus/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllCountries(c *gin.Context) {
	res, err := h.repo.GetAllCountries()

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) getCountry(c *gin.Context) {
	cname := c.Param("cname")
	res, err := h.repo.GetCountry(cname)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) createCountry(c *gin.Context) {
	var country *model.Country
	if err := c.BindJSON(&country); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.CreateCountry(country)

	if err != nil {
		defaultErrorHandler(c, err)
	}

	sendSuccess(c)
}

func (h *Handler) updateCountry(c *gin.Context) {
	var country *model.Country
	if err := c.BindJSON(&country); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateCountry(country)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) deleteCountry(c *gin.Context) {
	cname := c.Param("cname")
	err := h.repo.DeleteCountry(cname)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) getAllDiseases(c *gin.Context) {
	res, err := h.repo.GetAllDiseases()

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) getDisease(c *gin.Context) {
	diseaseCode := c.Param("disease_code")
	res, err := h.repo.GetDisease(diseaseCode)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) createDisease(c *gin.Context) {
	var disease *model.Disease
	if err := c.BindJSON(&disease); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreateDisease(disease)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) updateDisease(c *gin.Context) {
	var disease *model.Disease
	if err := c.BindJSON(&disease); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateDisease(disease)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) deleteDisease(c *gin.Context) {
	var req model.ReqDiseaseCode
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.DeleteDisease(req.DiseaseCode)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) createUser(c *gin.Context) {
	var user *model.Users
	if err := c.BindJSON(&user); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreateUsers(user)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	res, err := h.repo.GetAllUsers()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) getUser(c *gin.Context) {
	email := c.Param("email")
	res, err := h.repo.GetUsers(email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) updateUser(c *gin.Context) {
	var user *model.Users
	if err := c.BindJSON(&user); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateUsers(user)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) deleteUser(c *gin.Context) {
	email := c.Param("email")
	err := h.repo.DeleteUsers(email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) createDiseaseType(c *gin.Context) {
	var diseaseType *model.DiseaseType
	if err := c.BindJSON(&diseaseType); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreateDiseaseType(diseaseType)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) getAllDiseaseTypes(c *gin.Context) {
	res, err := h.repo.GetAllDiseaseTypes()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) getDiseaseType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	res, err := h.repo.GetDiseaseType(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) updateDiseaseType(c *gin.Context) {
	var diseaseType *model.DiseaseType
	if err := c.BindJSON(&diseaseType); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateDiseaseType(diseaseType)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) deleteDiseaseType(c *gin.Context) {
	var req model.ReqDiseaseType
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.DeleteDiseaseType(req.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) createDiscover(c *gin.Context) {
	var discover *model.Discover
	if err := c.BindJSON(&discover); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreateDiscover(discover)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) getAllDiscovers(c *gin.Context) {
	res, err := h.repo.GetAllDiscovers()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) getDiscover(c *gin.Context) {
	cname := c.Param("cname")
	diseaseCode := c.Param("diseaseCode")
	res, err := h.repo.GetDiscover(cname, diseaseCode)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) updateDiscover(c *gin.Context) {
	var discover *model.Discover
	if err := c.BindJSON(&discover); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateDiscover(discover)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) deleteDiscover(c *gin.Context) {
	var req model.ReqDiscover
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.DeleteDiscover(req.DiseaseCode, req.Cname)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) createDoctor(c *gin.Context) {
	var doctor *model.Doctor
	if err := c.BindJSON(&doctor); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreateDoctor(doctor)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) getAllDoctors(c *gin.Context) {
	res, err := h.repo.GetAllDoctors()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) getDoctor(c *gin.Context) {
	email := c.Param("email")
	res, err := h.repo.GetDoctor(email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) updateDoctor(c *gin.Context) {
	var doctor *model.Doctor
	if err := c.BindJSON(&doctor); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateDoctor(doctor)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) deleteDoctor(c *gin.Context) {
	var req model.ReqUser
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.DeleteDoctor(req.Email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) createPublicServant(c *gin.Context) {
	var publicServant *model.PublicServant
	if err := c.BindJSON(&publicServant); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreatePublicServant(publicServant)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) getAllPublicServants(c *gin.Context) {
	res, err := h.repo.GetAllPublicServants()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) getPublicServant(c *gin.Context) {
	email := c.Param("email")
	res, err := h.repo.GetPublicServant(email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) updatePublicServant(c *gin.Context) {
	var publicServant *model.PublicServant
	if err := c.BindJSON(&publicServant); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdatePublicServant(publicServant)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) deletePublicServant(c *gin.Context) {
	var req model.ReqUser
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.DeletePublicServant(req.Email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) createSpecialize(c *gin.Context) {
	var specialize *model.Specialize
	if err := c.BindJSON(&specialize); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreateSpecialize(specialize)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) getAllSpecializes(c *gin.Context) {
	res, err := h.repo.GetAllSpecializes()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) getSpecialize(c *gin.Context) {
	email := c.Param("email")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	res, err := h.repo.GetSpecialize(id, email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) updateSpecialize(c *gin.Context) {
	var specialize *model.Specialize
	if err := c.BindJSON(&specialize); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateSpecialize(specialize)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) deleteSpecialize(c *gin.Context) {
	var req model.ReqSpecialize
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.DeleteSpecialize(req.ID, req.Email)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) createRecord(c *gin.Context) {
	var record *model.Record
	if err := c.BindJSON(&record); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.CreateRecord(record)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) getAllRecords(c *gin.Context) {
	res, err := h.repo.GetAllRecords()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) getRecord(c *gin.Context) {
	email := c.Param("email")
	cname := c.Param("cname")
	diseaseCode := c.Param("diseaseCode")
	res, err := h.repo.GetRecord(email, cname, diseaseCode)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}

func (h *Handler) updateRecord(c *gin.Context) {
	var record *model.Record
	if err := c.BindJSON(&record); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.UpdateRecord(record)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) deleteRecord(c *gin.Context) {
	var req model.ReqRecord
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	err := h.repo.DeleteRecord(req.Email, req.Cname, req.DiseaseCode)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}
