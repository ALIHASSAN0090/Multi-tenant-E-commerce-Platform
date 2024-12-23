CREATE TABLE IF NOT EXISTS "payments"(
    "id" SERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "payment_method" VARCHAR(255) NOT NULL DEFAULT 'cod',
    "payment_status" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    "created_by" BIGINT NOT NULL,
    "updated_by" BIGINT
);
