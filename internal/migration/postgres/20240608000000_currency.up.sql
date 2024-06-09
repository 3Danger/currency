
CREATE TABLE currency
(
    code        VARCHAR(10)     NOT NULL,
    rate_to_usd DECIMAL(15, 6)  NOT NULL,
    updated_at  TIMESTAMP       NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_currency_code       ON currency USING HASH  (code);

CREATE INDEX idx_currency_updated_at ON currency USING BTREE (updated_at);
