-- name: CreateProductImage :one
INSERT INTO product_images (
    product_id,
    pv_id,
    image_url,
    position,
    is_primary
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetProductImageByPiid :many
SELECT 
    * 
FROM 
    product_images 
WHERE 
    pi_id = $1;

-- name: GetProductImageByPvid :many
SELECT 
    * 
FROM 
    product_images 
WHERE 
    product_id = $1 AND pv_id = $2;

-- name: GetProductImageByProductid :many
SELECT 
    * 
FROM 
    product_images 
WHERE 
    product_id = $1;

-- name: GetProductImagePrimary :many
SELECT 
    * 
FROM 
    product_images 
WHERE 
    is_primary = $1;

-- name: UpdateProductImage :one
UPDATE 
    product_images
SET 
    image_url = $2,
    position = $3,
    is_primary = $4,
    updated_at = NOW()
WHERE 
    pi_id = $1 RETURNING *;

-- name: DeleteProductImage :exec
DELETE FROM 
    product_images
WHERE 
    pi_id = $1;