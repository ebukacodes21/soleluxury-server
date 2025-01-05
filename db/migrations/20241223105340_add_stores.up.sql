-- Creating the "stores" table
CREATE TABLE "stores" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Creating the "billboards" table
CREATE TABLE "billboards" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "label" VARCHAR NOT NULL,
  "image_url" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the "categories" table with foreign keys
CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "billboard_id" bigserial NOT NULL,
  "store_name" VARCHAR NOT NULL,
  "billboard_label" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_category_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_category_billboard" FOREIGN KEY ("billboard_id") REFERENCES "billboards" ("id") ON DELETE CASCADE
);

-- Creating the "sizes" table with foreign keys
CREATE TABLE "sizes" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "store_name" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "value" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_size_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the "colors" table with foreign keys
CREATE TABLE "colors" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "store_name" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "value" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_color_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the "products" table (this must be created before the other tables referencing it)
CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "price" double precision NOT NULL,
  "is_featured" BOOLEAN NOT NULL DEFAULT false,
  "is_archived" BOOLEAN NOT NULL DEFAULT false,
  "description" VARCHAR NOT NULL,
  "images" JSONB NOT NULL, 
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now()
);

-- Creating the "orders" table with store_id
CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL, 
  "items" JSONB NOT NULL, 
  "is_paid" BOOLEAN NOT NULL DEFAULT false,
  "phone" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_order_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the "product_colors" table
CREATE TABLE "product_colors" (
  "product_id" bigserial NOT NULL,
  "color_id" bigserial NOT NULL,
  CONSTRAINT "fk_product_color" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_color_product" FOREIGN KEY ("color_id") REFERENCES "colors" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("product_id", "color_id")
);

-- Creating the "product_sizes" table
CREATE TABLE "product_sizes" (
  "product_id" bigserial NOT NULL,
  "size_id" bigserial NOT NULL,
  CONSTRAINT "fk_product_size" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_size_product" FOREIGN KEY ("size_id") REFERENCES "sizes" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("product_id", "size_id")
);

-- Creating the "product_stores" table
CREATE TABLE "product_stores" (
  "product_id" bigserial NOT NULL,
  "store_id" bigserial NOT NULL,
  CONSTRAINT "fk_product_store" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_store_product" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("product_id", "store_id")
);

-- Creating the "product_categories" table
CREATE TABLE "product_categories" (
  "product_id" bigserial NOT NULL,
  "category_id" bigserial NOT NULL,
  CONSTRAINT "fk_product_category" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_category_product" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("product_id", "category_id")
);
