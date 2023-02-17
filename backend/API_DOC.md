# GoodGrocer API Documentation

## Table of Contents

- [Errors](#errors)
- [User Tokens](#user-tokens)
- [Authentication](#authentication)
  - [Register User](#register-user)
  - [Login User](#login-user)
- [User](#user)
  - [User Commmunities](#user-commmunities)
- [Community](#community)
- [Request](#request)

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
{
  "communities": [
    {
      "id": "[int, community id]",
      "name": "[string, community name]",
      "admin": "[int, community user admin id]",
      "place_id": "[string, community google maps place id]",
      "center_x_coord": "[float, community address x coord]",
      "center_y_coord": "[float, community address x coord]",
      "range": "[int, community range of users]",
      "address": "[string, community address]",
      "created_at": "[timestamptz, community creation date]"
    }
  ]
}
```

---

## Community

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

### Get all Request in and Errand

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
      "store_id": "[int, id of store that request is associated with]", 
    }, 
  ]
  
}
```

**Extra notes** : 


---

## Request
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
  "store_id": "[int, id of store that request is associated with]", 

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
    "id":  "[int, id of the item]", 
    "requested_by":  "[int, id of user that requested item]",
    "request_id":  "[int, id of request]",
    "name":  "[string, name of item]",
    "quantity_type":  "[item_quantity_type, type of quantity (e.g. oz, lbs)]",
    "quantity":  "[float, quantity associated with item type]",
    "preferred_brand":  "[string, brand of item, not required]",
    "image":  "[string, image for item, not required]",
    "found":  "[bool, true if found, else false]",
    "extra_notes":  "[string, notes for shopper]",
  }
}
```

**Extra notes** : json in response returns data for item with 'found' field updated
