// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: request.sql

package db

import (
	"context"
	"database/sql"
)

const createRequest = `-- name: CreateRequest :one
INSERT INTO requests (
    user_id, 
    community_id, 
    store_id
) VALUES (
    $1, $2, $3
) RETURNING id, created_at, user_id, community_id, status, errand_id, store_id
`

type CreateRequestParams struct {
	UserID      int64         `json:"user_id"`
	CommunityID sql.NullInt64 `json:"community_id"`
	StoreID     sql.NullInt64 `json:"store_id"`
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) (Request, error) {
	row := q.db.QueryRowContext(ctx, createRequest, arg.UserID, arg.CommunityID, arg.StoreID)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.CommunityID,
		&i.Status,
		&i.ErrandID,
		&i.StoreID,
	)
	return i, err
}

const deleteRequest = `-- name: DeleteRequest :exec
DELETE FROM requests WHERE id = $1
`

func (q *Queries) DeleteRequest(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteRequest, id)
	return err
}

const deleteRequestsByErrand = `-- name: DeleteRequestsByErrand :exec
DELETE FROM requests WHERE errand_id = $1
`

func (q *Queries) DeleteRequestsByErrand(ctx context.Context, errandID sql.NullInt64) error {
	_, err := q.db.ExecContext(ctx, deleteRequestsByErrand, errandID)
	return err
}

const deleteRequestsByStore = `-- name: DeleteRequestsByStore :exec
DELETE FROM requests WHERE store_id = $1
`

func (q *Queries) DeleteRequestsByStore(ctx context.Context, storeID sql.NullInt64) error {
	_, err := q.db.ExecContext(ctx, deleteRequestsByStore, storeID)
	return err
}

const deleteRequestsByUser = `-- name: DeleteRequestsByUser :exec
DELETE FROM requests WHERE user_id = $1
`

func (q *Queries) DeleteRequestsByUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteRequestsByUser, userID)
	return err
}

const getPendingRequestsByCommunityId = `-- name: GetPendingRequestsByCommunityId :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests 
WHERE community_id = $1 AND status = 'pending'
`

func (q *Queries) GetPendingRequestsByCommunityId(ctx context.Context, communityID sql.NullInt64) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getPendingRequestsByCommunityId, communityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const getPendingRequestsByStoreId = `-- name: GetPendingRequestsByStoreId :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests 
WHERE store_id = $1 AND status = 'pending'
`

func (q *Queries) GetPendingRequestsByStoreId(ctx context.Context, storeID sql.NullInt64) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getPendingRequestsByStoreId, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const getRequest = `-- name: GetRequest :one
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests WHERE id = $1
`

func (q *Queries) GetRequest(ctx context.Context, id int64) (Request, error) {
	row := q.db.QueryRowContext(ctx, getRequest, id)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.CommunityID,
		&i.Status,
		&i.ErrandID,
		&i.StoreID,
	)
	return i, err
}

const getRequestsByCommunityId = `-- name: GetRequestsByCommunityId :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests 
WHERE community_id = $1
LIMIT $2
OFFSET $3
`

type GetRequestsByCommunityIdParams struct {
	CommunityID sql.NullInt64 `json:"community_id"`
	Limit       int32         `json:"limit"`
	Offset      int32         `json:"offset"`
}

func (q *Queries) GetRequestsByCommunityId(ctx context.Context, arg GetRequestsByCommunityIdParams) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getRequestsByCommunityId, arg.CommunityID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const getRequestsByErrandId = `-- name: GetRequestsByErrandId :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests WHERE errand_id = $1
`

func (q *Queries) GetRequestsByErrandId(ctx context.Context, errandID sql.NullInt64) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getRequestsByErrandId, errandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const getRequestsByStoreId = `-- name: GetRequestsByStoreId :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests WHERE store_id = $1
`

func (q *Queries) GetRequestsByStoreId(ctx context.Context, storeID sql.NullInt64) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getRequestsByStoreId, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const getRequestsByUserId = `-- name: GetRequestsByUserId :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests 
WHERE user_id = $1
LIMIT $2
OFFSET $3
`

type GetRequestsByUserIdParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetRequestsByUserId(ctx context.Context, arg GetRequestsByUserIdParams) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getRequestsByUserId, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const getRequestsForUserByStatus = `-- name: GetRequestsForUserByStatus :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests
WHERE user_id = $1 and status = $2
`

type GetRequestsForUserByStatusParams struct {
	UserID int64         `json:"user_id"`
	Status RequestStatus `json:"status"`
}

func (q *Queries) GetRequestsForUserByStatus(ctx context.Context, arg GetRequestsForUserByStatusParams) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, getRequestsForUserByStatus, arg.UserID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const listRequests = `-- name: ListRequests :many
SELECT id, created_at, user_id, community_id, status, errand_id, store_id FROM requests
LIMIT $1
OFFSET $2
`

type ListRequestsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRequests(ctx context.Context, arg ListRequestsParams) ([]Request, error) {
	rows, err := q.db.QueryContext(ctx, listRequests, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Request{}
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.CommunityID,
			&i.Status,
			&i.ErrandID,
			&i.StoreID,
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

const updateRequest = `-- name: UpdateRequest :one
UPDATE requests SET 
    user_id = $2, 
    community_id = $3, 
    status = $4, 
    errand_id = $5, 
    store_id = $6
WHERE id = $1
RETURNING id, created_at, user_id, community_id, status, errand_id, store_id
`

type UpdateRequestParams struct {
	ID          int64         `json:"id"`
	UserID      int64         `json:"user_id"`
	CommunityID sql.NullInt64 `json:"community_id"`
	Status      RequestStatus `json:"status"`
	ErrandID    sql.NullInt64 `json:"errand_id"`
	StoreID     sql.NullInt64 `json:"store_id"`
}

func (q *Queries) UpdateRequest(ctx context.Context, arg UpdateRequestParams) (Request, error) {
	row := q.db.QueryRowContext(ctx, updateRequest,
		arg.ID,
		arg.UserID,
		arg.CommunityID,
		arg.Status,
		arg.ErrandID,
		arg.StoreID,
	)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.CommunityID,
		&i.Status,
		&i.ErrandID,
		&i.StoreID,
	)
	return i, err
}

const updateRequestErrandAndStatus = `-- name: UpdateRequestErrandAndStatus :exec
UPDATE requests SET 
    errand_id = $2,
    status = $3
WHERE id = $1
`

type UpdateRequestErrandAndStatusParams struct {
	ID       int64         `json:"id"`
	ErrandID sql.NullInt64 `json:"errand_id"`
	Status   RequestStatus `json:"status"`
}

func (q *Queries) UpdateRequestErrandAndStatus(ctx context.Context, arg UpdateRequestErrandAndStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateRequestErrandAndStatus, arg.ID, arg.ErrandID, arg.Status)
	return err
}

const updateRequestStatus = `-- name: UpdateRequestStatus :one
UPDATE requests SET status = $2 
WHERE id = $1
RETURNING id, created_at, user_id, community_id, status, errand_id, store_id
`

type UpdateRequestStatusParams struct {
	ID     int64         `json:"id"`
	Status RequestStatus `json:"status"`
}

func (q *Queries) UpdateRequestStatus(ctx context.Context, arg UpdateRequestStatusParams) (Request, error) {
	row := q.db.QueryRowContext(ctx, updateRequestStatus, arg.ID, arg.Status)
	var i Request
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.CommunityID,
		&i.Status,
		&i.ErrandID,
		&i.StoreID,
	)
	return i, err
}
