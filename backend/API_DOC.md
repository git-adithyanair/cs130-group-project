# GoodGrocer API Documentation

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

## Authentication

### Register User

Registers user with provided details.

**URL** : `http://api.good-grocer.click/auth/register`

**Method** : `POST`

**Auth required** : NO

**Body Parameters**:

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

**Success Response**:`200 OK`

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

**Auth required** : NO

**Body Parameters**:

```json
{
  "email": "[string, valid email]",
  "password": "[string, min. 6 characters]"
}
```

**Success Response**:`200 OK`

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

**Auth required** : YES

**Body Parameters**:

```json
{}
```

**Success Response**:`200 OK`

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
