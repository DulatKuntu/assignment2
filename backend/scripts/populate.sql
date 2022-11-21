INSERT INTO diseaseType(id, description) VALUES (1, 'Description1');
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('COVID-19', 'bacteria', 'Description for coronavirus', 1);
INSERT INTO country(cname, population) VALUES ('China', 1412000000);
INSERT INTO discover(cname, disease_code, first_enc_date) VALUES ('China', 'COVID-19', '1980-05-28');

INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('Tuberculosis', 'bacteria', 'Description for tuberculosis', 1);
INSERT INTO country(cname, population) VALUES ('Qazaqstan', 18000000);
INSERT INTO discover(cname, disease_code, first_enc_date) VALUES ('Qazaqstan', 'Tuberculosis', '1960-05-28');

INSERT INTO diseaseType(id, description) VALUES (2, 'Description2');
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('Pneumonia', 'virus', 'Description for Pneumonia', 2);
INSERT INTO discover(cname, disease_code, first_enc_date) VALUES ('Qazaqstan', 'Pneumonia', '2000-05-28');

INSERT INTO diseaseType(id, description) VALUES (3, 'infectious diseases');
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('Infectious disease', 'virus', 'Description for Infectious disease', 3);

INSERT INTO country(cname, population) VALUES ('Italy', 62000000);

INSERT INTO diseaseType(id, description) VALUES (4, 'virology');
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('HIV', 'virus', 'Description for HIV', 4);
INSERT INTO country(cname, population) VALUES ('USA', 330000000);

INSERT INTO country(cname, population) VALUES ('Russia', 130000000);
INSERT INTO country(cname, population) VALUES ('Turkey', 150000000);
INSERT INTO country(cname, population) VALUES ('Germany', 200000000);
INSERT INTO country(cname, population) VALUES ('Spain', 189000000);
INSERT INTO country(cname, population) VALUES ('France', 200000000);
INSERT INTO country(cname, population) VALUES ('UK', 100000000);

INSERT INTO diseaseType(id, description) VALUES (5, 'fungus');
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('fungus1', 'virus', 'Description for fungus1', 5);
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('fungus2', 'virus', 'Description for fungus1', 5);

INSERT INTO diseaseType(id, description) VALUES (6, 'respiratory');
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('asthma', 'bacteria', 'Description for asthma', 6);
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('flu', 'virus', 'Description for flu', 6);

INSERT INTO diseaseType(id, description) VALUES (7, 'digestive');
INSERT INTO disease(disease_code, pathogen, description, id) VALUES ('gastritus', 'bacteria', 'Description for gastritus', 7);



INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor1@gmail.com', 'Dname1', 'Sname1', 300000, '87770002525', 'Qazaqstan');
INSERT INTO doctor(email, degree) VALUES ('doctor1@gmail.com', 'MD');
INSERT INTO specialize(id, email) VALUES (1, 'doctor1@gmail.com'); 

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor2@gmail.com', 'Dname2', 'Sname2', 400000, '87772222222', 'China');
INSERT INTO doctor(email, degree) VALUES ('doctor2@gmail.com', 'PhD');
INSERT INTO specialize(id, email) VALUES (2, 'doctor2@gmail.com');

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor3@gmail.com', 'Dname3', 'Sname3', 500000, '87773333333', 'Italy');
INSERT INTO doctor(email, degree) VALUES ('doctor3@gmail.com', 'PhD');
INSERT INTO specialize(id, email) VALUES (3, 'doctor3@gmail.com');

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor4@gmail.com', 'Dname4', 'Sname4', 400000, '87774444444', 'Qazaqstan');
INSERT INTO doctor(email, degree) VALUES ('doctor4@gmail.com', 'PhD');
INSERT INTO specialize(id, email) VALUES (4, 'doctor4@gmail.com'); 

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor5@gmail.com', 'Dname5', 'Sname5', 400000, '87775555555', 'China');
INSERT INTO doctor(email, degree) VALUES ('doctor5@gmail.com', 'PhD');
INSERT INTO specialize(id, email) VALUES (4, 'doctor5@gmail.com');

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor6@gmail.com', 'Dname6', 'Sname6', 300000, '87776666666', 'Italy');
INSERT INTO doctor(email, degree) VALUES ('doctor6@gmail.com', 'MD');
INSERT INTO specialize(id, email) VALUES (4, 'doctor6@gmail.com');

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor7@gmail.com', 'Dname7', 'Sname7', 3000000, '87777777777', 'Italy');
INSERT INTO doctor(email, degree) VALUES ('doctor7@gmail.com', 'PhD');
INSERT INTO specialize(id, email) VALUES (4, 'doctor7@gmail.com');

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor8@gmail.com', 'Dname8', 'Sname8', 400000, '87770002525', 'Spain');
INSERT INTO doctor(email, degree) VALUES ('doctor8@gmail.com', 'MD');
INSERT INTO specialize(id, email) VALUES (7, 'doctor8@gmail.com'); 
INSERT INTO public.publicservant(email, department)VALUES ('doctor8@gmail.com', 'department3');
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('doctor8@gmail.com', 'Russia', 'fungus1', 10000000, 53400000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('doctor8@gmail.com', 'China', 'gastritus', 1000000, 160700000);

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('doctor9@gmail.com', 'Dname9', 'Sname9', 400000, '87770002525', 'UK');
INSERT INTO doctor(email, degree) VALUES ('doctor9@gmail.com', 'PhD');
INSERT INTO specialize(id, email) VALUES (7, 'doctor9@gmail.com'); 
INSERT INTO public.publicservant(email, department)VALUES ('doctor9@gmail.com', 'department4');
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('doctor9@gmail.com', 'UK', 'fungus1', 10000000, 534000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('doctor9@gmail.com', 'France', 'gastritus', 1005000, 1607000);

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('public1@gmail.com', 'namePublic1', 'SnamePublic1', 3000000, '81111111111', 'Italy');
INSERT INTO public.publicservant(email, department)VALUES ('public1@gmail.com', 'department1');
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public1@gmail.com', 'Italy', 'COVID-19', 200000, 20000000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public1@gmail.com', 'Qazaqstan', 'COVID-19', 100000, 534000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public1@gmail.com', 'USA', 'COVID-19', 100000, 160000);


INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('public2@gmail.com', 'namePublic2', 'SnamePublic2', 600000, '82222222222', 'China');
INSERT INTO public.publicservant(email, department)VALUES ('public2@gmail.com', 'department1');
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public2@gmail.com', 'China', 'COVID-19', 2000000, 200000000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public2@gmail.com', 'Qazaqstan', 'COVID-19', 100000, 1600000);
INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('public4@gmail.com', 'namePublic4', 'SnamePublic4', 600000, '82222222222', 'Italy');
INSERT INTO public.publicservant(email, department)VALUES ('public4@gmail.com', 'department1');
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public4@gmail.com', 'Turkey', 'COVID-19', 185000, 6500000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public4@gmail.com', 'Germany', 'COVID-19', 9805000, 98200000);
INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('public5@gmail.com', 'namePublic5', 'SnamePublic5', 60000, '82222222222', 'Russia');
INSERT INTO public.publicservant(email, department)VALUES ('public5@gmail.com', 'department3');
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public5@gmail.com', 'Turkey', 'COVID-19', 185000, 6500000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public5@gmail.com', 'Spain', 'COVID-19', 8952, 289500);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public5@gmail.com', 'Qazaqstan', 'HIV', 9657, 12500);

INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('public3@gmail.com', 'namePublic3', 'SnamePublic3', 600000, '83333333333', 'Qazaqstan');
INSERT INTO public.publicservant(email, department)VALUES ('public3@gmail.com', 'department2');
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public3@gmail.com', 'Qazaqstan', 'COVID-19', 100000, 160000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public3@gmail.com', 'USA', 'COVID-19', 2000000, 200000000);
INSERT INTO public.record(email, cname, disease_code, total_deaths, total_patients) VALUES ('public3@gmail.com', 'Russia', 'COVID-19', 16800, 1895000);


INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('alibek@gmail.com', 'alibek', 'alibek', 600000, '83333333333', 'Qazaqstan');
INSERT INTO users(email, name, surname, salary, phone, cname) VALUES ('gulsim@gmail.com', 'gulsim', 'gulsim', 600000, '83333333333', 'Qazaqstan');
