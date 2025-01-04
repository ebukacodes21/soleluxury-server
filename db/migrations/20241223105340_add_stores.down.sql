-- Drop foreign key constraints first
ALTER TABLE "product_categories" DROP CONSTRAINT IF EXISTS "fk_product_category";
ALTER TABLE "product_stores" DROP CONSTRAINT IF EXISTS "fk_product_store";
ALTER TABLE "product_sizes" DROP CONSTRAINT IF EXISTS "fk_product_size";
ALTER TABLE "product_colors" DROP CONSTRAINT IF EXISTS "fk_product_color";
ALTER TABLE "images" DROP CONSTRAINT IF EXISTS "fk_image_product";  -- Foreign key on "images"
ALTER TABLE "categories" DROP CONSTRAINT IF EXISTS "fk_category_store";
ALTER TABLE "categories" DROP CONSTRAINT IF EXISTS "fk_category_billboard";
ALTER TABLE "billboards" DROP CONSTRAINT IF EXISTS "fk_store";
ALTER TABLE "sizes" DROP CONSTRAINT IF EXISTS "fk_size_store";
ALTER TABLE "colors" DROP CONSTRAINT IF EXISTS "fk_color_store";

-- Drop the tables in reverse order of dependencies
DROP TABLE IF EXISTS "product_categories" CASCADE;
DROP TABLE IF EXISTS "product_stores" CASCADE;
DROP TABLE IF EXISTS "product_sizes" CASCADE;
DROP TABLE IF EXISTS "product_colors" CASCADE;
DROP TABLE IF EXISTS "images" CASCADE;  -- Drops the images table and all associated image URLs

-- Now, drop the tables that do not depend on others
DROP TABLE IF EXISTS "products" CASCADE;

-- Drop the remaining tables in reverse dependency order
DROP TABLE IF EXISTS "categories" CASCADE;
DROP TABLE IF EXISTS "billboards" CASCADE;
DROP TABLE IF EXISTS "stores" CASCADE;
DROP TABLE IF EXISTS "sizes" CASCADE;
DROP TABLE IF EXISTS "colors" CASCADE;
