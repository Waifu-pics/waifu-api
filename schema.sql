CREATE TABLE uploads (
    uploaded datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id int AUTO_INCREMENT,
    file varchar(64) NOT NULL,
    md5 varchar(64) NOT NULL,
    TYPE varchar(16) NOT NULL,
    nsfw tinyint (1) DEFAULT 0 NULL,
    verified tinyint (1) DEFAULT 0 NOT NULL,
    CONSTRAINT uploads_file_uindex UNIQUE (file),
    CONSTRAINT uploads_id_uindex UNIQUE (id),
    CONSTRAINT uploads_md5_uindex UNIQUE (md5)
);

CREATE TABLE admins (
    Username varchar(32) NOT NULL,
    PASSWORD varchar(128) NOT NULL,
    CONSTRAINT admins_Username_uindex UNIQUE (Username)
);
