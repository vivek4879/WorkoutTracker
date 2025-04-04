import { Link } from "react-router-dom";
import Calendar from "./Calendar/Calendar";
// import Calendar from "./Calendar/Calendar";
import "./Dashboard.css";
import { useEffect } from "react";

const Dashboard = () => {
  useEffect(() => {
    document.title = "Dashboard | My App";
  }, []);
  return (
    <div className="container">
      <div className="container">
        <div className="topnav">
          <a className="active" href="/dashboard">
            Home
          </a>
          <a href="/addexercise">Add Workout</a>
          <a href="/profile" style={{ float: "right" }}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              height="23px"
              viewBox="0 -960 960 960"
              width="23px"
              fill="#000000"
            >
              <path d="M235.2-279.59q51-38.52 113.52-60.54 62.52-22.02 131.33-22.02t131.64 22.38q62.83 22.38 113.11 60.42 34.29-41.24 53.31-91.92 19.02-50.69 19.02-108.73 0-131.57-92.78-224.35T480-797.13q-131.57 0-224.35 92.78T162.87-480q0 57.8 18.9 108.49 18.9 50.68 53.43 91.92ZM480-437.85q-59.96 0-100.93-40.86-40.98-40.86-40.98-100.81 0-59.96 40.98-100.94 40.97-40.97 100.93-40.97 59.96 0 100.93 40.97 40.98 40.98 40.98 100.94 0 59.95-40.98 100.81-40.97 40.86-100.93 40.86Zm-.02 365.98q-84.65 0-159.09-32.1-74.43-32.1-129.63-87.29-55.19-55.2-87.29-129.65-32.1-74.46-32.1-159.11 0-84.65 32.1-159.09 32.1-74.43 87.29-129.63 55.2-55.19 129.65-87.29 74.46-32.1 159.11-32.1 84.65 0 159.09 32.1 74.43 32.1 129.63 87.29 55.19 55.2 87.29 129.65 32.1 74.46 32.1 159.11 0 84.65-32.1 159.09-32.1 74.43-87.29 129.63-55.2 55.19-129.65 87.29-74.46 32.1-159.11 32.1Zm.02-91q51.8 0 97.37-14.9 45.56-14.9 84.56-42.95-39.47-28.28-84.44-43.06-44.97-14.79-97.49-14.79-52.52 0-97.37 14.79-44.85 14.78-84.33 43.06 39 28.05 84.45 42.95 45.45 14.9 97.25 14.9Zm0-358.56q25.04 0 41.57-16.53 16.52-16.52 16.52-41.56 0-25.05-16.52-41.69-16.53-16.64-41.57-16.64t-41.57 16.64q-16.52 16.64-16.52 41.69 0 25.04 16.52 41.56 16.53 16.53 41.57 16.53Zm0-58.09Zm.24 358.8Z" />
            </svg>
          </a>
          <a href="/login" style={{ float: "right" }}>
            Logout
          </a>
           <Link style={{ float: "right" }} to="/profile">
            Profile
          </Link>
           <Link style={{ float: "right" }} to="/measurements">
  Measurements
</Link>

         

          {/* <Link style={{ float: "right" }} to="/usergoals">
            UserGoal
          </Link> */}
        </div>
        <div className="Logo">
          <h1 className="Page-Heading">Gambare!</h1>
        </div>
        <div className="header">
          <h2 className="Page-Heading">Current Streak: 7</h2>
        </div>
        <Calendar />
        <div className="WorkoutContainer">
          <div className="Workouts">
            <h3>Past Workouts:</h3>
            <div className="WorkoutCard">
              <h3 className="cardText">Monday</h3>
              <p className="cardText">Ab Wheel</p>
              <p className="cardText">Sets: 7 Reps: 5</p>
              <p className="cardText">Arnold Press (Dumbbell)</p>
              <p className="cardText">Sets: 4 Reps: 9</p>
              <p className="cardText">Around the World</p>
              <p className="cardText">Sets: 7 Reps: 7</p>
              <p className="cardText">Ab Wheel</p>
              <p className="cardText">Sets: 3 Reps: 4</p>
            </div>
            <div className="WorkoutCard">
              <h3 className="cardText">Tuesday</h3>
              <p className="cardText">Aerobics</p>
              <p className="cardText">Sets: 6 Reps: 2</p>
              <p className="cardText">Ab Wheel</p>
              <p className="cardText">Sets: 4 Reps: 8</p>
            </div>
            <div className="WorkoutCard">
              <h3 className="cardText">Wednesday</h3>
              <p className="cardText">Back Extension</p>
              <p className="cardText">Sets: 9 Reps: 5</p>
              <p className="cardText">Ab Wheel</p>
              <p className="cardText">Sets: 3 Reps: 2</p>
              <p className="cardText">Arnold Press (Dumbbell)</p>
              <p className="cardText">Sets: 6 Reps: 7</p>
            </div>
            <div className="WorkoutCard">
              <h3 className="cardText">Thursday</h3>
              <p className="cardText">
                Back Extensions <br />
                (Machine)
              </p>
              <p className="cardText">Sets: 4 Reps: 7</p>
              <p className="cardText">Back Extension</p>
              <p className="cardText">Sets: 3 Reps: 4</p>
            </div>
            <div className="WorkoutCard">
              <h3 className="cardText">Friday</h3>
              <p className="cardText">Ball Slams</p>
              <p className="cardText">Sets: 6 Reps: 8</p>
              <p className="cardText">Ab Wheel</p>
              <p className="cardText">Sets: 5 Reps: 7</p>
              <p className="cardText">Around the World</p>
              <p className="cardText">Sets: 4 Reps: 2</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
