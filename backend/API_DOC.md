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
  "email": "[valid email]", // string
  "password": "[string of 6 characters]", // string
  "full_name": "[]", // string
  "phone_number": "[numeric]", // string
  "address": "[]", // string
  "place_id": "[]", // string
  "x_coord": "[]", // float
  "y_coord": "[]" // float
}
```

**Success Response**:`200 OK`

```json
{
  "token": "[user token]", // string
  "user": {
    "id": "[user id]", // int
    "email": "[user email]", // string
    "full_name": "[user full name]", // string
    "created_at": "[user creation date]" // timestamptz
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
  "email": "[valid email]", // string
  "password": "[string of 6 characters]" // string
}
```

**Success Response**:`200 OK`

```json
{
  "token": "[user token]", // string
  "user": {
    "id": "[user id]", // int
    "email": "[user email]", // string
    "full_name": "[user full name]", // string
    "created_at": "[user creation date]" // timestamptz
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
NONE
```

**Success Response**:`200 OK`

```json
{
  "communities": db.Community[] // array of community objects
}
```

---

## Community

---

## Request
