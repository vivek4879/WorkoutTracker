# 🏋 Workout Tracker - Sprint Summary

## 📌 User Stories
## 📝 Full List of User Stories

1. *Workout Logging* - As a user, I want to log my workouts so that I can track my progress.
2. *Detailed Workout Tracking* - As a user, I want to log my workouts along with the number of sets, repetitions, and the weights used so that I can monitor my strength progression.
3. *Water Intake Tracking* - As a user, I want to keep track of my daily water intake so that I can ensure I am drinking enough water.
4. *Calorie Tracking* - As a user, I want to track my calorie intake so that I know how much I am consuming.
5. *Body Measurements Logging* - As a user, I want to log my body measurements as and when I want so that I can track my progress.
6. *Workout History* - As a user, I want to be able to view my past workouts so that I can compare my progress.
7. *Workout Reminders* - As a user, I want to receive reminders for my scheduled workouts so that I can stay consistent and avoid missed sessions.
8. *Social Workouts* - As a user, I want to know what my friends are working out today and when so that we can work out together if possible.
9. *Progress Photos* - As a user, I want to be able to save photos of my progress so that I can track my progress visually as well.
10. *Data Visualization* - As a user, I want readily accessible graphs and data analysis of my workout history so that I can make informed decisions about my workouts.
11. *Pre-Built Workout Plans* - As a user, I want to be able to follow structured workout plans so that I can maintain a consistent fitness routine.
12. *Session Management* - As the developer, I want session management so that I can save the user’s workout to a database while they are in session and keep the user's data safe.

## 🔧 Issues Planned to Address

1. *User Login & Session Management*  
   - Implement user authentication via email and password.  
   - Establish session management where a session token (cookie-based authentication) is generated upon successful login.  
   - Ensure persistent login by storing and validating session cookies.  

2. *User Registration*  
   - Develop a sign-up feature that captures essential details: *Email, First Name, Last Name, and Password*.  
   - Implement email validation to prevent duplicate accounts.  
   - Securely store user credentials using password hashing (e.g., *bcrypt*).  

3. *OTP-Based Authentication*  
   - Integrate *one-time password (OTP) verification via email* for user authentication.  
   - OTP expiration and invalidation mechanisms to enhance security.  

4. *Forgot Password Functionality*  
   - Enable users to reset their password using a *password reset link or OTP* sent to their registered email.  
   - Implement secure token-based password reset flow with expiration.  

5. *UI/UX Design with Figma*  
   - Create high-fidelity *Figma designs* for user authentication screens, including:  
     - *Sign Up Page*  
     - *Login Page*  
     - *Forgot Password Page*  

6. *Frontend Implementation*  
   - Develop *React-based frontend* for user authentication.  
   - Implement responsive UI components for a seamless user experience.  
   - Use *React Router* for page navigation and *Redux/Context API* (if required) for state management.  

7. *Frontend & Backend Integration*  
   - Establish API communication between frontend and backend using *RESTful APIs* or *GraphQL*.  
   - Secure API endpoints with *JWT-based authentication*.  
   - Implement error handling and response validation on the frontend.  

8. *Swagger API Documentation*  
   - Integrate *Swagger (OpenAPI)* for API documentation.  
   - Enable automatic API contract generation for easy debugging and testing.  
   - Provide example request-response structures for better API usability.  

9. *Profile Deletion*
   - User should be able to delete profile.
---


## 🏆 Successfully Completed
1. *User Login & Session Management*  
   - Implemented user authentication via email and password.  
   - Established session management with *cookie-based authentication*, ensuring users remain logged in.  

2. *User Registration*  
   - Developed a sign-up feature that captures *Email, First Name, Last Name, and Password*.  
   - Implemented email validation to prevent duplicate accounts.  
   - Securely stored user credentials using *argon2id* for password hashing which uses the Argon2id algorithm variant and cryptographically-secure random salts.  

3. *UI/UX Design with Figma*  
   - Created high-fidelity *Figma designs* for authentication screens:  
     - *Sign Up Page*  
     - *Login Page*  
     - *Forgot Password Page*  

4. *Frontend Implementation*  
   - Developed *React-based frontend* for authentication.  
   - Implemented responsive UI components for a seamless user experience.  
   - Used *React Router* for page navigation.
5. *Profile Deletion*
   - User is able to delete profile.
---

## ❌ Not Completed
1. *OTP-Based Authentication*  
   - *Email OTP verification* was not implemented due to prioritization of basic authentication.  
   - Future plans to include *OTP-based login for enhanced security*.  

2. *Forgot Password Functionality*  
   - The password reset flow using *OTP or token-based authentication* was not fully implemented because I want to do it using kafka and had trouble implementing it.  
   - Future implementation will include *email-based password reset with expiration*.  

3. *Swagger API Documentation*  
   - API documentation using *Swagger (OpenAPI)* was not integrated because I was having trouble understanding it.  
   - Future enhancements will include *automatic API contract generation and request-response documentation*.  

---
