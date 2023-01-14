-- Table: public.user

-- DROP TABLE public."user";

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS public."user"
(
    id bigint NOT NULL DEFAULT nextval('user_id_seq'::regclass),
    name varchar(15) NOT NULL, /* Nick */
    phone_number varchar(11) NOT NULL UNIQUE , /* 79999999999 */
    email varchar(30) NOT NULL UNIQUE , /* Nick@gmail.com */
    registration_date date NOT NULL DEFAULT now(),
    password_hash varchar(60) NOT NULL, /* BCrypt hash length */
    CONSTRAINT user_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE public."user"
    OWNER to nick;

/* Random data for test */

INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Diane Bender', 52707206228, 'kellybrooks@wilson-simmons.com', '$2b$12$MDh3u4FHBriVttpzq48f/udpUMTs39B7hLMKeuSsNejDswe0hW4k.');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Michelle Brooks', 00142360597, 'mary14@hotmail.com', '$2b$12$wEMYIlWeKvu1gO98V2IVbeKy9Sh0M0zaq.0DNaynSSS6DAR70kDwm');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Julie Ibarra', 00160163919, 'ericsmith@rodriguez.net', '$2b$12$NBkfBifEoFYPuMXCPAaICucqsMttVXM/n5SOZzGuEIFwDgFy77iF.');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Amy Lopez', 18499766332, 'maymadison@gmail.com', '$2b$12$pvMNQbBeWTYc1sa11KS7NehEZAAeJf8n7pgAQPKzgFpF6qPOV/sFm');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Tyler Watts', 18555865949, 'xgarner@yahoo.com', '$2b$12$LgFaPCYy0LMv7sS0pvk3HuBKb.mjM83W9x/nEuyFgfsbIpPpyrOSu');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Rebecca Cox', 99836023511, 'zwang@tyler-moore.com', '$2b$12$22JdHI9CUOjqF9Lc5ZZlquJ/r6wq1YFPL.y/71TINuSjGp3eh2O/e');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Nichole Jackson', 00118511377, 'owatson@mcguire.com', '$2b$12$.R79uVxFwdjgDbLkbt2V1ec7uDC5l3PVArfiZ.cV1hPDeLD0qljzO');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Steven Williams', 13396861976, 'cewing@gmail.com', '$2b$12$P3MnfZdggI5Pdj7FHNSTS.Lk3alTnpSLw81UHRQBOWrHSJZ40sUqi');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Benjamin Knight', 16653967395, 'stephanienash@davis.biz', '$2b$12$J/Y6tfBYeRGSNLy1G8WmKuMi4MGBJcZphQARRTnmn7ImFzudcmENW');
INSERT INTO public."user"(	name, phone_number, email, password_hash)	VALUES ('Diana Haynes', 20352487837, 'dharrington@silva.com', '$2b$12$5vSX/N9FA1wxbhKjixZqiePLYjZWCdogzeyA9PRqG2fl.Tae92ZVa');