-- name: CreateInventoryAdjustments :one
INSERT INTO inventory_adjustments (
    adjustment_id,
    product_id,
    pv_id,
    adjustment_type,
    quantity,
    reason
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetInventoryAdjustmentsByAdjustmentById :one
SELECT 
    * 
FROM 
    inventory_adjustments
WHERE 
    adjustment_id = $1;

-- name: GetInventoryAdjustmentsByPvid :many
SELECT 
    * 
FROM 
    inventory_adjustments
WHERE 
    product_id = $1 AND pv_id = $2;

-- name: GetInventoryAdjustmentsByType :many
SELECT 
    * 
FROM 
    inventory_adjustments
WHERE 
    adjustment_type = $1;

