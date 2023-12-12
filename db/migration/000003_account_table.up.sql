CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "category_id" int NOT NULL,
  "title" varchar NOT NULL,
  "type" varchar NOT NULL,
  "description" varchar NOT NULL,
  "value" integer NOT NULL,
  "date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");