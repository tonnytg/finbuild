CREATE TABLE Users (
                       "id" SERIAL PRIMARY KEY,
                       "user_id" varchar,
                       "first_name" varchar,
                       "last_name" varchar,
                       "social_id" varchar,
                       "phone" varchar,
                       "email" varchar,
                       "valid_email" bool,
                       "city" varchar,
                       "country" varchar,
                       "address" varchar,
                       "zipcode" varchar,
                       "age" int,
                       "sex" varchar,
                       "sign" varchar,
                       "created_at" TIMESTAMPTZ DEFAULT Now()
);

CREATE TABLE Wallets (
                         "id" SERIAL PRIMARY KEY,
                         "wallet_id" varchar,
                         "user_id" varchar,
                         "balance" float,
                         "created_at" TIMESTAMPTZ DEFAULT Now()
);

CREATE TABLE WalletReturn (
                              "id" SERIAL PRIMARY KEY,
                              "wallet_id" int,
                              "status" float,
                              "created_at" TIMESTAMPTZ DEFAULT Now()
);

CREATE TABLE Exchanges (
                           "id" SERIAL PRIMARY KEY,
                           "wallet_id" varchar,
                           "asset_id" varchar,
                           "price" float,
                           "tax" float,
                           "action" varchar,
                           "quantity" int DEFAULT 1,
                           "date" varchar,
                           "created_at" TIMESTAMPTZ DEFAULT Now()
);

CREATE TABLE Assets (
                        "id" SERIAL PRIMARY KEY,
                        "wallet_id" int,
                        "asset_id" int,
                        "medium_price" float,
                        "full_price" float,
                        "quantity" int DEFAULT 1,
                        "created_at" TIMESTAMPTZ DEFAULT Now()
);

ALTER TABLE Wallets ADD FOREIGN KEY ("user_id") REFERENCES Users ("id");

ALTER TABLE Assets ADD FOREIGN KEY ("wallet_id") REFERENCES Wallets ("id");
