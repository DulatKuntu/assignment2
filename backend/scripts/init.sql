CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    given_name VARCHAR(255),
    surname VARCHAR(255),
    city VARCHAR(255),
    phone_number VARCHAR(20),
    profile_description TEXT,
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS caregivers (
    caregiver_user_id SERIAL PRIMARY KEY,
    photo VARCHAR(255),
    gender VARCHAR(10),
    caregiving_type VARCHAR(255),
    hourly_rate DECIMAL,
    FOREIGN KEY (caregiver_user_id) REFERENCES users(user_id)  ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS members (
    member_user_id SERIAL PRIMARY KEY,
    house_rules TEXT,
    FOREIGN KEY (member_user_id) REFERENCES users(user_id)  ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS addresses (
    member_user_id SERIAL PRIMARY KEY,
    house_number INTEGER,
    street VARCHAR(255),
    town VARCHAR(255),
    FOREIGN KEY (member_user_id) REFERENCES users(user_id)  ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS jobs (
    job_id SERIAL PRIMARY KEY,
    member_user_id INTEGER,
    required_caregiving_type VARCHAR(255),
    other_requirements TEXT,
    date_posted DATE,
    FOREIGN KEY (member_user_id) REFERENCES users(user_id)  ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS job_applications (
    caregiver_user_id INTEGER,
    job_id INTEGER,
    date_applied DATE,
    PRIMARY KEY (caregiver_user_id, job_id),
    FOREIGN KEY (caregiver_user_id) REFERENCES users(user_id)  ON DELETE CASCADE,
    FOREIGN KEY (job_id) REFERENCES jobs(job_id)  ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS appointments (
    appointment_id SERIAL PRIMARY KEY,
    caregiver_user_id INTEGER,
    member_user_id INTEGER,
    appointment_date DATE,
    appointment_time TIME,
    work_hours INTEGER,
    status VARCHAR(20),
    FOREIGN KEY (caregiver_user_id) REFERENCES users(user_id)  ON DELETE CASCADE,
    FOREIGN KEY (member_user_id) REFERENCES users(user_id)  ON DELETE CASCADE
);