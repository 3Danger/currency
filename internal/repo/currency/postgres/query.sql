-- name: SetCurrenciesFiat :batchexec
INSERT INTO currency_fiat (code, updated_at, rate_to_usd)
VALUES (@code, @updated_at, @rate_to_usd)
ON CONFLICT (code) DO UPDATE
    SET updated_at  = EXCLUDED.updated_at,
        rate_to_usd = EXCLUDED.rate_to_usd;

-- name: SetCryptoPrices :batchexec
INSERT INTO currency_crypto_pair (code_crypto, code_fiat, updated_at, rate)
VALUES (@code_crypto, @code_fiat, @updated_at, @rate)
ON CONFLICT (code_crypto, code_fiat) DO UPDATE
    SET updated_at = EXCLUDED.updated_at,
        rate = EXCLUDED.rate;

-- name: CurrencyPriceByPair :one
SELECT *
FROM currency_crypto_pair
WHERE code_fiat   = @code_fiat
  AND code_crypto = @code_crypto;

-- name: Currency :one
SELECT * FROM currency_fiat WHERE code = @code;
