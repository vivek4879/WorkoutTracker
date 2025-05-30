import Calendar from "./Calendar/Calendar";
// import Calendar from "./Calendar/Calendar";
import "./Dashboard.css";
import { useEffect } from "react";
import { Link } from "react-router-dom";
import {
  CRow,
  CCol,
  CCard,
  CCardHeader,
  CCardBody,
  CCardText,
} from "@coreui/react";
const Dashboard = () => {
  useEffect(() => {
    document.title = "Dashboard | My App";
  }, []);
  const workouts = [
    {
      day: "Wednesday",
      exercises: [
        { name: "Ab Wheel", sets: 7, reps: 5 },
        { name: "Arnold Press (Dumbbell)", sets: 4, reps: 9 },
        { name: "Around the World", sets: 7, reps: 7 },
        { name: "Ab Wheel", sets: 3, reps: 4 },
      ],
    },
    {
      day: "Thursday",
      exercises: [
        { name: "Aerobics", sets: 6, reps: 2 },
        { name: "Ab Wheel", sets: 4, reps: 8 },
      ],
    },
    {
      day: "Friday",
      exercises: [
        { name: "Back Extension", sets: 9, reps: 5 },
        { name: "Ab Wheel", sets: 3, reps: 2 },
        { name: "Arnold Press (Dumbbell)", sets: 6, reps: 7 },
      ],
    },
    {
      day: "Saturday",
      exercises: [
        { name: "Back Extensions (Machine)", sets: 4, reps: 7 },
        { name: "Back Extension", sets: 3, reps: 4 },
      ],
    },
    {
      day: "Sunday",
      exercises: [
        { name: "Ball Slams", sets: 6, reps: 8 },
        { name: "Ab Wheel", sets: 5, reps: 7 },
        { name: "Around the World", sets: 4, reps: 2 },
      ],
    },
  ];
  return (
    <div className="dashcontainer">
      <div className="topnav">
        <a className="active" href="/dashboard">
          Home
        </a>
        <Link to="/addexercise">Add Workout</Link>
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
        <Link to="/login" style={{ float: "right" }}>
          Logout
        </Link>
        <a href="/measurements" style={{ float: "right" }}>
          Measurements
        </a>
      </div>
      <div className="Logo">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>
      <div className="Header">
        <h2 className="Page-Heading">Current Streak: 7</h2>
      </div>
      <Calendar />
      <div className="WorkoutContainer">
        <div className="Workouts">
          <h3>Past Workouts:</h3>
          <CRow className="justify-content-center">
            {workouts.map((workout) => (
              <CCol key={workout.day}>
                <CCard className="h-100 w-80">
                  <CCardHeader as="h4" className="fw-bold text-center">
                    {workout.day}
                  </CCardHeader>
                  <CCardBody>
                    {workout.exercises.map((exercise, idx) => (
                      <div key={idx} className="mb-1">
                        <CCardText className="mb-1">{exercise.name}</CCardText>
                        <CCardText>
                          <small>
                            Sets: {exercise.sets} Reps: {exercise.reps}
                          </small>
                        </CCardText>
                      </div>
                    ))}
                  </CCardBody>
                </CCard>
              </CCol>
            ))}
          </CRow>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
