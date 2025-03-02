## Detail work  completed in Sprint 2

### 1. Backend
In the Backend this sprint focused on implementing an interface for user management, developing workout related functionality, adding structured response handling, defining some new database tables and tests for the functions.

1.1 Implementing the `UserModelInterface`

One of the key improvements made in Sprint 2 was defining and implementing an interface (UserModelInterface) to enforce a standard structure for user-related operations in the database.This interface, located in models.go, defines essential functions for managing users, sessions, and workouts.
With this interface in place, the MyModel struct was updated to implement all the required functions. This implementation ensures that any future changes to database operations can be integrated seamlessly.

#### Why was this interface introduced?
* To standardize database operations across multiple models.
* To enable mocking in unit tests, making the codebase more testable and maintainable.
* To simplify future enhancements by allowing different database implementations to adhere to the same contract.

1.2 Implementation of `AddWorkoutHandler`

Introduced the AddWorkoutHandler in route_access.go to allow authenticated users to add workouts.
It Validates User Session before allowing modifications, parses JSON input to extract workout details,
Inserts Data into `workouts` and `workout_to_user` tables, and return appropriate success/error messages.

1.3 New Tables for Workouts & Exercises

To support workouts, I introduced two new tables in models.go which store information about workouts associated with a user.

1.4 Implementing Error and Success Response Functions

To standardize API responses, we introduced sendErrorResponse and sendSuccessResponse functions in utils.go.
The error response function handles API errors and returns JSON response. The success response function handles successful API responses.

1.5 Unit Testing and Testify library 

I wrote unit tests using the Testify library to validate API behavior. The Testify library  provides mocking for database interactions, simplifies assertions with assert functions, and makes unit tests more readable and maintainable.




### 2. Frontend


## Unit tests and Cypress test for frontend

## Unit tests for backend

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

