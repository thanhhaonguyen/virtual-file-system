CREATE SCHEMA file_system

CREATE TABLE file_system.folder (
    id bigserial NOT NULL,
    "name" varchar(255) NULL,
    parent_id int8 NULL,
    created_at int8 NOT NULL,
    CONSTRAINT folder_pkey PRIMARY KEY (id)
);

CREATE TABLE file_system.folder (
    id bigserial NOT NULL,
    "name" varchar(255) NULL,
    parent_id int8 NULL,
    created_at int8 NOT NULL,
    CONSTRAINT folder_pkey PRIMARY KEY (id)
);