
ALTER TABLE "users" ADD CONSTRAINT "users_role_id_foreign" FOREIGN KEY("role_id") REFERENCES "roles"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_created_by_foreign" FOREIGN KEY("created_by") REFERENCES "users"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_store_id_foreign" FOREIGN KEY("store_id") REFERENCES "stores"("id");
ALTER TABLE "payments" ADD CONSTRAINT "payments_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_updated_by_foreign" FOREIGN KEY("updated_by") REFERENCES "users"("id");
ALTER TABLE "order_items" ADD CONSTRAINT "order_items_item_id_foreign" FOREIGN KEY("item_id") REFERENCES "items"("id");
ALTER TABLE "payments" ADD CONSTRAINT "payments_updated_by_foreign" FOREIGN KEY("updated_by") REFERENCES "users"("id");
ALTER TABLE "seller" ADD CONSTRAINT "seller_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE "reviews" ADD CONSTRAINT "reviews_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE "reviews" ADD CONSTRAINT "reviews_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE "payments" ADD CONSTRAINT "payments_created_by_foreign" FOREIGN KEY("created_by") REFERENCES "users"("id");
ALTER TABLE "items" ADD CONSTRAINT "items_store_id_foreign" FOREIGN KEY("store_id") REFERENCES "stores"("id");
ALTER TABLE "stores" ADD CONSTRAINT "stores_client_id_foreign" FOREIGN KEY("client_id") REFERENCES "seller"("id");
ALTER TABLE "order_history" ADD CONSTRAINT "order_history_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");

