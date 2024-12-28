CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_verified" bool NOT NULL DEFAULT false,
  "verification_code" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL DEFAULT 'user',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);