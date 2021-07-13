CREATE TABLE "users"
(
    "id"         varchar(255) UNIQUE NOT NULL PRIMARY KEY,
    "email"      varchar(255) UNIQUE NOT NULL,
    "phone"      varchar(10) UNIQUE  NOT NULL,
    "password"   varchar(255)        NOT NULL,
    "address"    varchar(255),
    "fullName"   varchar(120),
    "avatar"     text,
    "role"       varchar(25),
    "created_at" timestamp           NOT NULL,
    "updated_at" timestamp
);