CREATE TABLE IF NOT EXISTS "stores"(
    "id" SERIAL PRIMARY KEY,
    "store_img" VARCHAR(255),
    "seller_id" BIGINT NOT NULL,
    "store_name" VARCHAR(255) NOT NULL,
    "store_description" VARCHAR(255),
    "store_address" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE ,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);