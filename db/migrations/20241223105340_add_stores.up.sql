CREATE TABLE "stores" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);