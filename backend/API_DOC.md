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

## Request