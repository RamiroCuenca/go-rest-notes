CREATE TABLE "notes" (
  "id" BIGSERIAL PRIMARY KEY,
  "owner" VARCHAR NOT NULL,
  "title" VARCHAR NOT NULL,
  "details" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);
