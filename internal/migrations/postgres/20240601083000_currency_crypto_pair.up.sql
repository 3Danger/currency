CREATE TYPE CRYPTO_CODE AS ENUM (
    'USDT', 'USDC', 'ETH'
);

CREATE TABLE currency_crypto_pair (
    code_crypto CRYPTO_CODE     NOT NULL,
    code_fiat   FIAT_CODE       NOT NULL,
    updated_at  TIMESTAMPTZ     NOT NULL,
    rate        NUMERIC(14,2)   NOT NULL
);

CREATE UNIQUE INDEX currency_crypto_pair_code_idx_hash ON currency_crypto_pair USING BTREE (code_crypto, code_fiat);
