CREATE TABLE symbol (
    id SERIAL PRIMARY KEY,
    exchange_id integer NULL,
    ticker TEXT NOT NULL,
    instrument TEXT NOT NULL,
    name TEXT NOT NULL,
    sector TEXT NOT NULL,
    currency VARCHAR(64) NULL,
    created_date TIMESTAMP NOT NULL,
    last_updated_date TIMESTAMP NOT NULL,
    FOREIGN KEY (exchange_id) REFERENCES exchange(id)
)