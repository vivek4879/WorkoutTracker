import React, { useEffect, useState } from "react";
// import axios from "axios"; // API call commented out
import "./AddExercise.css";
const AddExercise = () => {
  useEffect(() => {
    document.title = "Add Exercise";
  }, []);

  // State for expanded workouts, personal bests, and sets for each workout
  const [expandedItems, setExpandedItems] = useState({});
  const [personalBests, setPersonalBests] = useState({});
  const [setsData, setSetsData] = useState({});

  // Constant personal best data
  const personalBestData = {
    1: "150 lbs", // Ab Wheel
    2: "30 minutes", // Aerobics
    3: "75 lbs", // Arnold Press (Dumbbell)
    4: "10 rounds", // Around the World
    5: "100 lbs", // Back Extension
    6: "150 lbs", // Back Extensions (Machine)
    7: "50 reps", // Ball Slams
    8: "60 seconds", // Battle Ropes
  };

  // Toggle expanded state for a workout and initialize personal best and sets if needed
  const toggleExpand = (index, workoutId) => {
    setExpandedItems((prev) => ({ ...prev, [index]: !prev[index] }));

    if (!personalBests[workoutId]) {
      // Uncomment and replace the URL when integrating the API
      // axios.get(`/api/getPersonalBest/${workoutId}`)
      //   .then(response => {
      //     setPersonalBests(prev => ({
      //       ...prev,
      //       [workoutId]: response.data.personalBest,
      //     }));
      //   })
      //   .catch((error) => console.error("Error fetching personal best:", error));

      // Use constant data for now
      setPersonalBests((prev) => ({
        ...prev,
        [workoutId]:
          personalBestData[workoutId] || "No personal best available",
      }));
    }

    // Initialize the sets for this workout if not already set
    if (!setsData[workoutId]) {
      setSetsData((prev) => ({ ...prev, [workoutId]: [{}] }));
    }
  };

  // Add a new set row for a given workout
  const addSet = (workoutId) => {
    setSetsData((prev) => ({
      ...prev,
      [workoutId]: [...(prev[workoutId] || []), {}],
    }));
  };

  const items = [
    [1, "Ab Wheel"],
    [2, "Aerobics"],
    [3, "Arnold Press (Dumbbell)"],
    [4, "Around the World"],
    [5, "Back Extension"],
    [6, "Back Extensions (Machine)"],
    [7, "Ball Slams"],
    [8, "Battle Ropes"],
  ];

  return (
    <div className="container">
      <div className="topnav">
        <a href="#">Home</a>
        <a className="active" href="#">
          Add Workout
        </a>
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
      </div>
      <div className="Logo">
        <h1 className="Page-Heading">Select Exercise</h1>
      </div>
      <div className="addContainer">
        <div className="selectExercise">
          <h2>Exercise List</h2>
          {items.map((item, index) => (
            <div key={index}>
              <div
                className="workouts"
                onClick={() => toggleExpand(index, item[0])}
                style={{ cursor: "pointer" }}
              >
                {item[1]}
              </div>
              {expandedItems[index] && (
                <div className="expandedOptions">
                  <p>Personal Best: {personalBests[item[0]] || "Loading..."}</p>
                  <div className="workout-titles">
                    <p>Set</p>
                    <p>KG/s</p>
                    <p>Reps</p>
                  </div>
                  {setsData[item[0]] &&
                    setsData[item[0]].map((set, setIndex) => (
                      <div className="workout" key={setIndex}>
                        <input
                          className="metrics"
                          type="number"
                          placeholder="0"
                          min="0"
                          max="10"
                        />
                        <input
                          className="metrics"
                          type="number"
                          placeholder="0"
                          min="0"
                          max="50"
                        />
                        <input
                          className="metrics"
                          type="number"
                          placeholder="0"
                          min="0"
                          max="10"
                        />
                      </div>
                    ))}
                  <button className="addset" onClick={() => addSet(item[0])}>
                    Add Set
                  </button>
                </div>
              )}
            </div>
          ))}
          <button className="addButton" type="submit">
            Add Workout
          </button>
        </div>
      </div>
    </div>
  );
};

export default AddExercise;
