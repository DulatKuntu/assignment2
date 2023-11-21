package model

import "time"

// User represents the 'users' table
type User struct {
	UserID             uint   `json:"user_id"`
	Email              string `json:"email"`
	GivenName          string `json:"given_name"`
	Surname            string `json:"surname"`
	City               string `json:"city"`
	PhoneNumber        string `json:"phone_number"`
	ProfileDescription string `json:"profile_description"`
	Password           string `json:"password"`
}

// Caregiver represents the 'caregivers' table
type Caregiver struct {
	CaregiverUserID uint    `json:"caregiver_user_id"`
	Photo           string  `json:"photo"`
	Gender          string  `json:"gender"`
	CaregivingType  string  `json:"caregiving_type"`
	HourlyRate      float64 `json:"hourly_rate"`
}

// Member represents the 'members' table
type Member struct {
	MemberUserID uint   `json:"member_user_id"`
	HouseRules   string `json:"house_rules"`
}

// Address represents the 'addresses' table
type Address struct {
	MemberUserID uint   `json:"member_user_id"`
	HouseNumber  int    `json:"house_number"`
	Street       string `json:"street"`
	Town         string `json:"town"`
}

// Job represents the 'jobs' table
type Job struct {
	JobID                  uint      `json:"job_id"`
	MemberUserID           uint      `json:"member_user_id"`
	RequiredCaregivingType string    `json:"required_caregiving_type"`
	OtherRequirements      string    `json:"other_requirements"`
	DatePosted             time.Time `json:"date_posted"`
}

// JobApplication represents the 'job_applications' table
type JobApplication struct {
	CaregiverUserID uint      `json:"caregiver_user_id"`
	JobID           uint      `json:"job_id"`
	DateApplied     time.Time `json:"date_applied"`
}

// Appointment represents the 'appointments' table
type Appointment struct {
	AppointmentID   uint      `json:"appointment_id"`
	CaregiverUserID uint      `json:"caregiver_user_id"`
	MemberUserID    uint      `json:"member_user_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	WorkHours       int       `json:"work_hours"`
	Status          string    `json:"status"`
}
