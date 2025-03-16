CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "location" varchar NOT NULL,
  "datatimecreating" timestamp NOT NULL DEFAULT (now()),
  "userId" int NOT NULL,
  "datetimeupdating" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "register" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "userid" int NOT NULL,
  "eventid" int NOT NULL
);

CREATE INDEX ON "events" ("userId");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "register" ("userid");

CREATE INDEX ON "register" ("eventid");

ALTER TABLE "events" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id");

ALTER TABLE "register" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "register" ADD FOREIGN KEY ("eventid") REFERENCES "events" ("id");   