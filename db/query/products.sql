-- name: CreateProduct :one
INSERT INTO products (
    vendor_id,
    name,
    category_id,
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
    $7,
    $8
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

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE($2,name),
    category_id = COALESCE($3,category_id),
    description = COALESCE($4,description),
    price = COALESCE($5,price),
    discount = COALESCE($6,discount),
    stock = COALESCE($7,stock),
    status = COALESCE($8,status),
    updated_at = NOW() 
WHERE
    product_id = $1 RETURNING *;


-- name: DeletProduct :exec
DELETE FROM 
    products 
WHERE 
    product_id = $1;

