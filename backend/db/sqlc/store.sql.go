// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: store.sql

package db

import (
	"context"
)

const createStore = `-- name: CreateStore :one
INSERT INTO stores (
    name, 
    x_coord, 
    y_coord, 
    place_id,
    address
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, name, x_coord, y_coord, address, place_id
`

type CreateStoreParams struct {
	Name    string  `json:"name"`
	XCoord  float64 `json:"x_coord"`
	YCoord  float64 `json:"y_coord"`
	PlaceID string  `json:"place_id"`
	Address string  `json:"address"`
}

func (q *Queries) CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, createStore,
		arg.Name,
		arg.XCoord,
		arg.YCoord,
		arg.PlaceID,
		arg.Address,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.XCoord,
		&i.YCoord,
		&i.Address,
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
SELECT id, name, x_coord, y_coord, address, place_id FROM stores WHERE id = $1
`

func (q *Queries) GetStore(ctx context.Context, id int64) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStore, id)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.XCoord,
		&i.YCoord,
		&i.Address,
		&i.PlaceID,
	)
	return i, err
}

const getStoreByPlaceId = `-- name: GetStoreByPlaceId :one
SELECT id, name, x_coord, y_coord, address, place_id FROM stores WHERE place_id = $1
`

func (q *Queries) GetStoreByPlaceId(ctx context.Context, placeID string) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStoreByPlaceId, placeID)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.XCoord,
		&i.YCoord,
		&i.Address,
		&i.PlaceID,
	)
	return i, err
}

const getStoresByCommunity = `-- name: GetStoresByCommunity :many
SELECT stores.id, stores.name, stores.x_coord, stores.y_coord, stores.address, stores.place_id 
FROM stores 
LEFT JOIN community_stores ON community_stores.store_id = stores.id
WHERE community_stores.community_id = $1
`

func (q *Queries) GetStoresByCommunity(ctx context.Context, communityID int64) ([]Store, error) {
	rows, err := q.db.QueryContext(ctx, getStoresByCommunity, communityID)
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
			&i.XCoord,
			&i.YCoord,
			&i.Address,
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

const listStores = `-- name: ListStores :many
SELECT id, name, x_coord, y_coord, address, place_id FROM stores
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
			&i.XCoord,
			&i.YCoord,
			&i.Address,
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
    x_coord = $3, 
    y_coord = $4, 
    place_id = $5,
    address = $6
WHERE id = $1
RETURNING id, name, x_coord, y_coord, address, place_id
`

type UpdateStoreParams struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	XCoord  float64 `json:"x_coord"`
	YCoord  float64 `json:"y_coord"`
	PlaceID string  `json:"place_id"`
	Address string  `json:"address"`
}

func (q *Queries) UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, updateStore,
		arg.ID,
		arg.Name,
		arg.XCoord,
		arg.YCoord,
		arg.PlaceID,
		arg.Address,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.XCoord,
		&i.YCoord,
		&i.Address,
		&i.PlaceID,
	)
	return i, err
}
