-- name: CreateProduct :one
INSERT INTO products (
    vendor_id,
    name,
    description,
    price,
    discount,
    stock,
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

-- name: GetProductByProductId :one
SELECT 
    * 
FROM 
    products 
WHERE 
    product_id = $1;

-- name: GetProductByVendorID :many
SELECT
    *
FROM 
    products 
WHERE 
    vendor_id = $1;

-- name: UpdateProductDetail :one
UPDATE products
SET
    name = COALESCE($2,name),
    description = COALESCE($3,description),
    price = COALESCE($4,price),
    discount = COALESCE($5,discount),
    stock = COALESCE($6,stock),
    status = COALESCE($7,status),
    updated_at = NOW() 
WHERE
    product_id = $1 RETURNING *;


-- name: DeletProduct :exec
DELETE FROM 
    products 
WHERE 
    product_id = $1;

