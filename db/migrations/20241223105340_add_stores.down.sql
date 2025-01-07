-- Drop foreign key constraints first
ALTER TABLE "product_categories" DROP CONSTRAINT IF EXISTS "fk_product_category";
ALTER TABLE "product_sizes" DROP CONSTRAINT IF EXISTS "fk_product_size";
ALTER TABLE "product_colors" DROP CONSTRAINT IF EXISTS "fk_product_color";
ALTER TABLE "categories" DROP CONSTRAINT IF EXISTS "fk_category_store";
ALTER TABLE "categories" DROP CONSTRAINT IF EXISTS "fk_category_billboard";
ALTER TABLE "billboards" DROP CONSTRAINT IF EXISTS "fk_store";
ALTER TABLE "sizes" DROP CONSTRAINT IF EXISTS "fk_size_store";
ALTER TABLE "colors" DROP CONSTRAINT IF EXISTS "fk_color_store";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "fk_order_store"; 

-- Drop tables in reverse order of dependency
DROP TABLE IF EXISTS "product_categories" CASCADE;
DROP TABLE IF EXISTS "product_sizes" CASCADE;
DROP TABLE IF EXISTS "product_colors" CASCADE;

-- Drop the "orders" table
DROP TABLE IF EXISTS "orders" CASCADE;

-- Drop the "products" table last, as it's referenced by other tables
DROP TABLE IF EXISTS "products" CASCADE;

-- Now, drop the remaining tables that do not have dependencies
DROP TABLE IF EXISTS "categories" CASCADE;
DROP TABLE IF EXISTS "billboards" CASCADE;
DROP TABLE IF EXISTS "stores" CASCADE;
DROP TABLE IF EXISTS "sizes" CASCADE;
DROP TABLE IF EXISTS "colors" CASCADE;
