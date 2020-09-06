create table uploads
(
    id       int auto_increment,
    file     varchar(64)          not null,
    md5      varchar(64)          not null,
    type     varchar(16)          not null,
    nsfw     tinyint(1) default 0 not null,
    verified tinyint(1) default 0 not null,
    constraint uploads_file_uindex
        unique (file),
    constraint uploads_id_uindex
        unique (id),
    constraint uploads_md5_uindex
        unique (md5)
);

create table admins
(
    Username varchar(32)  not null,
    Password varchar(128) not null,
    constraint admins_Username_uindex
        unique (Username)
);