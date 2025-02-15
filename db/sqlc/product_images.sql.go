// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: product_images.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProductImage = `-- name: CreateProductImage :one
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
) RETURNING pi_id, product_id, pv_id, image_url, position, is_primary, created_at, updated_at
`

type CreateProductImageParams struct {
	ProductID string      `json:"product_id"`
	PvID      string      `json:"pv_id"`
	ImageUrl  string      `json:"image_url"`
	Position  int32       `json:"position"`
	IsPrimary pgtype.Bool `json:"is_primary"`
}

func (q *Queries) CreateProductImage(ctx context.Context, arg CreateProductImageParams) (ProductImage, error) {
	row := q.db.QueryRow(ctx, createProductImage,
		arg.ProductID,
		arg.PvID,
		arg.ImageUrl,
		arg.Position,
		arg.IsPrimary,
	)
	var i ProductImage
	err := row.Scan(
		&i.PiID,
		&i.ProductID,
		&i.PvID,
		&i.ImageUrl,
		&i.Position,
		&i.IsPrimary,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProductImage = `-- name: DeleteProductImage :exec
DELETE FROM 
    product_images
WHERE 
    pi_id = $1
`

func (q *Queries) DeleteProductImage(ctx context.Context, piID string) error {
	_, err := q.db.Exec(ctx, deleteProductImage, piID)
	return err
}

const getProductImageByPiid = `-- name: GetProductImageByPiid :many
SELECT 
    pi_id, product_id, pv_id, image_url, position, is_primary, created_at, updated_at 
FROM 
    product_images 
WHERE 
    pi_id = $1
`

func (q *Queries) GetProductImageByPiid(ctx context.Context, piID string) ([]ProductImage, error) {
	rows, err := q.db.Query(ctx, getProductImageByPiid, piID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductImage{}
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.PiID,
			&i.ProductID,
			&i.PvID,
			&i.ImageUrl,
			&i.Position,
			&i.IsPrimary,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductImageByProductid = `-- name: GetProductImageByProductid :many
SELECT 
    pi_id, product_id, pv_id, image_url, position, is_primary, created_at, updated_at 
FROM 
    product_images 
WHERE 
    product_id = $1
`

func (q *Queries) GetProductImageByProductid(ctx context.Context, productID string) ([]ProductImage, error) {
	rows, err := q.db.Query(ctx, getProductImageByProductid, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductImage{}
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.PiID,
			&i.ProductID,
			&i.PvID,
			&i.ImageUrl,
			&i.Position,
			&i.IsPrimary,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductImageByPvid = `-- name: GetProductImageByPvid :many
SELECT 
    pi_id, product_id, pv_id, image_url, position, is_primary, created_at, updated_at 
FROM 
    product_images 
WHERE 
    product_id = $1 AND pv_id = $2
`

type GetProductImageByPvidParams struct {
	ProductID string `json:"product_id"`
	PvID      string `json:"pv_id"`
}

func (q *Queries) GetProductImageByPvid(ctx context.Context, arg GetProductImageByPvidParams) ([]ProductImage, error) {
	rows, err := q.db.Query(ctx, getProductImageByPvid, arg.ProductID, arg.PvID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductImage{}
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.PiID,
			&i.ProductID,
			&i.PvID,
			&i.ImageUrl,
			&i.Position,
			&i.IsPrimary,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductImagePrimary = `-- name: GetProductImagePrimary :many
SELECT 
    pi_id, product_id, pv_id, image_url, position, is_primary, created_at, updated_at 
FROM 
    product_images 
WHERE 
    pi_id = $1 AND is_primary = $2
`

type GetProductImagePrimaryParams struct {
	PiID      string      `json:"pi_id"`
	IsPrimary pgtype.Bool `json:"is_primary"`
}

func (q *Queries) GetProductImagePrimary(ctx context.Context, arg GetProductImagePrimaryParams) ([]ProductImage, error) {
	rows, err := q.db.Query(ctx, getProductImagePrimary, arg.PiID, arg.IsPrimary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductImage{}
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.PiID,
			&i.ProductID,
			&i.PvID,
			&i.ImageUrl,
			&i.Position,
			&i.IsPrimary,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProductImage = `-- name: UpdateProductImage :one
UPDATE 
    product_images
SET 
    image_url = $2,
    position = $3,
    is_primary = $4,
    updated_at = NOW()
WHERE 
    pi_id = $1 RETURNING pi_id, product_id, pv_id, image_url, position, is_primary, created_at, updated_at
`

type UpdateProductImageParams struct {
	PiID      string      `json:"pi_id"`
	ImageUrl  string      `json:"image_url"`
	Position  int32       `json:"position"`
	IsPrimary pgtype.Bool `json:"is_primary"`
}

func (q *Queries) UpdateProductImage(ctx context.Context, arg UpdateProductImageParams) (ProductImage, error) {
	row := q.db.QueryRow(ctx, updateProductImage,
		arg.PiID,
		arg.ImageUrl,
		arg.Position,
		arg.IsPrimary,
	)
	var i ProductImage
	err := row.Scan(
		&i.PiID,
		&i.ProductID,
		&i.PvID,
		&i.ImageUrl,
		&i.Position,
		&i.IsPrimary,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
