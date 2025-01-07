-- Creating the "stores" table
CREATE TABLE "stores" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Creating the indexes for "stores"
CREATE INDEX "idx_stores_name" ON "stores" ("name");
CREATE INDEX "idx_stores_created_at" ON "stores" ("created_at");

-- Creating the "billboards" table
CREATE TABLE "billboards" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "label" VARCHAR NOT NULL,
  "image_url" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the indexes for "billboards"
CREATE INDEX "idx_billboards_store_id" ON "billboards" ("store_id");

-- Creating the "categories" table with foreign keys
CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "billboard_id" bigserial NOT NULL,
  "name" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_category_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_category_billboard" FOREIGN KEY ("billboard_id") REFERENCES "billboards" ("id") ON DELETE CASCADE
);

-- Creating the indexes for "categories"
CREATE INDEX "idx_categories_store_id" ON "categories" ("store_id");
CREATE INDEX "idx_categories_billboard_id" ON "categories" ("billboard_id");

-- Creating the "sizes" table with foreign keys
CREATE TABLE "sizes" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "name" VARCHAR NOT NULL,
  "value" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_size_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the indexes for "sizes"
CREATE INDEX "idx_sizes_store_id" ON "sizes" ("store_id");

-- Creating the "colors" table with foreign keys
CREATE TABLE "colors" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "name" VARCHAR NOT NULL,
  "value" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_color_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the indexes for "colors"
CREATE INDEX "idx_colors_store_id" ON "colors" ("store_id");

-- Creating the "products" table
CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "name" VARCHAR NOT NULL,
  "price" double precision NOT NULL,
  "is_featured" BOOLEAN NOT NULL DEFAULT false,
  "is_archived" BOOLEAN NOT NULL DEFAULT false,
  "description" VARCHAR NOT NULL,
  "images" JSONB NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "fk_product_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE
);

-- Creating the indexes for "products"
CREATE INDEX "idx_products_store_id" ON "products" ("store_id");
CREATE INDEX "idx_products_is_featured" ON "products" ("is_featured");
CREATE INDEX "idx_products_is_archived" ON "products" ("is_archived");

-- Creating the "orders" table
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

-- Creating the indexes for "orders"
CREATE INDEX "idx_orders_store_id" ON "orders" ("store_id");
CREATE INDEX "idx_orders_is_paid" ON "orders" ("is_paid");

-- Creating the "product_colors" table
CREATE TABLE "product_colors" (
  "product_id" bigserial NOT NULL,
  "color_id" bigserial NOT NULL,
  CONSTRAINT "fk_product_color" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_color_product" FOREIGN KEY ("color_id") REFERENCES "colors" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("product_id", "color_id")
);

-- Creating the indexes for "product_colors"
CREATE INDEX "idx_product_colors_product_id" ON "product_colors" ("product_id");
CREATE INDEX "idx_product_colors_color_id" ON "product_colors" ("color_id");

-- Creating the "product_sizes" table
CREATE TABLE "product_sizes" (
  "product_id" bigserial NOT NULL,
  "size_id" bigserial NOT NULL,
  CONSTRAINT "fk_product_size" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_size_product" FOREIGN KEY ("size_id") REFERENCES "sizes" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("product_id", "size_id")
);

-- Creating the indexes for "product_sizes"
CREATE INDEX "idx_product_sizes_product_id" ON "product_sizes" ("product_id");
CREATE INDEX "idx_product_sizes_size_id" ON "product_sizes" ("size_id");

-- Creating the "product_categories" table
CREATE TABLE "product_categories" (
  "product_id" bigserial NOT NULL,
  "category_id" bigserial NOT NULL,
  CONSTRAINT "fk_product_category" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_category_product" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE,
  PRIMARY KEY ("product_id", "category_id")
);

-- Creating the indexes for "product_categories"
CREATE INDEX "idx_product_categories_product_id" ON "product_categories" ("product_id");
CREATE INDEX "idx_product_categories_category_id" ON "product_categories" ("category_id");
