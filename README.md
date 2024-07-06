# Org-authenticator-api

This project implements a user authentication and organisation management system using Go.
The API supports user registration, login, and organisation management functionalities.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Database Setup](#database-setup)
- [Models](#models)
- [Endpoints](#endpoints)
- [Testing](#testing)

## Features

- User Registration and Login
- JWT Authentication
- Organisation Creation and Management
- User Organisation Association
- Field Validation and Error Handling

## Technologies Used

- Backend Language/Framework: Go
- Database: PostgreSQL
- ORM: TBD (optional)
- Authentication: JWT (JSON Web Tokens)

## Database Setup

1. Install PostgreSQL and create a database.
2. Configure your application to connect to the PostgreSQL database.
3. Optionally, set up an ORM of your choice.

## Models

### User Model

```json
{
    "userId": "string", // must be unique
    "firstName": "string", // must not be null
    "lastName": "string", // must not be null
    "email": "string", // must be unique and must not be null
    "password": "string", // must not be null
    "phone": "string"
}
```

### Organisation Model

```json
{
    "orgId": "string", // Unique
    "name": "string", // Required and cannot be null
    "description": "string"
}
```

## Endpoints

### User Authentication & Register User

Endpoint: POST /auth/register

**Request Body:**

```json
{
    "firstName": "string",
    "lastName": "string",
    "email": "string",
    "password": "string",
    "phone": "string"
}
```

**Successful Response:**

```json
{
    "status": "success",
    "message": "Registration successful",
    "data": {
        "accessToken": "eyJh...",
        "user": {
            "userId": "string",
            "firstName": "string",
            "lastName": "string",
            "email": "string",
            "phone": "string"
        }
    }
}
```

**Unsuccessful Response:**

```json
{
    "status": "Bad request",
    "message": "Registration unsuccessful",
    "statusCode": 400
}
```

### Login User

Endpoint: POST /auth/login

**Request Body:**

```json
{
    "email": "string",
    "password": "string"
}
```

**Successful Response:**

```json
{
    "status": "success",
    "message": "Login successful",
    "data": {
        "accessToken": "eyJh...",
        "user": {
            "userId": "string",
            "firstName": "string",
            "lastName": "string",
            "email": "string",
            "phone": "string"
        }
    }
}
```

**Unsuccessful Response:**

```json
{
    "status": "Bad request",
    "message": "Authentication failed",
    "statusCode": 401
}
```

## User Endpoints

### Get User Details

Endpoint: GET /api/users/:id

**Successful Response:**

```json
{
    "status": "success",
    "message": "<message>",
    "data": {
        "userId": "string",
        "firstName": "string",
        "lastName": "string",
        "email": "string",
        "phone": "string"
    }
}
```

## Organisation Endpoints

### Get All Organisations

Endpoint: GET /api/organisations

**Successful Response:**

```json
{
    "status": "success",
    "message": "<message>",
    "data": {
        "organisations": [
            {
                "orgId": "string",
                "name": "string",
                "description": "string"
            }
        ]
    }
}
```

### Get Single Organisation

Endpoint: GET /api/organisations/:orgId

**Successful Response:**

```json
{
    "status": "success",
    "message": "<message>",
    "data": {
        "orgId": "string",
        "name": "string",
        "description": "string"
    }
}
```

### Create Organisation

Endpoint: POST /api/organisations

**Request Body:**

```json
{
    "name": "string",
    "description": "string"
}
```

**Successful Response:**

```json
{
    "status": "success",
    "message": "Organisation created successfully",
    "data": {
        "orgId": "string",
        "name": "string",
        "description": "string"
    }
}
```

**Unsuccessful Response:**

```json
{
    "status": "Bad Request",
    "message": "Client error",
    "statusCode": 400
}
```

### Add User to Organisation

Endpoint: POST /api/organisations/:orgId/users

**Request Body:**

```json
{
    "userId": "string"
}
```

**Successful Response:**

```json
{
    "status": "success",
    "message": "User added to organisation successfully"
}
```

## Testing

- Unit Testing

- Token generation: Ensure token expires at the correct time and correct user details are found in token.

- Organisation: Ensure users can’t see data from organisations they don’t have access to.

- End-to-End Test Requirements for the Register Endpoint

- Directory Structure: Create a tests folder with the test file named `auth.spec.ext`.

### Test Scenarios:

- Register user successfully with default organisation.
- Log the user in successfully.
- Fail if required fields are missing.
- Fail if there’s a duplicate email or userID.