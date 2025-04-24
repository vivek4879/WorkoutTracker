## Detail work completed in Sprint 3

### 1. Frontend

# Frontend Sprint Report

In this sprint, we delivered several new UI components, enhanced existing pages, and integrated them with the backend APIs to create a seamless, interactive experience for the user. Key focuses included workout display, streak visualization, measurement tracking, and responsive design.

---

## 1. Workout List Component

- **Dynamic Rendering**  
  Replaced static markup with a `WorkoutList` component that maps over workout data to generate list items.

- **Expandable Details**  
  Clicking on a day header toggles dropdowns showing individual exercise details (name, sets, reps).

- **Editable Dropdowns**  
  Users can adjust `sets` and `reps` via `<select>` controls; state updates propagate immediately.

- **Styling & Accessibility**  
  Added cursor pointers on headers and semantic `<ul>/<li>` structure for screen-reader compatibility.

---

## 2. Streak Visualization

- **Streak Calendar Widget**  
  Integrated the streak-data API to fetch and render user's current and longest streak on the dashboard.

- **Conditional Rendering**  
  Show loading spinners until API returns; display error messages if session validation fails.

- **Real-time Updates**  
  Calendar highlights update immediately after a workout is logged, using WebSocket events or polling (as configured).

---

## 3. Measurement Tracking UI

- **Measurement Form**  
  Created a form component for `update-measurements`, allowing users to input metrics (weight, body fat, etc.).

- **Data Retrieval**  
  On mount, calls `get-measurements` API to pre-fill fields with existing values.

- **Validation & Feedback**  
  Inline validation for numeric ranges; success and error toasts inform the user of API responses.

---

## 4. API Integration & State Management

- **Centralized Store**  
  Leveraged React Context or Redux (as appropriate) to hold workout data, measurements, and streaks globally.

- **Error Handling**  
  Uniform error boundary component catches network and rendering errors, prompting the user to retry.

- **Session Management**  
  Axios interceptors include `Authorization` token in headers and redirect to login on 401 responses.

---

## 5. Responsive & Mobile Design

- **Flex & Grid Layouts**  
  Replaced fixed widths with CSS Grid for workout cards and Flexbox for form layouts to support small screens.

- **Media Queries**  
  Ensured font sizes, paddings, and dropdown widths adjust gracefully from mobile (320px) to desktop (1920px).

---

## 6. Testing & Quality Assurance

- **Unit Tests** (Jest + React Testing Library)

  - `WorkoutList` rendering and toggle behavior
  - Dropdown change events updating component state
  - API-mocking for measurement form submission

- **End-to-End Tests** (Cypress)
  - Workout expansion and data modification
  - Streak calendar display under valid/invalid sessions
  - Measurement form validation and submission errors

---

## Summary of Frontend Changes

- Built dynamic **WorkoutList** with expandable, editable dropdowns for exercises.
- Integrated **streak** and **measurement** endpoints into dashboard and forms.
- Centralized state management with consistent error/session handling.
- Enhanced responsiveness and accessibility across components.
- Increased test coverage with unit and end-to-end scenarios.

---

# Frontend Testing Matrix

---

## Cypress E2E Tests

#### 1. AddExercise.component.cy.js

1. Should render the exercise list
2. Should render the calendar
3. Should check if logout works
4. Should check if dropdown menu works
5. Should check if workout is added to database
6. Should check if Measurements works

#### 2. Dashboard.component.cy.js

1. Should render the max streak
2. Should render the calendar
3. Should check if logout works
4. Should check if addWorkout works
5. Should check if workout cards are rendered
6. Should check if Measurements works

#### 3. Measurements.cy.js

1. Should render the exercise list
2. Should render the measurements form and data
3. Should check if logout works
4. Should check if dropdown menu works
5. Should check if new measurements are added to database
6. Should check if Measurements works

#### 4. forgotpass.cy.js

1. Should render the “Forgot Password” form

#### 5. login.cy.js

1. Should show the login form
2. Should allow a user to log in

#### 6. signup.cy.js

1. Should render the signup form fields
2. Should validate first name input
3. Should validate last name input
4. Should validate email input
5. Should validate password input
6. Should show validation errors appropriately

#### 7. userprof.cy.js

1. Should render the exercise list on the profile page
2. Should render the measurements
3. Should check if logout works
4. Should check if Home navigation works

---

## Unit Tests for Frontend Components

#### 1. WorkoutList Component (WorkoutList.spec.js)

1. Renders one `<li>` per workout day from props/state
2. Toggles exercise details when a day header is clicked
3. Updates internal state when “Sets” dropdown changes
4. Updates internal state when “Reps” dropdown changes
5. Applies correct ARIA attributes for expanded/collapsed state

#### 2. StreakCalendar Component (StreakCalendar.spec.js)

1. Displays loading indicator before data arrives
2. Fetches streak data on mount and renders current/max streak
3. Shows an error message if the API returns a 4xx/5xx
4. Updates highlighted dates when new streak data is pushed

