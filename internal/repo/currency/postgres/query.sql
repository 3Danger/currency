-- name: Currency :one
SELECT *
FROM currency
WHERE code = @code;


-- name: Upsert :exec
INSERT INTO currency (code, rate_to_usd, updated_at)
VALUES (@Code, @RateToUSD, NOW())
ON CONFLICT (code) DO UPDATE
    SET rate_to_usd = EXCLUDED.rate_to_usd,
        updated_at = EXCLUDED.updated_at;
