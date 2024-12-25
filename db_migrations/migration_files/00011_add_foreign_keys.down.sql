ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_role_id_foreign";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_created_by_foreign";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_store_id_foreign";
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_order_id_foreign";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_updated_by_foreign";
ALTER TABLE "order_items" DROP CONSTRAINT IF EXISTS "order_items_item_id_foreign";
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_updated_by_foreign";
ALTER TABLE "seller" DROP CONSTRAINT IF EXISTS "seller_user_id_foreign";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_user_id_foreign";
ALTER TABLE "reviews" DROP CONSTRAINT IF EXISTS "reviews_order_id_foreign";
ALTER TABLE "reviews" DROP CONSTRAINT IF EXISTS "reviews_user_id_foreign";
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_created_by_foreign";
ALTER TABLE "items" DROP CONSTRAINT IF EXISTS "items_store_id_foreign";
ALTER TABLE "stores" DROP CONSTRAINT IF EXISTS "stores_seller_id_foreign";
ALTER TABLE "order_history" DROP CONSTRAINT IF EXISTS "order_history_order_id_foreign";

