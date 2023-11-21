package repository

import (
	"bonus/model"
	"database/sql"
)

type AllDB struct {
	db *sql.DB
}

func (r *AllDB) GetAllUsers() ([]*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		user := new(model.User)
		err := rows.Scan(&user.UserID, &user.Email, &user.GivenName, &user.Surname, &user.City, &user.PhoneNumber, &user.ProfileDescription, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AllDB) GetUser(id int) (*model.User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE user_id=$1", id)

	user := new(model.User)
	err := row.Scan(&user.UserID, &user.Email, &user.GivenName, &user.Surname, &user.City, &user.PhoneNumber, &user.ProfileDescription, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AllDB) CreateUser(user *model.User) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO users (email, given_name, surname, city, phone_number, profile_description, password) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id", user.Email, user.GivenName, user.Surname, user.City, user.PhoneNumber, user.ProfileDescription, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AllDB) UpdateUser(user *model.User) error {
	_, err := r.db.Exec("UPDATE users SET email=$1, given_name=$2, surname=$3, city=$4, phone_number=$5, profile_description=$6, password=$7 WHERE user_id=$8", user.Email, user.GivenName, user.Surname, user.City, user.PhoneNumber, user.ProfileDescription, user.Password, user.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *AllDB) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE user_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AllDB) GetAllCaregivers() ([]*model.Caregiver, error) {
	rows, err := r.db.Query("SELECT * FROM caregivers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	caregivers := make([]*model.Caregiver, 0)
	for rows.Next() {
		caregiver := new(model.Caregiver)
		err := rows.Scan(&caregiver.CaregiverUserID, &caregiver.Photo, &caregiver.Gender, &caregiver.CaregivingType, &caregiver.HourlyRate)
		if err != nil {
			return nil, err
		}
		caregivers = append(caregivers, caregiver)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return caregivers, nil
}

func (r *AllDB) GetCaregiver(id int) (*model.Caregiver, error) {
	row := r.db.QueryRow("SELECT * FROM caregivers WHERE caregiver_user_id=$1", id)

	caregiver := new(model.Caregiver)

	err := row.Scan(&caregiver.CaregiverUserID, &caregiver.Photo, &caregiver.Gender, &caregiver.CaregivingType, &caregiver.HourlyRate)
	if err != nil {
		return nil, err
	}

	return caregiver, nil
}

func (r *AllDB) CreateCaregiver(caregiver *model.Caregiver) (int, error) {
	var id int

	err := r.db.QueryRow("INSERT INTO caregivers (caregiver_user_id, photo, gender, caregiving_type, hourly_rate) VALUES ($1, $2, $3, $4, $5) RETURNING caregiver_user_id", caregiver.CaregiverUserID, caregiver.Photo, caregiver.Gender, caregiver.CaregivingType, caregiver.HourlyRate).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AllDB) UpdateCaregiver(caregiver *model.Caregiver) error {
	_, err := r.db.Exec("UPDATE caregivers SET photo=$1, gender=$2, caregiving_type=$3, hourly_rate=$4 WHERE caregiver_user_id=$5", caregiver.Photo, caregiver.Gender, caregiver.CaregivingType, caregiver.HourlyRate, caregiver.CaregiverUserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) DeleteCaregiver(id int) error {
	_, err := r.db.Exec("DELETE FROM caregivers WHERE caregiver_user_id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
func (r *AllDB) GetAllMembers() ([]*model.Member, error) {
	rows, err := r.db.Query("SELECT * FROM members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make([]*model.Member, 0)
	for rows.Next() {
		member := new(model.Member)
		err := rows.Scan(&member.MemberUserID, &member.HouseRules)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return members, nil
}

func (r *AllDB) GetMember(id int) (*model.Member, error) {
	row := r.db.QueryRow("SELECT * FROM members WHERE member_user_id=$1", id)

	member := new(model.Member)
	err := row.Scan(&member.MemberUserID, &member.HouseRules)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (r *AllDB) CreateMember(member *model.Member) (int, error) {
	var id int

	err := r.db.QueryRow("INSERT INTO members (member_user_id, house_rules) VALUES ($1, $2) RETURNING member_user_id", member.MemberUserID, member.HouseRules).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AllDB) UpdateMember(member *model.Member) error {
	_, err := r.db.Exec("UPDATE members SET house_rules=$1 WHERE member_user_id=$2", member.HouseRules, member.MemberUserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) DeleteMember(id int) error {
	_, err := r.db.Exec("DELETE FROM members WHERE member_user_id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) GetAllAddresses() ([]*model.Address, error) {
	rows, err := r.db.Query("SELECT * FROM addresses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := make([]*model.Address, 0)
	for rows.Next() {
		address := new(model.Address)
		err := rows.Scan(&address.MemberUserID, &address.HouseNumber, &address.Street, &address.Town)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return addresses, nil
}

func (r *AllDB) GetAddress(id int) (*model.Address, error) {
	row := r.db.QueryRow("SELECT * FROM addresses WHERE member_user_id=$1", id)

	address := new(model.Address)
	err := row.Scan(&address.MemberUserID, &address.HouseNumber, &address.Street, &address.Town)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (r *AllDB) CreateAddress(address *model.Address) (int, error) {
	var id int

	err := r.db.QueryRow("INSERT INTO addresses (member_user_id, house_number, street, town) VALUES ($1, $2, $3, $4) RETURNING member_user_id", address.MemberUserID, address.HouseNumber, address.Street, address.Town).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AllDB) UpdateAddress(address *model.Address) error {
	_, err := r.db.Exec("UPDATE addresses SET house_number=$1, street=$2, town=$3 WHERE member_user_id=$4", address.HouseNumber, address.Street, address.Town, address.MemberUserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) DeleteAddress(id int) error {
	_, err := r.db.Exec("DELETE FROM addresses WHERE member_user_id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) GetAllJobs() ([]*model.Job, error) {
	rows, err := r.db.Query("SELECT * FROM jobs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobs := make([]*model.Job, 0)
	for rows.Next() {
		job := new(model.Job)
		err := rows.Scan(&job.JobID, &job.MemberUserID, &job.RequiredCaregivingType, &job.OtherRequirements, &job.DatePosted)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *AllDB) GetJob(id int) (*model.Job, error) {
	row := r.db.QueryRow("SELECT * FROM jobs WHERE job_id=$1", id)

	job := new(model.Job)
	err := row.Scan(&job.JobID, &job.MemberUserID, &job.RequiredCaregivingType, &job.OtherRequirements, &job.DatePosted)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (r *AllDB) CreateJob(job *model.Job) (int, error) {
	var id int

	err := r.db.QueryRow("INSERT INTO jobs (member_user_id, required_caregiving_type, other_requirements, date_posted) VALUES ($1, $2, $3, $4) RETURNING job_id", job.MemberUserID, job.RequiredCaregivingType, job.OtherRequirements, job.DatePosted).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AllDB) UpdateJob(job *model.Job) error {
	_, err := r.db.Exec("UPDATE jobs SET member_user_id=$1, required_caregiving_type=$2, other_requirements=$3, date_posted=$4 WHERE job_id=$5", job.MemberUserID, job.RequiredCaregivingType, job.OtherRequirements, job.DatePosted, job.JobID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) DeleteJob(id int) error {
	_, err := r.db.Exec("DELETE FROM jobs WHERE job_id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) GetAllJobApplications() ([]*model.JobApplication, error) {
	rows, err := r.db.Query("SELECT * FROM job_applications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobApplications := make([]*model.JobApplication, 0)
	for rows.Next() {
		jobApplication := new(model.JobApplication)
		err := rows.Scan(&jobApplication.CaregiverUserID, &jobApplication.JobID, &jobApplication.DateApplied)
		if err != nil {
			return nil, err
		}
		jobApplications = append(jobApplications, jobApplication)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return jobApplications, nil
}

func (r *AllDB) GetJobApplication(id int) (*model.JobApplication, error) {
	row := r.db.QueryRow("SELECT * FROM job_applications WHERE caregiver_user_id=$1", id)

	jobApplication := new(model.JobApplication)
	err := row.Scan(&jobApplication.CaregiverUserID, &jobApplication.JobID, &jobApplication.DateApplied)
	if err != nil {
		return nil, err
	}

	return jobApplication, nil
}

func (r *AllDB) CreateJobApplication(jobApplication *model.JobApplication) (int, error) {
	var id int

	err := r.db.QueryRow("INSERT INTO job_applications (caregiver_user_id, job_id, date_applied) VALUES ($1, $2, $3) RETURNING caregiver_user_id", jobApplication.CaregiverUserID, jobApplication.JobID, jobApplication.DateApplied).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AllDB) UpdateJobApplication(jobApplication *model.JobApplication) error {
	_, err := r.db.Exec("UPDATE job_applications SET job_id=$1, date_applied=$2 WHERE caregiver_user_id=$3", jobApplication.JobID, jobApplication.DateApplied, jobApplication.CaregiverUserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) DeleteJobApplication(id int) error {
	_, err := r.db.Exec("DELETE FROM job_applications WHERE caregiver_user_id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) GetAllAppointments() ([]*model.Appointment, error) {
	rows, err := r.db.Query("SELECT * FROM appointments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments := make([]*model.Appointment, 0)
	for rows.Next() {
		appointment := new(model.Appointment)
		err := rows.Scan(&appointment.AppointmentID, &appointment.CaregiverUserID, &appointment.MemberUserID, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.WorkHours, &appointment.Status)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AllDB) GetAppointment(id int) (*model.Appointment, error) {
	row := r.db.QueryRow("SELECT * FROM appointments WHERE appointment_id=$1", id)

	appointment := new(model.Appointment)
	err := row.Scan(&appointment.AppointmentID, &appointment.CaregiverUserID, &appointment.MemberUserID, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.WorkHours, &appointment.Status)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (r *AllDB) CreateAppointment(appointment *model.Appointment) (int, error) {
	var id int

	err := r.db.QueryRow("INSERT INTO appointments (caregiver_user_id, member_user_id, appointment_date, appointment_time, work_hours, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING appointment_id", appointment.CaregiverUserID, appointment.MemberUserID, appointment.AppointmentDate, appointment.AppointmentTime, appointment.WorkHours, appointment.Status).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AllDB) UpdateAppointment(appointment *model.Appointment) error {
	_, err := r.db.Exec("UPDATE appointments SET caregiver_user_id=$1, member_user_id=$2, appointment_date=$3, appointment_time=$4, work_hours=$5, status=$6 WHERE appointment_id=$7", appointment.CaregiverUserID, appointment.MemberUserID, appointment.AppointmentDate, appointment.AppointmentTime, appointment.WorkHours, appointment.Status, appointment.AppointmentID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AllDB) DeleteAppointment(id int) error {
	_, err := r.db.Exec("DELETE FROM appointments WHERE appointment_id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func NewAllRepository(db *sql.DB) *AllDB {
	return &AllDB{db: db}
}
