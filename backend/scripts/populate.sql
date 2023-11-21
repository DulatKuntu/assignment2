INSERT INTO users (email, given_name, surname, city, phone_number, profile_description, password)
VALUES
    ('user1@example.com', 'Askar', 'Askarov', 'City1', '1234567890', 'Profile 1', 'password1'),
    ('user2@example.com', 'Bolat', 'Bolatov', 'City2', '9876543210', 'Profile 2', 'password2'),
    ('user3@example.com', 'Bob', 'Johnson', 'Astana', '5551234567', 'Profile 3', 'password3'),
    ('user4@example.com', 'Alice', 'Williams', 'City4', '1112233445', 'Profile 4', 'password4'),
    ('user5@example.com', 'Charlie', 'Brown', 'City5', '9998887776', 'Profile 5', 'password5'),
    ('user6@example.com', 'Eva', 'Anderson', 'City6', '4445556668', 'Profile 6', 'password6'),
    ('user7@example.com', 'David', 'Lee', 'City7', '7776665554', 'Profile 7', 'password7'),
    ('user8@example.com', 'Grace', 'Smith', 'City8', '6665554443', 'Profile 8', 'password8'),
    ('user9@example.com', 'Sam', 'Taylor', 'City9', '3334445552', 'Profile 9', 'password9'),
    ('user10@example.com', 'Lily', 'Johnson', 'City10', '2223334441', 'Profile 10', 'password10'),
    ('user11@example.com', 'John', 'Doe', 'City11', '5556667778', 'Profile 11', 'password11'),
    ('user12@example.com', 'Emma', 'Miller', 'City12', '7778889990', 'Profile 12', 'password12'),
    ('user13@example.com', 'Daniel', 'White', 'City13', '8889990001', 'Profile 13', 'password13'),
    ('user14@example.com', 'Sophia', 'Jones', 'City14', '2221110002', 'Profile 14', 'password14'),
    ('user15@example.com', 'William', 'Moore', 'City15', '4443332223', 'Profile 15', 'password15'),
    ('user16@example.com', 'Olivia', 'Brown', 'City16', '6665554444', 'Profile 16', 'password16'),
    ('user17@example.com', 'James', 'Taylor', 'City17', '8887776665', 'Profile 17', 'password17'),
    ('user18@example.com', 'Ava', 'Smith', 'City18', '1110009996', 'Profile 18', 'password18'),
    ('user19@example.com', 'Benjamin', 'Clark', 'City19', '3332221117', 'Profile 19', 'password19'),
    ('user20@example.com', 'Mia', 'Adams', 'City20', '5554443338', 'Profile 20', 'password20');

INSERT INTO caregivers (caregiver_user_id,photo, gender, caregiving_type, hourly_rate)
VALUES
    (11,'photo1.jpg', 'Male', 'Baby Sitter', 20),
    (12,'photo2.jpg', 'Female', 'Home Health Aide', 14),
    (13,'photo3.jpg', 'Male', 'Nurse', 7),
    (14,'photo4.jpg', 'Female', 'Therapist', 35),
    (15,'photo5.jpg', 'Male', 'Baby Sitter', 2),
    (16,'photo6.jpg', 'Female', 'Home Health Aide', 28),
    (17,'photo7.jpg', 'Male', 'Baby Sitter', 9),
    (18,'photo8.jpg', 'Female', 'Therapist', 10),
    (19,'photo9.jpg', 'Male', 'Baby Sitter', 9.5),
    (20,'photo10.jpg', 'Female', 'Home Health Aide', 6);

INSERT INTO members (member_user_id,house_rules)
VALUES
    (1,'No smoking in the house.'),
    (2,'Quiet hours after 10 PM.'),
    (3,'No pets.'),
    (4,'Clean up after using the shared spaces.'),
    (5,'Respect others privacy.'),
    (6,'Contribute to shared expenses.'),
    (7,'Limit guests to weekends only.'),
    (8,'Use headphones for loud music or TV.'),
    (9,'Participate in monthly house meetings.'),
    (10,'Notify in advance about overnight guests.');

INSERT INTO addresses (member_user_id,house_number, street, town)
VALUES
    (1,123, 'Kabanbay', 'City1'),
    (2,456, 'Kabanbay', 'City2'),
    (3,789, 'Kabanbay', 'Astana'),
    (4,101, 'Kabanbay', 'City4'),
    (5,202, 'Kabanbay', 'City5'),
    (6,303, 'Turan', 'City6'),
    (7,404, 'Kabanbay', 'City7'),
    (8,505, 'Turan', 'City8'),
    (9,606, 'Kabanbay', 'City9'),
    (10,707, 'Kabanbay', 'City10');

INSERT INTO jobs (member_user_id, required_caregiving_type, other_requirements, date_posted)
VALUES
    (1, 'Companion', 'Experience with elderly care preferred.', '2023-01-01'),
    (2, 'Home Health Aide', 'CPR certification required.', '2023-01-02'),
    (3, 'Nurse', 'Valid nursing license required.', '2023-01-03'),
    (4, 'Therapist', 'Experience in physical therapy required.', '2023-01-04'),
    (5, 'Companion', 'Experience with elderly care preferred.', '2023-01-05'),
    (6, 'Home Health Aide', 'CPR certification required.', '2023-01-06'),
    (7, 'Nurse', 'gentle', '2023-01-07'),
    (8, 'Therapist', 'Experience in physical therapy required, gentle', '2023-01-08'),
    (9, 'Companion', 'Experience with elderly care preferred, gentle', '2023-01-09'),
    (10, 'Home Health Aide', 'CPR certification required.', '2023-01-10');

INSERT INTO job_applications (caregiver_user_id, job_id, date_applied)
VALUES
    (11, 1, '2023-02-01'),
    (12, 2, '2023-02-02'),
    (13, 3, '2023-02-03'),
    (14, 4, '2023-02-04'),
    (15, 5, '2023-02-05'),
    (16, 6, '2023-02-06'),
    (17, 7, '2023-02-07'),
    (18, 8, '2023-02-08'),
    (19, 9, '2023-02-09'),
    (20, 10, '2023-02-10'),
    (12, 1, '2023-02-01'),
    (14, 9, '2023-02-09');

INSERT INTO appointments (caregiver_user_id, member_user_id, appointment_date, appointment_time, work_hours, status)
VALUES
    (11, 1, '2023-03-01', '08:00:00', 4, 'Scheduled'),
    (12, 2, '2023-03-02', '09:30:00', 5, 'Scheduled'),
    (13, 3, '2023-03-03', '10:15:00', 6, 'Scheduled'),
    (14, 4, '2023-03-04', '11:00:00', 7, 'Scheduled'),
    (15, 5, '2023-03-05', '13:30:00', 8, 'Accepted'),
    (16, 6, '2023-03-06', '14:45:00', 4, 'Scheduled'),
    (17, 7, '2023-03-07', '16:00:00', 5, 'Scheduled'),
    (18, 8, '2023-03-08', '17:15:00', 6, 'Accepted'),
    (19, 9, '2023-03-09', '18:30:00', 7, 'Scheduled'),
    (20, 10, '2023-03-10', '19:45:00', 8, 'Accepted');
