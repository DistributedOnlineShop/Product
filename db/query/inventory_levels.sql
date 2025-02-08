-- name: CreateInventoryLevels :one
INSERT INTO inventory_levels (
    product_id,
    pv_id,
    stock
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetInventoryLevelByInventoryId :one
SELECT 
    * 
FROM 
    inventory_levels 
WHERE 
    inventory_id = $1;

-- name: GetInventoryLevelByPvid :many
SELECT 
    * 
FROM 
    inventory_levels 
WHERE 
    product_id = $1 AND pv_id = $2;

-- name: UpdateInventoryLevel :one
UPDATE 
    inventory_levels
SET 
    stock = $2,
    updated_at = NOW()
WHERE 
    inventory_id = $1 RETURNING *;
