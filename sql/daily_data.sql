CREATE TABLE daily_data (
    id SERIAL PRIMARY KEY,
    data_vendor_id INTEGER NOT NULL,
    stock_id INTEGER NOT NULL,
    created_date TIMESTAMP NOT NULL,
    last_updated_date TIMESTAMP NOT NULL,
    date_price DATE,
    open_price NUMERIC,
    high_price NUMERIC,
    low_price NUMERIC,
    close_price NUMERIC,
    adj_close_price NUMERIC,
    volume BIGINT,
    FOREIGN KEY (data_vendor_id) REFERENCES data_vendor(id),
    FOREIGN KEY (stock_id) REFERENCES symbol(id)
)