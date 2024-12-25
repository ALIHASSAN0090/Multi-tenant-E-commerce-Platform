CREATE TABLE IF NOT EXISTS "seller"(
    "id" SERIAL PRIMARY KEY,
    "seller_img" VARCHAR(255),
    "cnic_number" VARCHAR(15)NOT NULL,
    "cnic_image" VARCHAR(255)NOT NULL,
    "user_id" BIGINT NOT NULL,
    "business_name" VARCHAR(255) NOT NULL,
    "contact_number" VARCHAR(20), 
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE ,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);