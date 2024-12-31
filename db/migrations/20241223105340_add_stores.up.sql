CREATE TABLE "stores" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "billboards" (
  "id" bigserial PRIMARY KEY,
  "store_id" bigserial NOT NULL,
  "label" VARCHAR NOT NULL,
  "image_url" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "billboards"
ADD CONSTRAINT "fk_store" FOREIGN KEY ("store_id") REFERENCES "stores" ("id") ON DELETE CASCADE;