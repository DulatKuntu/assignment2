package handler

import (
	"bonus/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Description Get a list of all users
// @Produce json
// @Tags users
// @Success 200 {array} model.User
// @Router /users/getAll [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	res, err := h.repo.All.GetAllUsers()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Get a user by ID
// @Description Get a user based on the provided ID
// @Produce json
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Router /users/get/{id} [get]
func (h *Handler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	res, err := h.repo.All.GetUser(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Create a new user
// @Description Create a new user with the provided data
// @Produce json
// @Tags users
// @Consumes json
// @Param user body model.User true "User object"
// @Router /users/create [post]
func (h *Handler) createUser(c *gin.Context) {
	req := &model.User{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	id, err := h.repo.CreateUser(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"id": id})
}

// @Summary Update a user
// @Description Update a user with the provided data
// @Produce json
// @Tags users
// @Consumes json
// @Param user body model.User true "User object"
// @Router /users/update [post]
func (h *Handler) updateUser(c *gin.Context) {
	req := &model.User{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.UpdateUser(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Delete a user
// @Description Delete a user based on the provided ID
// @Produce json
// @Tags users
// @Param id path int true "User ID"
// @Router /users/delete/{id} [get]
func (h *Handler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.repo.DeleteUser(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Get all caregivers
// @Description Get a list of all caregivers
// @Tags caregivers
// @Produce json
// @Success 200 {array} model.Caregiver
// @Router /caregivers/getAll [get]
func (h *Handler) getAllCaregivers(c *gin.Context) {
	res, err := h.repo.All.GetAllCaregivers()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Get a caregiver by ID
// @Description Get a caregiver based on the provided ID
// @Tags caregivers
// @Produce json
// @Param id path int true "Caregiver ID"
// @Success 200 {object} model.Caregiver
// @Router /caregivers/get/{id} [get]
func (h *Handler) getCaregivers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	res, err := h.repo.All.GetCaregiver(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Create a new caregiver
// @Description Create a new caregiver with the provided data
// @Tags caregivers
// @Produce json
// @Consumes json
// @Param caregiver body model.Caregiver true "Caregiver object"
// @Router /caregivers/create [post]
func (h *Handler) createCaregivers(c *gin.Context) {
	req := &model.Caregiver{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	id, err := h.repo.CreateCaregiver(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"id": id})
}

// @Summary Update a caregiver
// @Description Update a caregiver with the provided data
// @Tags caregivers
// @Produce json
// @Consumes json
// @Param caregiver body model.Caregiver true "Caregiver object"
// @Router /caregivers/update [post]
func (h *Handler) updateCaregivers(c *gin.Context) {
	req := &model.Caregiver{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.UpdateCaregiver(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Delete a caregiver
// @Description Delete a caregiver based on the provided ID
// @Tags caregivers
// @Produce json
// @Param id path int true "Caregiver ID"
// @Router /caregivers/delete/{id} [get]
func (h *Handler) deleteCaregivers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.repo.DeleteCaregiver(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Get all members
// @Description Get a list of all members
// @Tags members
// @Produce json
// @Success 200 {array} model.Member
// @Router /members/getAll [get]
func (h *Handler) getAllMembers(c *gin.Context) {
	res, err := h.repo.All.GetAllMembers()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Get a member by ID
// @Description Get a member based on the provided ID
// @Tags members
// @Produce json
// @Param id path int true "Member ID"
// @Success 200 {object} model.Member
// @Router /members/get/{id} [get]
func (h *Handler) getMembers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	res, err := h.repo.All.GetMember(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Create a new member
// @Description Create a new member with the provided data
// @Tags members
// @Produce json
// @Consumes json
// @Param member body model.Member true "Member object"
// @Router /members/create [post]
func (h *Handler) createMembers(c *gin.Context) {
	req := &model.Member{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	id, err := h.repo.CreateMember(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"id": id})
}

// @Summary Update a member
// @Description Update a member with the provided data
// @Tags members
// @Produce json
// @Consumes json
// @Param member body model.Member true "Member object"
// @Router /members/update [post]
func (h *Handler) updateMembers(c *gin.Context) {
	req := &model.Member{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.UpdateMember(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Delete a member
// @Description Delete a member based on the provided ID
// @Tags members
// @Produce json
// @Param id path int true "Member ID"
// @Router /members/delete/{id} [get]
func (h *Handler) deleteMembers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.repo.DeleteMember(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Get all addresses
// @Description Get a list of all addresses
// @Tags addresses
// @Produce json
// @Success 200 {array} model.Address
// @Router /addresses/getAll [get]
func (h *Handler) getAllAddresses(c *gin.Context) {
	res, err := h.repo.All.GetAllAddresses()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Get an address by ID
// @Description Get an address based on the provided ID
// @Tags addresses
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} model.Address
// @Router /addresses/get/{id} [get]
func (h *Handler) getAddresses(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	res, err := h.repo.All.GetAddress(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Create a new address
// @Description Create a new address with the provided data
// @Tags addresses
// @Produce json
// @Consumes json
// @Param address body model.Address true "Address object"
// @Router /addresses/create [post]
func (h *Handler) createAddresses(c *gin.Context) {
	req := &model.Address{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	id, err := h.repo.CreateAddress(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"id": id})
}

// @Summary Update an address
// @Description Update an address with the provided data
// @Tags addresses
// @Produce json
// @Consumes json
// @Param address body model.Address true "Address object"
// @Router /addresses/update [post]
func (h *Handler) updateAddresses(c *gin.Context) {
	req := &model.Address{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.UpdateAddress(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Delete an address
// @Description Delete an address based on the provided ID
// @Tags addresses
// @Produce json
// @Param id path int true "Address ID"
// @Router /addresses/delete/{id} [get]
func (h *Handler) deleteAddresses(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.repo.DeleteAddress(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Get all jobs
// @Description Get a list of all jobs
// @Tags jobs
// @Produce json
// @Success 200 {array} model.Job
// @Router /jobs/getAll [get]
func (h *Handler) getAllJobs(c *gin.Context) {
	res, err := h.repo.All.GetAllJobs()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Get a job by ID
// @Description Get a job based on the provided ID
// @Tags jobs
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} model.Job
// @Router /jobs/get/{id} [get]
func (h *Handler) getJobs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	res, err := h.repo.All.GetJob(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Create a new job
// @Description Create a new job with the provided data
// @Tags jobs
// @Produce json
// @Consumes json
// @Param job body model.Job true "Job object"
// @Router /jobs/create [post]
func (h *Handler) createJobs(c *gin.Context) {
	req := &model.Job{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	id, err := h.repo.CreateJob(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"id": id})
}

// @Summary Update a job
// @Description Update a job with the provided data
// @Tags jobs
// @Produce json
// @Consumes json
// @Param job body model.Job true "Job object"
// @Router /jobs/update [post]
func (h *Handler) updateJobs(c *gin.Context) {
	req := &model.Job{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.UpdateJob(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Delete a job
// @Description Delete a job based on the provided ID
// @Tags jobs
// @Produce json
// @Param id path int true "Job ID"
// @Router /jobs/delete/{id} [get]
func (h *Handler) deleteJobs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.repo.DeleteJob(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Get all job applications
// @Description Get a list of all job applications
// @Tags jobApplications
// @Produce json
// @Success 200 {array} model.JobApplication
// @Router /jobApplications/getAll [get]
func (h *Handler) getAllJobApplications(c *gin.Context) {
	res, err := h.repo.All.GetAllJobApplications()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Get a job application by ID
// @Description Get a job application based on the provided ID
// @Tags jobApplications
// @Produce json
// @Param id path int true "Job Application ID"
// @Success 200 {object} model.JobApplication
// @Router /jobApplications/get/{id} [get]
func (h *Handler) getJobApplications(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	res, err := h.repo.All.GetJobApplication(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Create a new job application
// @Description Create a new job application with the provided data
// @Tags jobApplications
// @Produce json
// @Consumes json
// @Param jobApplication body model.JobApplication true "Job Application object"
// @Router /jobApplications/create [post]
func (h *Handler) createJobApplications(c *gin.Context) {
	req := &model.JobApplication{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	id, err := h.repo.CreateJobApplication(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"id": id})
}

// @Summary Update a job application
// @Description Update a job application with the provided data
// @Tags jobApplications
// @Produce json
// @Consumes json
// @Param jobApplication body model.JobApplication true "Job Application object"
// @Router /jobApplications/update [post]
func (h *Handler) updateJobApplications(c *gin.Context) {
	req := &model.JobApplication{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.UpdateJobApplication(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Delete a job application
// @Description Delete a job application based on the provided ID
// @Tags jobApplications
// @Produce json
// @Param id path int true "Job Application ID"
// @Router /jobApplications/delete/{id} [get]
func (h *Handler) deleteJobApplications(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.repo.DeleteJobApplication(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Get all appointments
// @Description Get a list of all appointments
// @Tags appointments
// @Produce json
// @Success 200 {array} model.Appointment
// @Router /appointments/getAll [get]
func (h *Handler) getAllAppointments(c *gin.Context) {
	res, err := h.repo.All.GetAllAppointments()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Get an appointment by ID
// @Description Get an appointment based on the provided ID
// @Tags appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} model.Appointment
// @Router /appointments/get/{id} [get]
func (h *Handler) getAppointments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	res, err := h.repo.All.GetAppointment(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary Create a new appointment
// @Description Create a new appointment with the provided data
// @Tags appointments
// @Produce json
// @Consumes json
// @Param appointment body model.Appointment true "Appointment object"
// @Router /appointments/create [post]
func (h *Handler) createAppointments(c *gin.Context) {
	req := &model.Appointment{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	id, err := h.repo.CreateAppointment(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"id": id})
}

// @Summary Update an appointment
// @Description Update an appointment with the provided data
// @Tags appointments
// @Produce json
// @Consumes json
// @Param appointment body model.Appointment true "Appointment object"
// @Router /appointments/update [post]
func (h *Handler) updateAppointments(c *gin.Context) {
	req := &model.Appointment{}
	if err := c.BindJSON(req); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err := h.repo.UpdateAppointment(req)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}

// @Summary Delete an appointment
// @Description Delete an appointment based on the provided ID
// @Tags appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Router /appointments/delete/{id} [get]
func (h *Handler) deleteAppointments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.repo.DeleteAppointment(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	c.JSON(200, map[string]interface{}{"status": "ok"})
}
