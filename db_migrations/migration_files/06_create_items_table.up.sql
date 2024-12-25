CREATE TABLE IF NOT EXISTS "items"(
    "id" SERIAL PRIMARY KEY,
    "item_img" VARCHAR(255),
    "store_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "price" DECIMAL(8, 2) NOT NULL,
    "stock_quantity" INTEGER NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE ,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    "discount" BIGINT NOT NULL DEFAULT 0
);