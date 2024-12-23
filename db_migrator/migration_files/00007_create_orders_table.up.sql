CREATE TABLE IF NOT EXISTS "orders"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "store_id" BIGINT NOT NULL,
    "total_price" DECIMAL(10, 2) NOT NULL, 
    "status" VARCHAR(255) NOT NULL DEFAULT 'pending',
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    "created_by" BIGINT NOT NULL,
    "updated_by" BIGINT
);