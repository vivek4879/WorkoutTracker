


## Backend API Documentation
### API Endpoints

### <b>1. User Signup</b>

<b>Endpoint</b>: POST /signup

<b>Description</b>: Registers a new user.

<b>Request Body (JSON)</b>:
```
{
  "firstname": "John",
  "lastname": "Doe",
  "email": "johndoe@example.com",
  "password": "securepassword"
}
```
<b>Response</b>:

`201 Created` if signup is successful.

`409 Conflict` if the email already exists.

`500 Internal Server Error` for other failures.

### 2. User Authentication

<b>Endpoint</b>: POST /authenticate

<b>Description</b>: Authenticates a user and starts a session.

<b>Request Body (JSON)</b>:

```
{
  "email": "johndoe@example.com",
  "password": "securepassword"
}
```
<b>Response</b>:


``` 
{
  "message": "Authentication successful",
  "session_token": "mock-session-token",
  "user_id": "1"
} 
```

<b>Status Codes</b>:

`200 OK` on success.

`401 Unauthorized` on incorrect credentials.

### 3. User Logout

<b>Endpoint</b>: POST /logout

<b>Description</b>: Logs the user out and invalidates their session.

<b>Response</b>:

```
{
  "message": "User successfully logged out"
}
```

<b>Status Codes</b>:

`200 OK` on success.

`500 Internal Server Error` if session deletion fails.

### 4. Dashboard



<b>Endpoint</b>: GET /dashboard


<b>Description</b>: Fetches the user dashboard.


<b>Response</b>:

```
{
  "message": "Welcome to the dashboard page"
}
```
<b>Status Codes</b>:

`200 OK` on success.

`401 Unauthorized` if session is invalid.


### 5. Add Workout

<b>Endpoint</b>: POST /add-workout

<b>Description</b>: Adds a workout for the authenticated user.

<b>Request Body (JSON)</b>:

```
{
  "workouts": [
    {
      "exerciseid": 1,
      "sets": [
        {"setno": 1, "repetitions": 10, "weights": 50},
        {"setno": 2, "repetitions": 8, "weights": 55}
      ],
      "created_at": "2025-02-25T12:00:00Z"
    }
  ]
}

```

<b>Response</b>:
```
{
  "message": "Workout added successfully"
}
```

<b>Status Codes<b>:

`201 Created` on success.

`400 Bad Request` for invalid JSON.

`401 Unauthorized` if session is invalid.

### 6. Delete User

**Endpoint**: `DELETE /delete`

**Description**: Deletes an authenticated user's account and session.

**Request Headers**:
- `Cookie: session_token=<user_session_token>`

**Response**:
```json
{
  "message": "User successfully deleted"
}
```
<b>Status Codes<b>:

`200 OK` on success.

`401 Unauthorized` if session is invalid or expired.
`401 Unauthorized` if user retrieval fails.
`401 Unauthorized` if deletion fails.
`401 Unauthorized` if session deletion fails.

### Unit Tests

#### Backend Unit Tests:

#### The following unit tests have been implemented to validate backend functionality:

#### 1. Access Tests (in route_access_test.go)

    1.1. TestSignupHandler: Ensures user signup works correctly.
    1.2. TestDeleteHandler: Ensures the function deletes user.

#### 2. Authentication Tests (in route_authentication_test.go)

    2.2. TestAuthenticationHandler: Validates login functionality.

#### 3. Dashboard Tests (in route_dashboard_test.go)

    3.1. TestDashboardHandler: Ensures authenticated users can access the dashboard.
    3.2. TestAddWorkoutHandler: Ensures users can add workouts.

#### 4. Other Tests (in utils_test.go)
    4.1. TestSession: Validates session retrieval.
    4.2. TestHashing: Validates hashing of the password.
    4.3. TestSendErrorResponse: Checks if the function is sending the required error.
    4.4. TestSendSuccessResponse: Check if we receive a success response.


