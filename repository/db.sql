-- SCHEMA: test

-- DROP SCHEMA test ;

CREATE SCHEMA test
    AUTHORIZATION postgres;

-- Table: test.permission

-- DROP TABLE test.permission;

CREATE TABLE test.permission
(
    permission_id integer NOT NULL DEFAULT nextval('test.permission_permission_id_seq'::regclass),
    name character varying(45) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT permission_pkey PRIMARY KEY (permission_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE test.permission
    OWNER to postgres;

-- Table: test.role

-- DROP TABLE test.role;

CREATE TABLE test.role
(
    role_id integer NOT NULL DEFAULT nextval('test.role_role_id_seq'::regclass),
    name character varying(45) COLLATE pg_catalog."default",
    CONSTRAINT role_pkey PRIMARY KEY (role_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE test.role
    OWNER to postgres;

    -- Table: test.role_permission

-- DROP TABLE test.role_permission;

CREATE TABLE test.role_permission
(
    role_id integer NOT NULL,
    permission_id integer NOT NULL,
    CONSTRAINT role_permission_pkey PRIMARY KEY (role_id, permission_id),
    CONSTRAINT rel_permission_role_permission FOREIGN KEY (permission_id)
        REFERENCES test.permission (permission_id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT,
    CONSTRAINT rel_role_role_permission FOREIGN KEY (role_id)
        REFERENCES test.role (role_id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE test.role_permission
    OWNER to postgres;

    -- Table: test."user"

-- DROP TABLE test."user";

CREATE TABLE test."user"
(
    user_id integer NOT NULL DEFAULT nextval('test.user_user_id_seq'::regclass),
    email character varying(50) COLLATE pg_catalog."default" NOT NULL,
    password character varying(25) COLLATE pg_catalog."default" NOT NULL,
    name character varying(45) COLLATE pg_catalog."default" NOT NULL,
    surname character varying(45) COLLATE pg_catalog."default" NOT NULL,
    role_id integer NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (user_id),
    CONSTRAINT uq_email UNIQUE (email)
,
    CONSTRAINT rel_user_role FOREIGN KEY (role_id)
        REFERENCES test.role (role_id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE test."user"
    OWNER to postgres;