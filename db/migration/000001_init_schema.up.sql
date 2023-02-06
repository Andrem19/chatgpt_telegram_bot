CREATE TABLE "gpt_user" (
  "id" BIGSERIAL PRIMARY KEY,
  "chat_id" varchar (100) NOT NULL,
  "gpt_token" varchar (100) NOT NULL UNIQUE,
  "step" smallint NOT NULL,
  "last_answer" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);