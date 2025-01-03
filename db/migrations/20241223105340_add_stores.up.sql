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
