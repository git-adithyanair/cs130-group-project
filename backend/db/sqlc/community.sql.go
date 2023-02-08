// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: community.sql

package db

import (
	"context"
)

const createCommunity = `-- name: CreateCommunity :one
INSERT INTO communities (
  name,
  admin,
  center_x_coord,
  center_y_coord,
  range,
  place_id,
  address
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING id, name, admin, place_id, center_x_coord, center_y_coord, range, address, created_at
`

type CreateCommunityParams struct {
	Name         string  `json:"name"`
	Admin        int64   `json:"admin"`
	CenterXCoord float64 `json:"center_x_coord"`
	CenterYCoord float64 `json:"center_y_coord"`
	Range        int32   `json:"range"`
	PlaceID      string  `json:"place_id"`
	Address      string  `json:"address"`
}

func (q *Queries) CreateCommunity(ctx context.Context, arg CreateCommunityParams) (Community, error) {
	row := q.db.QueryRowContext(ctx, createCommunity,
		arg.Name,
		arg.Admin,
		arg.CenterXCoord,
		arg.CenterYCoord,
		arg.Range,
		arg.PlaceID,
		arg.Address,
	)
	var i Community
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Admin,
		&i.PlaceID,
		&i.CenterXCoord,
		&i.CenterYCoord,
		&i.Range,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCommunity = `-- name: DeleteCommunity :exec
DELETE FROM communities WHERE id = $1
`

func (q *Queries) DeleteCommunity(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCommunity, id)
	return err
}

const getCommunitiesByAdmin = `-- name: GetCommunitiesByAdmin :many
SELECT id, name, admin, place_id, center_x_coord, center_y_coord, range, address, created_at FROM communities WHERE admin = $1
`

func (q *Queries) GetCommunitiesByAdmin(ctx context.Context, admin int64) ([]Community, error) {
	rows, err := q.db.QueryContext(ctx, getCommunitiesByAdmin, admin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Community{}
	for rows.Next() {
		var i Community
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Admin,
			&i.PlaceID,
			&i.CenterXCoord,
			&i.CenterYCoord,
			&i.Range,
			&i.Address,
			&i.CreatedAt,
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

const getCommunity = `-- name: GetCommunity :one
SELECT id, name, admin, place_id, center_x_coord, center_y_coord, range, address, created_at FROM communities WHERE id = $1
`

func (q *Queries) GetCommunity(ctx context.Context, id int64) (Community, error) {
	row := q.db.QueryRowContext(ctx, getCommunity, id)
	var i Community
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Admin,
		&i.PlaceID,
		&i.CenterXCoord,
		&i.CenterYCoord,
		&i.Range,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const getCommunityByPlaceID = `-- name: GetCommunityByPlaceID :one
SELECT id, name, admin, place_id, center_x_coord, center_y_coord, range, address, created_at FROM communities WHERE place_id = $1
`

func (q *Queries) GetCommunityByPlaceID(ctx context.Context, placeID string) (Community, error) {
	row := q.db.QueryRowContext(ctx, getCommunityByPlaceID, placeID)
	var i Community
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Admin,
		&i.PlaceID,
		&i.CenterXCoord,
		&i.CenterYCoord,
		&i.Range,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const getUserCommunities = `-- name: GetUserCommunities :many
SELECT communities.id, communities.name, communities.admin, communities.place_id, communities.center_x_coord, communities.center_y_coord, communities.range, communities.address, communities.created_at 
FROM communities
LEFT JOIN members ON members.community_id = communities.id
WHERE members.user_id = $1
`

func (q *Queries) GetUserCommunities(ctx context.Context, userID int64) ([]Community, error) {
	rows, err := q.db.QueryContext(ctx, getUserCommunities, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Community{}
	for rows.Next() {
		var i Community
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Admin,
			&i.PlaceID,
			&i.CenterXCoord,
			&i.CenterYCoord,
			&i.Range,
			&i.Address,
			&i.CreatedAt,
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

const listCommunities = `-- name: ListCommunities :many
SELECT id, name, admin, place_id, center_x_coord, center_y_coord, range, address, created_at FROM communities
LIMIT $1
OFFSET $2
`

type ListCommunitiesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCommunities(ctx context.Context, arg ListCommunitiesParams) ([]Community, error) {
	rows, err := q.db.QueryContext(ctx, listCommunities, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Community{}
	for rows.Next() {
		var i Community
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Admin,
			&i.PlaceID,
			&i.CenterXCoord,
			&i.CenterYCoord,
			&i.Range,
			&i.Address,
			&i.CreatedAt,
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

const updateCommunity = `-- name: UpdateCommunity :one
UPDATE communities SET
  name = $2,
  admin = $3,
  center_x_coord = $4,
  center_y_coord = $5,
  range = $6,
  place_id = $7,
  address = $8
WHERE id = $1
RETURNING id, name, admin, place_id, center_x_coord, center_y_coord, range, address, created_at
`

type UpdateCommunityParams struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Admin        int64   `json:"admin"`
	CenterXCoord float64 `json:"center_x_coord"`
	CenterYCoord float64 `json:"center_y_coord"`
	Range        int32   `json:"range"`
	PlaceID      string  `json:"place_id"`
	Address      string  `json:"address"`
}

func (q *Queries) UpdateCommunity(ctx context.Context, arg UpdateCommunityParams) (Community, error) {
	row := q.db.QueryRowContext(ctx, updateCommunity,
		arg.ID,
		arg.Name,
		arg.Admin,
		arg.CenterXCoord,
		arg.CenterYCoord,
		arg.Range,
		arg.PlaceID,
		arg.Address,
	)
	var i Community
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Admin,
		&i.PlaceID,
		&i.CenterXCoord,
		&i.CenterYCoord,
		&i.Range,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const updateCommunityAdmin = `-- name: UpdateCommunityAdmin :one
UPDATE communities SET
  admin = $2
WHERE id = $1
RETURNING id, name, admin, place_id, center_x_coord, center_y_coord, range, address, created_at
`

type UpdateCommunityAdminParams struct {
	ID    int64 `json:"id"`
	Admin int64 `json:"admin"`
}

func (q *Queries) UpdateCommunityAdmin(ctx context.Context, arg UpdateCommunityAdminParams) (Community, error) {
	row := q.db.QueryRowContext(ctx, updateCommunityAdmin, arg.ID, arg.Admin)
	var i Community
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Admin,
		&i.PlaceID,
		&i.CenterXCoord,
		&i.CenterYCoord,
		&i.Range,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}
