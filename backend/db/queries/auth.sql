-- name: FindOrCreate :one
INSERT INTO users (email , provider, phone, name, password_hash, created_at)
VALUES ($1, $2, $3, $4, $5, NOW())
ON CONFLICT(email)
DO UPDATE SET
   phone = EXCLUDED.phone,
   provider = EXCLUDED.provider,
   updated_at = now()
RETURNING id , name, phone, provider , role, password_hash;

