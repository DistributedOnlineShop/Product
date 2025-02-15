-- name: CreateVendors :one
INSERT INTO vendors (
    vendor_id,
    vendor_name,
    contact_name,
    product_type,
    email,
    phone,
    status
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING *;


-- name: GetVendorsList :many
SELECT
    *
FROM
    vendors;


-- name: GetVendorsByStatus :many
SELECT
    *
FROM
    vendors
WHERE
    status = $1;

-- name: UpdateVendorStatus :one
UPDATE
    vendors
SET
    status = $2,
    last_updated_at = NOW()
WHERE
    vendor_id = $1 RETURNING *;

-- name: UpdateVendor :one
UPDATE
    vendors
SET
    vendor_name = COALESCE($2,vendor_name),
    contact_name = COALESCE($3,contact_name),
    email = COALESCE($4,email),
    phone = COALESCE($5,phone),
    status = COALESCE($6,status),
    updated_at = NOW()
WHERE
    vendor_id = $1 RETURNING *;