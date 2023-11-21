package handler

import (
	_ "bonus/docs"
	"bonus/pkg/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.GET("/getAll", h.getAllUsers)
		users.GET("/get/:id", h.getUser)
		users.POST("/create", h.createUser)
		users.POST("/update", h.updateUser)
		users.GET("/delete/:id", h.deleteUser)
	}
	caregivers := router.Group("/caregivers")
	{
		caregivers.GET("/getAll", h.getAllCaregivers)
		caregivers.GET("/get/:id", h.getCaregivers)
		caregivers.POST("/create", h.createCaregivers)
		caregivers.POST("/update", h.updateCaregivers)
		caregivers.GET("/delete/:id", h.deleteCaregivers)
	}
	members := router.Group("/members")
	{
		members.GET("/getAll", h.getAllMembers)
		members.GET("/get/:id", h.getMembers)
		members.POST("/create", h.createMembers)
		members.POST("/update", h.updateMembers)
		members.GET("/delete/:id", h.deleteMembers)
	}
	addresses := router.Group("/addresses")
	{
		addresses.GET("/getAll", h.getAllAddresses)
		addresses.GET("/get/:id", h.getAddresses)
		addresses.POST("/create", h.createAddresses)
		addresses.POST("/update", h.updateAddresses)
		addresses.GET("/delete/:id", h.deleteAddresses)
	}
	jobs := router.Group("/jobs")
	{
		jobs.GET("/getAll", h.getAllJobs)
		jobs.GET("/get/:id", h.getJobs)
		jobs.POST("/create", h.createJobs)
		jobs.POST("/update", h.updateJobs)
		jobs.GET("/delete/:id", h.deleteJobs)
	}
	jobApplications := router.Group("/job_applications")
	{
		jobApplications.GET("/getAll", h.getAllJobApplications)
		jobApplications.GET("/get/:id", h.getJobApplications)
		jobApplications.POST("/create", h.createJobApplications)
		jobApplications.POST("/update", h.updateJobApplications)
		jobApplications.GET("/delete/:id", h.deleteJobApplications)
	}
	appointments := router.Group("/appointments")
	{
		appointments.GET("/getAll", h.getAllAppointments)
		appointments.GET("/get/:id", h.getAppointments)
		appointments.POST("/create", h.createAppointments)
		appointments.POST("/update", h.updateAppointments)
		appointments.GET("/delete/:id", h.deleteAppointments)
	}

	return router
}
