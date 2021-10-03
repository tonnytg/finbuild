CREATE TABLE "Client" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "created_at" varchar
);

CREATE TABLE "Wallet" (
  "id" SERIAL PRIMARY KEY,
  "user_id" varchar,
  "asset_ids" varchar,
  "created_at" varchar
);

CREATE TABLE "WalletReturn" (
  "id" SERIAL PRIMARY KEY,
  "wallet_id" int,
  "user_id" varchar,
  "status" float,
  "created_at" varchar
);

CREATE TABLE "Assets" (
  "id" SERIAL PRIMARY KEY,
  "wallet_id" int,
  "asset_id" int,
  "price" float,
  "tax" float,
  "quantity" int DEFAULT 1,
  "created_at" varchar
);

ALTER TABLE "Wallet" ADD FOREIGN KEY ("user_id") REFERENCES "Client" ("id");

ALTER TABLE "Assets" ADD FOREIGN KEY ("wallet_id") REFERENCES "Wallet" ("id");

COMMENT ON COLUMN "Client"."created_at" IS 'When order created';

COMMENT ON COLUMN "Wallet"."created_at" IS 'When order created';

COMMENT ON COLUMN "WalletReturn"."created_at" IS 'When order created';

COMMENT ON COLUMN "Assets"."created_at" IS 'When order created';
