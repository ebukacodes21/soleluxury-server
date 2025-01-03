-- Drop foreign key constraints first
ALTER TABLE "categories" DROP CONSTRAINT IF EXISTS "fk_category_store";
ALTER TABLE "categories" DROP CONSTRAINT IF EXISTS "fk_category_billboard";
ALTER TABLE "billboards" DROP CONSTRAINT IF EXISTS "fk_store";
ALTER TABLE "sizes" DROP CONSTRAINT IF EXISTS "fk_size_store";

-- Drop the tables in reverse order of dependencies
DROP TABLE IF EXISTS "sizes" CASCADE;
DROP TABLE IF EXISTS "categories" CASCADE;
DROP TABLE IF EXISTS "billboards" CASCADE;
DROP TABLE IF EXISTS "stores" CASCADE;
