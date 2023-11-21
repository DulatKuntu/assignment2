package repository

import (
	"bonus/model"
	"database/sql"
)

type All interface {
	GetAllUsers() ([]*model.User, error)
	GetUser(id int) (*model.User, error)
	CreateUser(user *model.User) (int, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error

	GetAllCaregivers() ([]*model.Caregiver, error)
	GetCaregiver(id int) (*model.Caregiver, error)
	CreateCaregiver(caregiver *model.Caregiver) (int, error)
	UpdateCaregiver(caregiver *model.Caregiver) error
	DeleteCaregiver(id int) error

	GetAllMembers() ([]*model.Member, error)
	GetMember(id int) (*model.Member, error)
	CreateMember(member *model.Member) (int, error)
	UpdateMember(member *model.Member) error
	DeleteMember(id int) error

	GetAllAddresses() ([]*model.Address, error)
	GetAddress(id int) (*model.Address, error)
	CreateAddress(address *model.Address) (int, error)
	UpdateAddress(address *model.Address) error
	DeleteAddress(id int) error

	GetAllJobs() ([]*model.Job, error)
	GetJob(id int) (*model.Job, error)
	CreateJob(job *model.Job) (int, error)
	UpdateJob(job *model.Job) error
	DeleteJob(id int) error

	GetAllJobApplications() ([]*model.JobApplication, error)
	GetJobApplication(id int) (*model.JobApplication, error)
	CreateJobApplication(jobApplication *model.JobApplication) (int, error)
	UpdateJobApplication(jobApplication *model.JobApplication) error
	DeleteJobApplication(id int) error

	GetAllAppointments() ([]*model.Appointment, error)
	GetAppointment(id int) (*model.Appointment, error)
	CreateAppointment(appointment *model.Appointment) (int, error)
	UpdateAppointment(appointment *model.Appointment) error
	DeleteAppointment(id int) error
}

type Repository struct {
	All
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		All: NewAllRepository(db),
	}
}
