CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_verified" bool NOT NULL DEFAULT false,
  "verification_code" varchar UNIQUE NOT NULL,
  "country" varchar NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL DEFAULT 'user',
  "profile_pic" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);