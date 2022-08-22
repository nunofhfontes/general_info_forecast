DROP TABLE "account";

CREATE TABLE "account" (
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    name TEXT NOT NULL,
    role text NOT NULL,
    password text NOT NULL
);