#### 3. MeasurementsForm Component (MeasurementsForm.spec.js)

1. Fetches and pre‑fills existing measurements on mount
2. Validates numeric inputs (e.g. weight > 0, body fat 0–100)
3. Submits updated measurements and shows a success toast
4. Handles and displays API errors on submission failure
5. Disables submit button while request is in flight

#### 4. LoginForm Component (LoginForm.spec.js)

1. Renders email and password fields plus submit button
2. Validates that both fields are required
3. Calls login API with correct payload on submit
4. Redirects to dashboard on successful login
5. Displays error message on authentication failure

#### 5. SignupForm Component (SignupForm.spec.js)

1. Renders first name, last name, email, password, and submit button
2. Validates required fields and email format
3. Calls signup API with correct payload
4. Redirects or shows confirmation on success
5. Displays server‑side validation errors (e.g. “email already exists”)

#### 6. UserProfile Component (UserProfile.spec.js)

1. Renders user’s name and basic info from context/store
2. Renders list of past workouts via `WorkoutList`
3. Renders measurements via `MeasurementsForm`
4. Allows logout and redirects to login on click
5. Displays loading and error states for profile data fetch

---

### 2. Backend

In this sprint, several important APIs were added and some updated to enhance user interaction with the application. The updated and new APIs focus on measurements,user performance tracking, exercise management, and updating the streaks for the user and showing the same on the front end. Additionally, significant updates were made to some existing tests and extensive unhappy path tests were also added.

<b>2.1. Streak Data:</b> This API returns the user's streak data so that it could be displayed on the dashboard. It helps to keep the dashboard updated for the user.

<b>Key Functionality:</b>

- The calendar on the dashboard gets updated using this API. This way the user will know about the current streak and the max streak.

- Validates session to ensure only authorized users are able to see streaks.

- Returns success or error messages based on the outcome.

In addition to this, various tests were added for extensive testing which will be discussed in the tests section.

### Summary of Changes in This Sprint:

- Streak Data Management: Worked on streak data management for seamless display of streak data for the user. This will help the user to stay motivated.

- User Data Management: APIs for updating and retrieving user measurements (update-measurements, get-measurements) were improved enabling users to track fitness metrics.

- Extensive testing for various paths.

These improvements add more value to the platform by making it easier for users to track their progress, keep motivated through streaks, and enhance their workout experience.

## Unit Tests for FrontEnd

## Unit tests for backend

#### 1. Access Tests (in route_access_test.go)

    1.1. TestSignupHandler: Ensures user signup works correctly.
    1.2. TestDeleteHandler: Ensures the function deletes user.
    1.3. TestSignupHandler_UserAlreadyExists: Ensures that duplicate email registrations are rejected with an appropriate error response.
    1.4. TestDeleteHandler_SessionNotFound: Simulates failure cases where the session is missing or the user cannot be found during deletion.
    1.5. TestDeleteHandler_MissingSessionToken: Simulates failure cases where the session is missing or the user cannot be found during deletion.

#### 2. Authentication Tests (in route_authentication_test.go)

    2.1. TestAuthenticationHandler: Validates login functionality.
    2.2. TestAuthenticationHandler_UserNotFound: Confirms that login fails for nonexistent users or incorrect passwords, and no session is created.
    2.3. TestAuthenticationHandler_InvalidPassword

#### 3. Dashboard Tests (in route_dashboard_test.go)

    3.1. TestDashboardHandler: Ensures authenticated users can access the dashboard.
    3.2. TestAddWorkoutHandler: Ensures users can add workouts.
    3.3. TestGetAllExercisesHandler: Ensures users get list of exercises when they want to add exercise.
    3.4. TestAddHandler: Ensures return of user's best for that particular exercise.
    3.5. TestGetAllExercisesHandler_DBError: Simulates a backend failure when fetching exercises and checks for a 500 error response.
    3.6. TestDashboardHandler_InvalidTokenInDB: Checks unauthorized access behavior when session token is missing or invalid.
    3.7. TestDashboardHandler_MissingSessionCookie: Simulates failure cases where the session is missing or the user cannot be found during deletion.
    3.8. TestAddWorkoutHandler_InsertWorkoutFails: Simulates a failure in inserting workout data and confirms that an appropriate error is returned without calling downstream logic.
    3.9. TestAddHandler_QueryUserBestFails: Tests the failure scenario when best performance data cannot be retrieved, returning an appropriate server error.
    3.10. TestGetStreakDataHandler_FetchStreakFails: Simulates a failure to fetch streak data and checks for a proper 500 error response.

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

### <b>8. Measurements</b>

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

### <b>9. user-bests</b>

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

### <b>11. Streak</b>

**Endpoint**: `GET /get-streak-data`

**Description**: Returns the list of streak for the user available in the database.

**Response**:

```json
{
  "currentStreak": 1,
  "maxStreak": 1,
  "user_id": 6
}
```

<b>Status Codes</b>:

`200 OK` on success.

`401 Unauthorized` if session is invalid or expired.
`500 InternalServerError` if retrieval fails.
