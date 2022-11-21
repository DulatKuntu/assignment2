package handler

import (
	"bonus/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	country := router.Group("/country")
	{
		country.GET("/getAll", h.getAllCountries)
		country.GET("/get/:cname", h.getCountry)
		country.POST("/create", h.createCountry)
		country.POST("/update", h.updateCountry)
		country.GET("/delete/:cname", h.deleteCountry)
	}
	users := router.Group("/users")
	{
		users.GET("/getAll", h.getAllUsers)
		users.GET("/get/:email", h.getUser)
		users.POST("/create", h.createUser)
		users.POST("/update", h.updateUser)
		users.GET("/delete/:email", h.deleteUser)
	}
	disease := router.Group("/disease")
	{
		disease.GET("/getAll", h.getAllDiseases)
		disease.GET("/get/:disease_code", h.getDisease)
		disease.POST("/create", h.createDisease)
		disease.POST("/update", h.updateDisease)
		disease.POST("/delete", h.deleteDisease)
	}
	diseaseType := router.Group("/diseaseType")
	{
		diseaseType.GET("/getAll", h.getAllDiseaseTypes)
		diseaseType.GET("/get/:id", h.getDiseaseType)
		diseaseType.POST("/create", h.createDiseaseType)
		diseaseType.POST("/update", h.updateDiseaseType)
		diseaseType.POST("/delete", h.deleteDiseaseType)
	}
	discover := router.Group("/discover")
	{
		discover.GET("/getAll", h.getAllDiscovers)
		discover.GET("/get/:cname/:disease_code", h.getDiscover)
		discover.POST("/create", h.createDiscover)
		discover.POST("/update", h.updateDiscover)
		discover.POST("/delete", h.deleteDiscover)
	}
	doctor := router.Group("/doctor")
	{
		doctor.GET("/getAll", h.getAllDoctors)
		doctor.GET("/get/:email", h.getDoctor)
		doctor.POST("/create", h.createDoctor)
		doctor.POST("/update", h.updateDoctor)
		doctor.POST("/delete", h.deleteDoctor)
	}
	publicServant := router.Group("/publicServant")
	{
		publicServant.GET("/getAll", h.getAllPublicServants)
		publicServant.GET("/get/:email", h.getPublicServant)
		publicServant.POST("/create", h.createPublicServant)
		publicServant.POST("/update", h.updatePublicServant)
		publicServant.POST("/delete", h.deletePublicServant)
	}
	specialize := router.Group("/specialize")
	{
		specialize.GET("/getAll", h.getAllSpecializes)
		specialize.GET("/get/:id/:email", h.getSpecialize)
		specialize.POST("/create", h.createSpecialize)
		specialize.POST("/update", h.updateSpecialize)
		specialize.POST("/delete", h.deleteSpecialize)
	}
	record := router.Group("/record")
	{
		record.GET("/getAll", h.getAllRecords)
		record.GET("/get/:email/:cname/:diseaseCode", h.getRecord)
		record.POST("/create", h.createRecord)
		record.POST("/update", h.updateRecord)
		record.POST("/delete", h.deleteRecord)
	}

	return router
}
