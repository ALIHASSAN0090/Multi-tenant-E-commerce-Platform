CREATE TABLE IF NOT EXISTS "roles"(
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "name" VARCHAR(255) NOT NULL UNIQUE
);

insert into roles (name) values ('seller'),('admin'),('user')