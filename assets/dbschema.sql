CREATE TABLE tpermission(
  tpe_id INTEGER NOT NULL PRIMARY KEY,
  tpe_name TEXT NOT NULL
);

CREATE TABLE trole(
    tro_id INTEGER NOT NULL PRIMARY KEY,
    tro_name TEXT NOT NULL
);

CREATE TABLE rrol_permission(
    rrp_role_id INTEGER NOT NULL,
    rrp_permission_id INTEGER NOT NULL,
    CONSTRAINT rel_permission_role_permission FOREIGN KEY (rrp_permission_id)
        REFERENCES tpermission (tpe_id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT,
    CONSTRAINT rel_role_role_permission FOREIGN KEY (rrp_role_id)
        REFERENCES trole (tro_id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT

);

CREATE TABLE tuser(
    tus_id INTEGER NOT NULL PRIMARY KEY,
    tus_email TEXT NOT NULL,
    tus_password TEXT NOT NULL,
    tus_name TEXT NOT NULL,
    tus_surname TEXT NOT NULL,
    tus_role_id integer NOT NULL,
    CONSTRAINT uq_email UNIQUE (tus_email),
    CONSTRAINT rel_user_role FOREIGN KEY (tus_role_id)
        REFERENCES trole (tro_id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
);

CREATE UNIQUE INDEX uq_email ON tuser(tus_email);


INSERT INTO tpermission(tpe_name)
VALUES('EditUsername');

INSERT INTO tpermission(tpe_name)
VALUES('Editemail');

INSERT INTO trole(tro_name)
VALUES('dummy');

INSERT INTO rrol_permission(rrp_role_id, rrp_permission_id)
VALUES(1,1);

INSERT INTO rrol_permission(rrp_role_id, rrp_permission_id)
VALUES(1,2);

INSERT INTO tuser(tus_email, tus_password, tus_name, tus_surname, tus_role_id)
VALUES ('at@dot.com','asd', 'John', 'Lennon',1)