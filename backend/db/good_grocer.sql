CREATE TYPE "item_quantity_type" AS ENUM (
  'numerical',
  'oz',
  'lbs',
  'fl_oz',
  'gal',
  'litres'
);

CREATE TYPE "request_status" AS ENUM (
  'pending',
  'in_progress',
  'completed'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "phone_number" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "address_line_1" varchar NOT NULL,
  "address_line_2" varchar,
  "zip_code" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL
);

CREATE TABLE "communities" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "admin" bigint UNIQUE NOT NULL,
  "center_x_coord" float NOT NULL,
  "center_y_coord" float NOT NULL,
  "range" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "members" (
  "user_id" bigint,
  "community_id" bigint,
  "joined_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "community_id")
);

CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "requested_by" bigint NOT NULL,
  "request_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "quantity_type" item_quantity_type NOT NULL,
  "quantity" float NOT NULL,
  "preferred_brand" varchar,
  "image" varchar,
  "found" bool DEFAULT null,
  "extra_notes" varchar
);

CREATE TABLE "stores" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "address_line_1" varchar NOT NULL,
  "address_line_2" varchar,
  "zip_code" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "x_coord" float NOT NULL,
  "y_coord" varchar NOT NULL,
  "place_id" varchar NOT NULL
);

CREATE TABLE "requests" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" bigint NOT NULL,
  "community_id" bigint,
  "status" request_status NOT NULL DEFAULT 'pending',
  "errand_id" bigint NOT NULL,
  "store_id" bigint
);

CREATE TABLE "errands" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "community_id" bigint NOT NULL,
  "is_complete" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "completed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("phone_number");

CREATE INDEX ON "communities" ("admin");

CREATE INDEX ON "items" ("requested_by");

CREATE INDEX ON "items" ("request_id");

CREATE INDEX ON "stores" ("place_id");

CREATE INDEX ON "requests" ("store_id");

COMMENT ON COLUMN "items"."image" IS 'base64 encoded';

ALTER TABLE "communities" ADD FOREIGN KEY ("admin") REFERENCES "users" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("community_id") REFERENCES "communities" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("requested_by") REFERENCES "users" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("request_id") REFERENCES "requests" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("community_id") REFERENCES "communities" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("errand_id") REFERENCES "errands" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "errands" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "errands" ADD FOREIGN KEY ("community_id") REFERENCES "communities" ("id");
