// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: store.sql

package db

import (
	"context"
	"database/sql"
)

const createStore = `-- name: CreateStore :one
INSERT INTO stores (
    name, 
    address_line_1, 
    address_line_2, 
    zip_code, 
    city, 
    state, 
    x_coord, 
    y_coord, 
    place_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, name, address_line_1, address_line_2, zip_code, city, state, x_coord, y_coord, place_id
`

type CreateStoreParams struct {
	Name         string         `json:"name"`
	AddressLine1 string         `json:"address_line_1"`
	AddressLine2 sql.NullString `json:"address_line_2"`
	ZipCode      string         `json:"zip_code"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	XCoord       float64        `json:"x_coord"`
	YCoord       string         `json:"y_coord"`
	PlaceID      string         `json:"place_id"`
}

func (q *Queries) CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, createStore,
		arg.Name,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.ZipCode,
		arg.City,
		arg.State,
		arg.XCoord,
		arg.YCoord,
		arg.PlaceID,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.ZipCode,
		&i.City,
		&i.State,
		&i.XCoord,
		&i.YCoord,
		&i.PlaceID,
	)
	return i, err
}

const deleteStore = `-- name: DeleteStore :exec
DELETE FROM stores WHERE id = $1
`

func (q *Queries) DeleteStore(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStore, id)
	return err
}

const getStore = `-- name: GetStore :one
SELECT id, name, address_line_1, address_line_2, zip_code, city, state, x_coord, y_coord, place_id FROM stores WHERE id = $1
`

func (q *Queries) GetStore(ctx context.Context, id int64) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStore, id)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.ZipCode,
		&i.City,
		&i.State,
		&i.XCoord,
		&i.YCoord,
		&i.PlaceID,
	)
	return i, err
}

const getStoreByPlaceId = `-- name: GetStoreByPlaceId :one
SELECT id, name, address_line_1, address_line_2, zip_code, city, state, x_coord, y_coord, place_id FROM stores WHERE place_id = $1
`

func (q *Queries) GetStoreByPlaceId(ctx context.Context, placeID string) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStoreByPlaceId, placeID)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.ZipCode,
		&i.City,
		&i.State,
		&i.XCoord,
		&i.YCoord,
		&i.PlaceID,
	)
	return i, err
}

const listStores = `-- name: ListStores :many
SELECT id, name, address_line_1, address_line_2, zip_code, city, state, x_coord, y_coord, place_id FROM stores
LIMIT $1
OFFSET $2
`

type ListStoresParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStores(ctx context.Context, arg ListStoresParams) ([]Store, error) {
	rows, err := q.db.QueryContext(ctx, listStores, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Store{}
	for rows.Next() {
		var i Store
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AddressLine1,
			&i.AddressLine2,
			&i.ZipCode,
			&i.City,
			&i.State,
			&i.XCoord,
			&i.YCoord,
			&i.PlaceID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStore = `-- name: UpdateStore :one
UPDATE stores SET
    name = $2, 
    address_line_1 = $3, 
    address_line_2 = $4,
    zip_code = $5, 
    city = $6, 
    state = $7, 
    x_coord = $8, 
    y_coord = $9, 
    place_id = $10
WHERE id = $1
RETURNING id, name, address_line_1, address_line_2, zip_code, city, state, x_coord, y_coord, place_id
`

type UpdateStoreParams struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	AddressLine1 string         `json:"address_line_1"`
	AddressLine2 sql.NullString `json:"address_line_2"`
	ZipCode      string         `json:"zip_code"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	XCoord       float64        `json:"x_coord"`
	YCoord       string         `json:"y_coord"`
	PlaceID      string         `json:"place_id"`
}

func (q *Queries) UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, updateStore,
		arg.ID,
		arg.Name,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.ZipCode,
		arg.City,
		arg.State,
		arg.XCoord,
		arg.YCoord,
		arg.PlaceID,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.ZipCode,
		&i.City,
		&i.State,
		&i.XCoord,
		&i.YCoord,
		&i.PlaceID,
	)
	return i, err
}