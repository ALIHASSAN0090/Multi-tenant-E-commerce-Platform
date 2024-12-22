CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY,
    "role_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "hash_password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE "stores"(
    "id" BIGINT PRIMARY KEY,
    "client_id" BIGINT NOT NULL,
    "store_name" VARCHAR(255) NOT NULL,
    "store_description" VARCHAR(255),
    "store_address" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE "order_items"(
    "id" SERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "item_id" BIGINT NOT NULL,
    "quantity" INTEGER NOT NULL, 
    "price_per_item" DECIMAL(10, 2) NOT NULL,  
    "total_price" DECIMAL(10, 2) NOT NULL, 
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE "roles"(
    "id" BIGINT PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE "clients"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "business_name" VARCHAR(255) NOT NULL,
    "contact_number" VARCHAR(20), 
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);


CREATE TABLE "items"(
    "id" BIGINT PRIMARY KEY,
    "store_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "price" DECIMAL(8, 2) NOT NULL,
    "stock_quantity" INTEGER NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    "discount" BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE "orders"(
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

CREATE TABLE "payments"(
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

CREATE TABLE "order_history"(
    "id" SERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "changed_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "changed_by" BIGINT NOT NULL
);

CREATE TABLE "reviews"(
    "id" SERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "rating" INTEGER NOT NULL,
    "comment" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

ALTER TABLE "users" ADD CONSTRAINT "users_role_id_foreign" FOREIGN KEY("role_id") REFERENCES "roles"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_created_by_foreign" FOREIGN KEY("created_by") REFERENCES "users"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_store_id_foreign" FOREIGN KEY("store_id") REFERENCES "stores"("id");
ALTER TABLE "payments" ADD CONSTRAINT "payments_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_updated_by_foreign" FOREIGN KEY("updated_by") REFERENCES "users"("id");
ALTER TABLE "order_items" ADD CONSTRAINT "order_items_item_id_foreign" FOREIGN KEY("item_id") REFERENCES "items"("id");
ALTER TABLE "payments" ADD CONSTRAINT "payments_updated_by_foreign" FOREIGN KEY("updated_by") REFERENCES "users"("id");
ALTER TABLE "clients" ADD CONSTRAINT "clients_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE "reviews" ADD CONSTRAINT "reviews_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE "reviews" ADD CONSTRAINT "reviews_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE "payments" ADD CONSTRAINT "payments_created_by_foreign" FOREIGN KEY("created_by") REFERENCES "users"("id");
ALTER TABLE "items" ADD CONSTRAINT "items_store_id_foreign" FOREIGN KEY("store_id") REFERENCES "stores"("id");
ALTER TABLE "stores" ADD CONSTRAINT "stores_client_id_foreign" FOREIGN KEY("client_id") REFERENCES "clients"("id");
ALTER TABLE "order_history" ADD CONSTRAINT "order_history_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
