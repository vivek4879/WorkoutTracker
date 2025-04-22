#  Workout Tracker Application

A full-stack workout tracking web application built with Golang (backend), React + Vite (frontend), and PostgreSQL (database). The application allows users to log workouts, track personal bests, view streak progress, and more. In this project we are working on a web application to track workouts, fitness activities, and nutrition. The user will be able to register and log workouts corresponding to a date to track progress. User can also access their workout history using various parameters like workout type, day, frequency of each type, muscle targeted. The total workout time will also be logged to track effectivenes of each workout. At its core, this is a workout and nutrition Tracker, and we plan to add other functions like suggesting workouts on the basis of the user's workout history, suggestions on the ideal exercise times for the user in the future.

---

##  Prerequisites

Ensure the following tools are installed before setting up the application:

### 1.  PostgreSQL

- **Version:** PostgreSQL 13 or above recommended
- **Install:**
    - **macOS:** `brew install postgresql`
    - **Ubuntu:** `sudo apt install postgresql postgresql-contrib`
    - **Windows:** [Download from official site](https://www.postgresql.org/download/)

- **PostgreSQL Setup:**
    - Create a database:
      ```bash
      createdb workout_tracker
      ```
    - Create a user and set privileges:
      ```bash
      createuser workout_user --pwprompt
      psql -d workout_tracker -c "GRANT ALL PRIVILEGES ON DATABASE workout_tracker TO workout_user;"
      ```

### 2. Golang

- **Version:** Go 1.20 or later
- **Install:**
    - **macOS:** `brew install go`
    - **Ubuntu:** `sudo apt install golang`
    - **Windows:** [Download from official site](https://go.dev/doc/install)

- **Verify installation:**
  ```bash
  go version

### 3.  React + Vite

- **Node.js(with npm) required**
- **Install:** [Download Node.js](https://nodejs.org/en)

- **Verify installation:**
  ```bash
  node -v
  npm -v

- **Install Vite:** 

  ```bash
  npm create vite@latest
  cd frontend
  npm install
  

# Project Setup

## Backend (Golang)
1. Navigate to the backend folder
```bash
  cd backend
```
2. Navigate to the backend folder
```bash
  go mod tidy
```
3. Set up environment variables in a .env file:
```bash
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=workout_user
    DB_PASSWORD=your_password
    DB_NAME=workout_tracker

```
4. NRun the server
```bash
  go run main.go
```

## Frontend (React + Vite)
1. Navigate to the frontend folder
```bash
  cd frontend
```
2. Install dependencies
```bash
  npm install
```
3. Start development server:
```bash
    npm run dev
```
4. Access the app at   http://localhost:5173


## Running Tests (Golang)
1. Backend Tests
```bash
  go test -v
```
2. Frontend Tests
```bash
  npm run test
```
# Members
1. Vivek Aher 
2. Anurag Kelkar 
3. Rishitha Adepu 

 
