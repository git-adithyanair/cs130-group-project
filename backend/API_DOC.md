# GoodGrocer API Documentation

## Table of Contents

- [Errors](#errors)
- [User Tokens](#user-tokens)
- [Authentication](#authentication)
  - [Register User](#register-user)
  - [Login User](#login-user)
- [User](#user)
  - [Update User Location](#update-user-location)
  - [Update Profile Picture](#update-profile-picture)
  - [Update Name](#update-name)
  - [User Commmunities](#user-commmunities)
  - [User Requests](#user-requests)
- [Community](#community)
  - [Get Community](#get-community)
  - [Get All Communities](#get-all-communities)
  - [Get Community Stores](#get-community-stores)
  - [Create Community](#create-community)
  - [Join Community](#join-community)
  - [Get Community Requests](#get-community-requests)
- [Errand](#errand)
  - [Create Errand](#create-errand)
  - [Update Errand Status](#update-errand-status)
  - [Get all Requests in an Errand](#get-all-request-in-an-errand)
  - [Get Active Errand](#get-active-errand)
- [Request](#request)
  - [Create Request](#create-request)
  - [Get Items in Request](#get-items-in-request)
  - [Change Request Status](#change-request-status)
- [Item](#update-found-status)
  - [Update Found Status](#update-found-status)

## Errors

All errors are returned in the following format:

```json
{
  "id": "ERROR_ID",
  "error": "User friendly error message.",
  "raw": "Raw error message from server."
}
```

See `errors/api_error.go` for a list of all errors.

---

## User Tokens

The user token needs to be added to the request header of all API calls that require authentication. The header key is `Authorization` and should be in the form of `Bearer <token>`.

Examle: `Authorization: Bearer v2.local.6-C_cbq2lrWNlu1vSd9ud2Lt33aSHIbDp7_2_hhXCC3myunfw16IDEGCXGGu-HJI2ef8mQDCl7saj3bCLus8DkrWAIhk9iqa8YabEZAOz0iABhs7DRRIQDsXFnQfMKdXeAkwrisfgaYdzg1EUzuq93tDsvx0H2136S0.bnVsbA`

---

## Authentication

### Register User

Registers user with provided details.

**URL** : `http://api.good-grocer.click/auth/register`

**Method** : `POST`

**Auth Required** : NO

**Body Parameters** :

```json
{
  "email": "[string, valid email]",
  "password": "[string, min. 6 characters]",
  "full_name": "[string]",
  "phone_number": "[string, numeric]",
  "address": "[string]",
  "place_id": "[string]",
  "x_coord": "[float]",
  "y_coord": "[float]"
}
```

**Success Response** : `200 OK`

```json
{
  "token": "[string, user token]",
  "user": {
    "id": "[string, user id]",
    "email": "[string, user email]",
    "full_name": "[string, user full name]",
    "created_at": "[timestamptz, user creation date]"
  }
}
```

### Login User

Logs in user and returns auth token.

**URL** : `http://api.good-grocer.click/auth/login`

**Method** : `POST`

**Auth Required** : NO

**Body Parameters** :

```json
{
  "email": "[string, valid email]",
  "password": "[string, min. 6 characters]"
}
```

**Success Response** : `200 OK`

```json
{
  "token": "[string, user token]",
  "user": {
    "id": "[string, user id]",
    "email": "[string, user email]",
    "full_name": "[string, user full name]",
    "created_at": "[timestamptz, user creation date]"
  }
}
```

---

## User

### Update User Location

Endpoint to update a user's home location.

**URL** : `http://api.good-grocer.click/user/update-location`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "address": "[string, full written address that is displayed to users]",
  "place_id": "[string, place_id from the google response]",
  "x_coord": "[float, x coordinate of location]",
  "y_coord": "[float, y coordinate of location]"
}
```

**Success Response** : `200 OK`

```json
{}
```

**Extra notes** :

### Update Profile Picture

Endpoint to change user profile picture.

**URL** : `http://api.good-grocer.click/update-profile-pic`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "image": "[string, data for the image in string (base64) form]"
}
```

**Success Response** : `200 OK`

```json
{}
```

**Extra notes** : Image data should start something like this: data:image/[IMAGE TYPE];base64

### Update Name

Endpoint to change user's full name.

**URL** : `http://api.good-grocer.click/update-name`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "name": "[string, full name of user]"
}
```

**Success Response** : `200 OK`

```json
{}
```

### User Commmunities

Gets all the communities the user is a member of.

**URL** : `http://api.good-grocer.click/user/community`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

```json
[
  {
    "community": {
      "id": "[int, community id]",
      "name": "[string, community name]",
      "admin": "[int, community user admin id]",
      "place_id": "[string, community google maps place id]",
      "center_x_coord": "[float, community address x coord]",
      "center_y_coord": "[float, community address x coord]",
      "range": "[int, community range of users]",
      "address": "[string, community address]",
      "created_at": "[timestamptz, community creation date]"
    },
    "member_count": "[int, number of members in the community]"
  }
]
```

### User Requests

Endpoint to get users requests grouped by status.

**URL** : `http://api.good-grocer.click/user/requests`

**Method** : `GET`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

```json
{
  "pending": [
    {
      "id": "[int, id of request]",
      "created_at": "[date, time when request was created]",
      "user_id": "[int, id of user who created the request]",
      "community_id": "[int, id of community that request belongs to]",
      "status": "[RequestStatus, the status of the request (pending, in_progress, completed)]",
      "errand_id": "[int, id of errand associated with request, could be null",
      "store_id": "[int, id of store that request is associated with]"
    }
  ],
  "in_progress": [],
  "complete": []
}
```

---

## Community

### Get Community

Endpoint to get info for a specific community.

**URL** : `http://api.good-grocer.click/community/:id`

**Method** : `GET`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

```json
{
  "id": "[int, id of the community]",
  "name": "[string, name of community]",
  "admin": "[int, id of the user who created the community]",
  "place_id": "[string, place_id associated with google map's API location place_id]",
  "center_x_coord": "[float, x coordinate of google maps location for center of community]",
  "center_y_coord": "[float, y coordinate of google maps location for center of community]",
  "range": "[int, range of community (in meters)]",
  "address": "[string, full address of center of community]",
  "created_at": "[date, time when community was created]"
}
```

**Extra Notes**: id in uri should be integer id for the community

### Get All Communities

Endpoint to get all communities.

**URL** : `http://api.good-grocer.click/community`

**Method** : `GET`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

```json
["[json, list of all communities]"]
```

### Get Community Stores

Endpoint to get all stores for a community.

**URL** : `http://api.good-grocer.click/community/stores/:id`

**Method** : `GET`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

```json
[
  {
    "id": "[int, id for store]",
    "name": "[string, name of store]",
    "place_id": "[string, place_id associated with google map's API location place_id]",
    "x_coord": "[float, x coordinate of google maps location]",
    "y_coord": "[float, y coordinate of google maps location]",
    "address": "[string, full address of the store]"
  }
]
```

**Extra Notes**: id in uri should be integer id for the community

### Create Community

Endpoint to create a community.

**URL** : `http://api.good-grocer.click/community`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "name": "[string, name of community]",
  "place_id": "[string, place_id associated with google map's API location place_id]",
  "center_x_coord": "[float, x coordinate of google maps location for center of community]",
  "center_y_coord": "[float, y coordinate of google maps location for center of community]",
  "range": "[int, range of community (in meters)]",
  "address": "[string, full address of center of community]",
  "stores": [
    {
      "name": "[string, name of store]",
      "place_id": "[string, place_id associated with google map's API location place_id]",
      "x_coord": "[float, x coordinate of google maps location]",
      "y_coord": "[float, y coordinate of google maps location]",
      "address": "[string, full address of the store]"
    }
  ]
}
```

**Success Response** : `200 OK`

```json
{
  "id": "[int, id of the community]",
  "name": "[string, name of community]",
  "admin": "[int, id of the user who created the community]",
  "place_id": "[string, place_id associated with google map's API location place_id]",
  "center_x_coord": "[float, x coordinate of google maps location for center of community]",
  "center_y_coord": "[float, y coordinate of google maps location for center of community]",
  "range": "[int, range of community (in meters)]",
  "address": "[string, full address of center of community]",
  "created_at": "[date, time when community was created]"
}
```

### Join Community

Endpoint to join community.

**URL** : `http://api.good-grocer.click/community/join`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "id": "[int, required id of the community]"
}
```

**Success Response** : `200 OK`

```json
{
  "id": "[int, id of the community]",
  "name": "[string, name of community]",
  "admin": "[int, id of the user who created the community]",
  "place_id": "[string, place_id associated with google map's API location place_id]",
  "center_x_coord": "[float, x coordinate of google maps location for center of community]",
  "center_y_coord": "[float, y coordinate of google maps location for center of community]",
  "range": "[int, range of community (in meters)]",
  "address": "[string, full address of center of community]",
  "created_at": "[date, time when community was created]"
}
```

### Get Community Requests

Endpoint to get all of the requests for a community.

**URL** : `http://api.good-grocer.click/community/requests`

**Method** : `GET`

**Auth Required** : YES

**Query Parameters** :

```json
{
  "id": "[int, required id of the community]",
  "limit": "[int, optional limit of requests to return]",
  "offset": "[int, optional offset for paging]"
}
```

**Success Response** : `200 OK`

```json
[
  {
    "request": {
      "id": "[int, id of request]",
      "created_at": "[date, time when request was created]",
      "user_id": "[int, id of user who created the request]",
      "community_id": "[int, id of community that request belongs to]",
      "status": "[RequestStatus, the status of the request (pending, in_progress, completed)]",
      "errand_id": "[int, id of errand associated with request, could be null]",
      "store_id": "[int, id of store that request is associated with]"
    },
    "store": {
      "id": "[int, id for store]",
      "name": "[string, name of store]",
      "place_id": "[string, place_id associated with google map's API location place_id]",
      "x_coord": "[float, x coordinate of google maps location]",
      "y_coord": "[float, y coordinate of google maps location]",
      "address": "[string, full address of the store]"
    },
    "user": {
      "id": "[string, user id]",
      "email": "[string, user email]",
      "full_name": "[string, user full name]",
      "created_at": "[timestamptz, user creation date]"
    }
  },
  ...
]
```

---

## Errand

### Create Errand

Endpoint to create and errand.

**URL** : `http://api.good-grocer.click/errand`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "community_id": "[int, community id]",
  "request_ids": "[int[], array of request ids]"
}
```

**Success Response** : `200 OK`

```json
{
  "id": "[int, id of errand]",
  "user_id": "[int, id of user who is completing errand]",
  "community_id": "[int, id of community that errand belongs to]",
  "is_complete": "[bool, true if errand is complete, false otherwise]",
  "created_at": "[date, when the errand was created]",
  "completed_at": "[date, 0001-01-01T00:00:00Z if errand not complete, otherwise time when errand was completed]"
}
```

**Extra notes** :

### Update Errand Status

Endpoint to update errand status (is_complete field).

**URL** : `http://api.good-grocer.click/errand/update-status`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "id": "[int, required id of the errand]",
  "is_complete": "[bool, true if complete, false is not]"
}
```

**Success Response** : `200 OK`

```json
{
  "id": "[int, id of errand]",
  "user_id": "[int, id of user who is completing errand]",
  "community_id": "[int, id of community that errand belongs to]",
  "is_complete": "[bool, true if errand is complete, false otherwise]",
  "created_at": "[date, when the errand was created]",
  "completed_at": "[date, 0001-01-01T00:00:00Z if errand not complete, otherwise time when errand was completed]"
}
```

### Get all Request in an Errand

Endpoint that returns all data for requests given an errand id.

**URL** : `http://api.good-grocer.click/errand/requests/:id`

**Method** : `GET`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

```json
{
  "requests": [
    {
      "id": "[int, id of request]",
      "created_at": "[date, time when request was created]",
      "user_id": "[int, id of user who created the request]",
      "community_id": "[int, id of community that request belongs to]",
      "status": "[RequestStatus, the status of the request (pending, in_progress, completed)]",
      "errand_id": "[int, id of errand associated with request, could be null]",
      "store_id": "[int, id of store that request is associated with]"
    }
  ]
}
```

**Extra notes** : id in uri is the id of the errand

### Get Active Errand

Endpoint returns a user's active errand and its requests if there is one

**URL** : `http://api.good-grocer.click/errand/active/`

**Method** : `GET`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

If there is an active errand:

```json
{
  "errand": {
    "id": "[int, id of errand]",
    "user_id": "[int, id of user who is completing errand]",
    "community_id": "[int, id of community that errand belongs to]",
    "is_complete": "[bool, true if errand is complete, false otherwise]",
    "created_at": "[date, when the errand was created]",
    "completed_at": "[date, 0001-01-01T00:00:00Z if errand not complete, otherwise time when errand was completed]"
  },
  "requests": [
    {
      "request": {
        "id": "[int, id of request]",
        "created_at": "[date, time when request was created]",
        "user_id": "[int, id of user who created the request]",
        "community_id": "[int, id of community that request belongs to]",
        "status": "[RequestStatus, the status of the request (pending, in_progress, completed)]",
        "errand_id": "[int, id of errand associated with request, could be null]",
        "store_id": "[int, id of store that request is associated with]"
      },
      "items": [
        {
          "id": "[int, id of the item]",
          "requested_by": "[int, id of user that requested item]",
          "request_id": "[int, id of request]",
          "name": "[string, name of item]",
          "quantity_type": "[item_quantity_type, type of quantity (e.g. oz, lbs)]",
          "quantity": "[float, quantity associated with item type]",
          "preferred_brand": "[string, brand of item, not required]",
          "image": "[string, image for item, not required]",
          "found": "[bool, true if found, else false]",
          "extra_notes": "[string, notes for shopper]"
        }
      ],
      "store": {
        "id": "[int, id for store]",
        "name": "[string, name of store]",
        "place_id": "[string, place_id associated with google map's API location place_id]",
        "x_coord": "[float, x coordinate of google maps location]",
        "y_coord": "[float, y coordinate of google maps location]",
        "address": "[string, full address of the store]"
      },
      "user": {
        "id": "[string, user id]",
        "email": "[string, user email]",
        "full_name": "[string, user full name]",
        "phone_number": "[string, user phone number]",
        "x_coord": "[float, user address x coord]",
        "y_coord": "[float, user address y coord]",
        "address": "[string, user address]"
      }
    }
  ]
}
```

If there is no active errand:

```json
{}
```

**Extra Notes**: store could be null

---

## Request

### Create Request

Endpoint to create a request

**URL** : `http://api.good-grocer.click/request`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "community_id": "[int, required id of the community]",
  "store_id": "[int, optional id of the store]",
  "items": [
    {
      "name": "[string, name of item]",
      "quantity_type": "[item_quantity_type, type of quantity (e.g. oz, lbs)]",
      "quantity": "[float, quantity associated with item type]",
      "preferred_brand": "[string, brand of item, not required]",
      "image": "[string, image for item, not required]",
      "extra_notes": "[string, notes for shopper]"
    }
  ]
}
```

**Success Response** : `200 OK`

```json
{
  "id": "[int, id of request]",
  "created_at": "[date, time when request was created]",
  "user_id": "[int, id of user who created the request]",
  "community_id": "[int, id of community that request belongs to]",
  "status": "[RequestStatus, the status of the request (pending, in_progress, completed)]",
  "errand_id": "[int, id of errand associated with request, could be null]",
  "store_id": "[int, id of store that request is associated with]"
}
```

### Get Items in Request

Endpoint to get all items in a request.

**URL** : `http://api.good-grocer.click/request/items/:id`

**Method** : `GET`

**Auth Required** : YES

**Body Parameters** :

```json
{}
```

**Success Response** : `200 OK`

```json
{
  "items": [
    {
      "id": "[int, id of the item]",
      "requested_by": "[int, id of user that requested item]",
      "request_id": "[int, id of request]",
      "name": "[string, name of item]",
      "quantity_type": "[item_quantity_type, type of quantity (e.g. oz, lbs)]",
      "quantity": "[float, quantity associated with item type]",
      "preferred_brand": "[string, brand of item, not required]",
      "image": "[string, image for item, not required]",
      "found": "[bool, true if found, else false]",
      "extra_notes": "[string, notes for shopper]"
    }
  ]
}
```

**Extra notes** : id in uri should be id for request

### Change Request Status

Endpoint to change the status on a request.

**URL** : `http://api.good-grocer.click/request/update-status`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "id": "[int, required id of the request]",
  "status": "[RequestStatus, the status of the request (pending, in_progress, completed)]"
}
```

**Success Response** : `200 OK`

```json
{
  "id": "[int, id of request]",
  "created_at": "[date, time when request was created]",
  "user_id": "[int, id of user who created the request]",
  "community_id": "[int, id of community that request belongs to]",
  "status": "[RequestStatus, the status of the request (pending, in_progress, completed)]",
  "errand_id": "[int, id of errand associated with request, could be null]",
  "store_id": "[int, id of store that request is associated with]"
}
```

**Extra notes** :

---

## Item

### Update Found Status

Updates the status of if an item is found or not.

**URL** : `http://api.good-grocer.click/item/update-status`

**Method** : `POST`

**Auth Required** : YES

**Body Parameters** :

```json
{
  "id": "[int, required id of the item]",
  "found": "[bool, true if found, false is not]"
}
```

**Success Response** : `200 OK`

```json
{
  "item": {
    "id": "[int, id of the item]",
    "requested_by": "[int, id of user that requested item]",
    "request_id": "[int, id of request]",
    "name": "[string, name of item]",
    "quantity_type": "[item_quantity_type, type of quantity (e.g. oz, lbs)]",
    "quantity": "[float, quantity associated with item type]",
    "preferred_brand": "[string, brand of item, not required]",
    "image": "[string, image for item, not required]",
    "found": "[bool, true if found, else false]",
    "extra_notes": "[string, notes for shopper]"
  }
}
```

**Extra notes** : json in response returns data for item with 'found' field updated
