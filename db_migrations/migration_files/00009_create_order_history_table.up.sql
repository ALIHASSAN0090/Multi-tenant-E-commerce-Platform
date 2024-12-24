
CREATE TABLE IF NOT EXISTS "order_history"(
    "id" SERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "changed_at" TIMESTAMP WITHOUT TIME ZONE ,
    "changed_by" BIGINT NOT NULL
);