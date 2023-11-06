-- create "products" table
CREATE TABLE "public"."products" (
  "id" bigserial NOT NULL,
  "product_code" text NULL,
  "product_name" text NULL,
  "price" bigint NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "email" text NULL,
  "password" text NULL,
  "name" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
