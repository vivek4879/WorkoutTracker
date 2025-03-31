## Detail work completed in Sprint 3

### 1. Frontend

### 2. Backend
In this sprint, several important APIs were added to enhance user interaction with the application. The new APIs focus on measurements,user performance tracking, and exercise management. Additionally, significant updates were made to the Add Workout API to track personal bests and update workout streaks.
    
<b>2.1. Update Measurements:</b> This API was introduced to allow users to update their body measurements, which include data like height, weight, and other fitness-related metrics.The user can send a JSON object containing updated body measurements, such as height, weight, chest, waist, and biceps.The API ensures that the user is authenticated by validating the session token before updating their measurements in the database. Upon successful update, a success message is returned. If the session is invalid or the data is malformed, appropriate error responses are sent.

<b>Key Functionality:</b>

- Allows users to update their personal fitness data.

- Validates session to ensure only authorized users can update measurements.

- Returns success or error messages based on the outcome.

<b>2.2. Get Measurements:</b> This API allows users to retrieve their current body measurements. It is useful for tracking progress over time.

- Description:

    - Response: The API fetches and returns the stored measurements for the authenticated user in JSON format.

    - Process: The API checks the session token to ensure the user is authenticated before retrieving the measurements from the database.


- Key Functionality:

    - Fetches stored measurements for the authenticated user.

    - Ensures that only logged-in users can retrieve their measurements.

    - Provides a structured response with height, weight, and other key measurements.

<b>2.3. Get User Best for Exercise :</b> This API allows users to track their best performance for a particular exercise.

- Description:

    - Query Parameter: Users specify the exercise_id in the query parameters.

    - Response: The user’s personal best for the given exercise, including the highest weight lifted and the corresponding repetitions, is returned.
    - Process: The API checks the session and fetches the user's best performance for the specified exercise from the database.

    
- Key Functionality:

    - Retrieves the user’s best performance for a specific exercise.

    - Ensures that the user is authenticated before fetching the data.

    - Provides the best weight and corresponding repetitions for the specified exercise.

<b>2.4. Get All Exercises :</b> This API provides a list of all available exercises in the system, enabling users to explore exercises they can log in their workout sessions.

- Description:

    - Response: A list of exercises with details like exercise_id, name, and URL (image or link related to the exercise).

    - Process: This API does not require session validation and returns a public list of exercises available in the system.


- Key Functionality:

    - Lists all available exercises in the system.

    - Provides exercise details such as name and image URL.

    - Helps users discover exercises to log in their workout sessions.

<b>2.5. Changes Made to the Add Workout API: </b>The Add Workout API was updated to not only add a workout for the user but also to track personal bests and update workout streaks.

- Description:

    - Request Body: Users provide workout details, including exerciseid, sets (with repetitions and weights), and the timestamp for when the workout was performed.

    - Process: The API follows these steps:
      - Session Validation: The session token is checked to ensure the user is authenticated.
      - Adding the Workout: The workout data is inserted into the database for the authenticated user.
      - Updating Best Performance: For each exercise in the workout, the API checks if the user’s current performance surpasses their personal best. If so, the new performance is recorded as their new best.
      - Updating Streaks: The API tracks the user's streak, i.e., the number of consecutive workouts logged. If the user continues to log workouts without breaks, their streak is updated accordingly.

- Response:

    - On success, the API returns a message confirming the workout was added, along with any updates to the user's best performance and streak.

    - If any error occurs (such as invalid session or workout data), an error message is returned.

- Key Updates:

    - Tracking Personal Bests: For each exercise logged, the API checks whether the new performance (maximum weight lifted) exceeds the user's previous best. If it does, the personal best is updated.

    - Streak Updates: The API also tracks the user's workout streak, encouraging consistent workout behavior. If a user logs workouts consecutively, their streak is updated.

    - Detailed Response: The response includes information about updated personal bests and the streak (old and new streak).

### Summary of Changes in This Sprint:

- User Data Management: APIs for updating and retrieving user measurements (update-measurements, get-measurements) were introduced, enabling users to track fitness metrics.

- Performance Tracking: New API (user-bests) allows users to see their best performance for specific exercises, providing an overview of their progress.

- Exercise Management: The exercises API lists all exercises available in the system, allowing users to explore available exercises.

- Workout Tracking Enhancements: The add-workout API was enhanced to track personal bests and update the user's streak based on consecutive workout logs.

These improvements add more value to the platform by making it easier for users to track their progress, keep motivated through streaks, and enhance their workout experience.
## Unit Tests for FrontEnd



## Unit tests for backend
#### 1. Access Tests (in route_access_test.go)

    1.1. TestSignupHandler: Ensures user signup works correctly.
    1.2. TestDeleteHandler: Ensures the function deletes user.

#### 2. Authentication Tests (in route_authentication_test.go)

    2.2. TestAuthenticationHandler: Validates login functionality.

