CREATE TYPE FIAT_CODE AS ENUM (
    'EUR', 'USD', 'CNY'
);

CREATE TABLE currency_fiat (
    code        FIAT_CODE       NOT NULL,
    updated_at  TIMESTAMPTZ     NOT NULL,
    rate_to_usd NUMERIC(14,2)   NOT NULL
);

CREATE UNIQUE INDEX currency_fiat_code_idx_hash ON currency_fiat USING BTREE (code);
