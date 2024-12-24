CREATE TABLE IF NOT EXISTS "order_items"(
    "id" SERIAL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "item_id" BIGINT NOT NULL,
    "quantity" INTEGER NOT NULL, 
    "price_per_item" DECIMAL(10, 2) NOT NULL,  
    "total_price" DECIMAL(10, 2) NOT NULL, 
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE ,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE
);