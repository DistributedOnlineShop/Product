-- name: CreateProductVariants :one
INSERT INTO product_variants (
    product_id,
    sku,
    price,
    stock,
    attributes,
    status
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetProductVariantsByPvid :one
SELECT 
    * 
FROM 
    product_variants 
WHERE 
    pv_id = $1;

-- name: GetProductVariantsByProductId :many
SELECT 
    * 
FROM 
    product_variants 
WHERE 
    product_id = $1;

-- name: GetProductVariantsByStatus :many
SELECT 
    * 
FROM 
    product_variants 
WHERE 
    status = $1;

-- name: UpdateProductVariant :one
UPDATE 
    product_variants 
SET
    sku = COALESCE($2, sku),
    price = COALESCE($3, price),
    stock = COALESCE($4, stock),
    attributes = COALESCE($5, attributes),
    status = COALESCE($6, status),
    updated_at = NOW()
WHERE 
    pv_id = $1 RETURNING *;

-- name: DeleteProductVariantsByPvid :exec
DELETE FROM 
    product_variants 
WHERE 
    pv_id = $1;