#### 3. Dashboard Tests (in route_dashboard_test.go)

    3.1. TestDashboardHandler: Ensures authenticated users can access the dashboard.
    3.2. TestAddWorkoutHandler: Ensures users can add workouts.
    3.3. TestGetAllExercisesHandler: Ensures users get list of exercises when they want to add exercise.
    3.4. TestAddHandler: Ensures return of user's best for that particular exercise.

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
  "Firstname": "John",
  "Lastname": "Doe",
  "Email": "johndoe@example.com",
  "Password": "securepassword"
}
```

<b>Response</b>:

`201 Created` if signup is successful.

`409 Conflict` if the email already exists.

`500 Internal Server Error` for other failures.


### <b>2. User Authentication</b>

<b>Endpoint</b>: POST /authenticate

<b>Description</b>: Authenticates a user and starts a session.

<b>Request Body (JSON)</b>:

```
{
  "Email": "johndoe@example.com",
  "Password": "securepassword"
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


### <b>3. User Logout</b>

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


### <b>4. Dashboard</b>

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

### <b>5. Add Workout</b>


<b>Endpoint</b>: POST /add-workout

<b>Description</b>: Adds a workout for the authenticated user. Also checks for the User's best for all the exercises and updates the best if there is a new best. Check's the user's streak details and updates it.

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
    "Old_streak": {
        "UserID": 6,
        "LastWorkoutDate": "2025-03-30T13:54:16.563934-04:00",
        "CurrentStreak": 1,
        "MaxStreak": 1
    },
    "message": "Workout added successfully",
    "new_streak": {
        "UserID": 6,
        "LastWorkoutDate": "2025-03-30T13:54:16.563934-04:00",
        "CurrentStreak": 1,
        "MaxStreak": 1
    },
    "updated_bests": null
}
```

<b>Status Codes</b>:

`201 Created` on success.

`400 Bad Request` for invalid JSON.

`401 Unauthorized` if session is invalid.

### <b>6. Delete User</b>

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

<b>Status Codes</b>:

`200 OK` on success.

`401 Unauthorized` if session is invalid or expired.
`401 Unauthorized` if user retrieval fails.
`401 Unauthorized` if deletion fails.
`401 Unauthorized` if session deletion fails.



### <b>7. Update Measurements</b>

**Endpoint**: `PUT /update-measurements`

**Description**: Updates Measurements for the user depending on the existing data available. If there is no data then it just enters the new data.

<b>Request Body (JSON)</b>:

```json
{
  "weight": 80.0,
  "neck": 38.0,
  "chest": 100.0
}

```

**Response**:

```json
{
  "message": "Measurements updated successfully"
}
```

<b>Status Codes</b>:

`200 OK` on success.

`401 Unauthorized` if session is invalid or expired.
`400 Badrequest` if JSON invalid.
`500 InternalServerError` if update fails.


### <b>8.  Measurements</b>

**Endpoint**: `GET /measurements`

**Description**: Returns the user's measurements if available. Returns null values if fields are blank. 

**Response**:

```json
{
  "userid": 6,
  "weight": 80,
  "neck": 38,
  "chest": 100,
  "user": {
    "ID": 0,
    "FirstName": "",
    "LastName": "",
    "Email": "",
    "Password": ""
  }
}
```

<b>Status Codes</b>:

`200 OK` on success.
`401 Unauthorized` if session is invalid or expired.
`500 InternalServerError` if retrieval fails.


### <b>9.  user-bests</b>

**Endpoint**: `GET /user-bests`

**Description**: Returns the user's bests if available for the particular exercises. Returns null values if fields are blank.

**Query Params**:

- `key: exercise_id; Value:3`

**Response**:

```json
{
  "best_weight": 50,
  "exercise_id": 3,
  "reps": 10,
  "user_id": 6
}
```

<b>Status Codes</b>:

`200 OK` on success.
`401 Unauthorized` if session is invalid or expired.
`500 InternalServerError` if retrieval fails.


### <b>10. Exercises</b>

**Endpoint**: `GET /exercises`

**Description**: Returns the list of all exercises available in the database.

**Response**:

```json
[
  {
    "exercise_id": 2,
    "name": "Ab Wheel",
    "url": "https://workoutexercises.s3.us-east-2.amazonaws.com/AbWheel.jpeg"
  },
  {
    "exercise_id": 3,
    "name": "Arnold Press",
    "url": "https://workoutexercises.s3.us-east-2.amazonaws.com/ArnoldPress.jpeg"
  },
  {
    "exercise_id": 4,
    "name": "Around The World",
    "url": "https://workoutexercises.s3.us-east-2.amazonaws.com/AroundTheWorld.jpeg"
  },
  {
    "exercise_id": 5,
    "name": "BenchPress(Smith Machine)",
    "url": "https://workoutexercises.s3.us-east-2.amazonaws.com/BenchPress(Smith+Machine).jpeg"
  },
  {
    "exercise_id": 6,
    "name": "Bicep Curl",
    "url": "https://workoutexercises.s3.us-east-2.amazonaws.com/BicepCurl.jpeg"
  }
]
```

<b>Status Codes</b>:

`200 OK` on success.

`401 Unauthorized` if session is invalid or expired.
`500 InternalServerError` if retrieval fails.