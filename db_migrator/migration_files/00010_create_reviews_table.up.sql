CREATE TABLE IF NOT EXISTS "reviews"(
    "id" SERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "rating" INTEGER NOT NULL,
    "comment" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL
);