CREATE TABLE IF NOT EXISTS "clients"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "business_name" VARCHAR(255) NOT NULL,
    "contact_number" VARCHAR(20), 
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);