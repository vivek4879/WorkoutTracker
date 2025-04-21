import { useEffect, useState } from "react";
import "./AddExercise.css";

const AddExercise = () => {
  useEffect(() => {
    document.title = "Add Exercise";
  }, []);

  // State
  const [items, setItems] = useState([]);
  const [expandedItems, setExpandedItems] = useState({});
  const [personalBests, setPersonalBests] = useState({});
  const [setsData, setSetsData] = useState({});

  // Constant personal best data
  const personalBestData = {
    1: "150 lbs",
    2: "30 minutes",
    3: "75 lbs",
    4: "10 rounds",
    5: "100 lbs",
    6: "150 lbs",
    7: "50 reps",
    8: "60 seconds",
  };

  // Fetch exercises on mount
  useEffect(() => {
    fetch("http://192.168.0.12:4000/exercises")
      .then((res) => res.json())
      .then((data) => setItems(data))
      .catch((err) => console.error(err));
  }, []);

  // Toggle expand and initialize data
  const toggleExpand = (index, workoutId) => {
    setExpandedItems((prev) => ({ ...prev, [index]: !prev[index] }));

    if (!personalBests[workoutId]) {
      setPersonalBests((prev) => ({
        ...prev,
        [workoutId]:
          personalBestData[workoutId] || "No personal best available",
      }));
    }

    if (!setsData[workoutId]) {
      setSetsData((prev) => ({
        ...prev,
        [workoutId]: [{ setno: 1, weights: 0, repetitions: 0 }],
      }));
    }
  };

  // Add new set row
  const addSet = (workoutId) => {
    setSetsData((prev) => {
      const current = prev[workoutId] || [];
      return {
        ...prev,
        [workoutId]: [
          ...current,
          { setno: current.length + 1, weights: 0, repetitions: 0 },
        ],
      };
    });
  };

  // Handle input changes
  const handleInputChange = (workoutId, setIndex, field, value) => {
    setSetsData((prev) => {
      const arr = [...(prev[workoutId] || [])];
      arr[setIndex] = { ...arr[setIndex], [field]: Number(value) };
      return { ...prev, [workoutId]: arr };
    });
  };

  // Build payload
  const handleAddWorkout = () => {
    const payload = {
      workouts: Object.entries(setsData).map(([exerciseid, sets]) => ({
        exerciseid: Number(exerciseid),
        sets: sets.map(({ setno, weights, repetitions }) => ({
          setno,
          repetitions,
          weights,
        })),
        created_at: new Date().toISOString(),
      })),
    };
    console.log(JSON.stringify(payload, null, 2));
    // send payload to API
  };

  return (
    <div className="dashcontainer">
      <div className="topnav">
        <a href="/dashboard">Home</a>
        <a className="active" href="#">
          Add Workout
        </a>
        <a href="/profile" style={{ float: "right" }}>
          {/* profile icon */}
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
          <h2 className="exeListHead">Exercise List</h2>
          {items.map((item, index) => {
            const [id, name] = item;
            return (
              <div key={index}>
                <div
                  className="workouts"
                  onClick={() => toggleExpand(index, id)}
                  style={{ cursor: "pointer" }}
                >
                  {name}
                </div>
                {expandedItems[index] && (
                  <div className="expandedOptions">
                    <p className="pBest">
                      Personal Best: {personalBests[id] || "Loading..."}
                    </p>
                    <div className="workout-titles">
                      <p>Set</p>
                      <p>KG/s</p>
                      <p>Reps</p>
                    </div>
                    {setsData[id]?.map((set, setIndex) => (
                      <div className="workout" key={setIndex}>
                        <input
                          className="metrics"
                          type="number"
                          value={set.setno}
                          onChange={(e) =>
                            handleInputChange(
                              id,
                              setIndex,
                              "setno",
                              e.target.value
                            )
                          }
                          min="1"
                        />
                        <input
                          className="metrics"
                          type="number"
                          value={set.weights}
                          onChange={(e) =>
                            handleInputChange(
                              id,
                              setIndex,
                              "weights",
                              e.target.value
                            )
                          }
                          min="0"
                        />
                        <input
                          className="metrics"
                          type="number"
                          value={set.repetitions}
                          onChange={(e) =>
                            handleInputChange(
                              id,
                              setIndex,
                              "repetitions",
                              e.target.value
                            )
                          }
                          min="0"
                        />
                      </div>
                    ))}
                    <button className="addset" onClick={() => addSet(id)}>
                      Add Set
                    </button>
                  </div>
                )}
              </div>
            );
          })}
          <button className="addButton" onClick={handleAddWorkout}>
            Add Workout
          </button>
        </div>
      </div>
    </div>
  );
};

export default AddExercise;
