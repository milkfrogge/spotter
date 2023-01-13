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
    phone_number varchar(11) NOT NULL, /* 79999999999 */
    email varchar(30) NOT NULL, /* Nick@gmail.com */
    registration_date date NOT NULL DEFAULT now(),
    password_hash varchar(60) NOT NULL, /* BCrypt hash length */
    CONSTRAINT user_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE public."user"
    OWNER to nick;