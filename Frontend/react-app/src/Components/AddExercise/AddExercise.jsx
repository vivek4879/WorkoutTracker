import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import "./AddExercise.css";

const AddExercise = () => {
  const navigate = useNavigate();

  // UI state
  const [items, setItems] = useState([]);
  const [expandedItems, setExpandedItems] = useState({});
  const [personalBests, setPersonalBests] = useState({});
  const [setsData, setSetsData] = useState({});
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  // Static personal bests
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
    const fetchExercises = async () => {
      const token = localStorage.getItem("token");

      try {
        const response = await axios.get("http://192.168.0.12:4000/exercises", {
          headers: {},
          withCredentials: true,
        });
        setItems(response.data);
      } catch (err) {
        if (err.response?.status === 401) {
          localStorage.removeItem("token");
        } else {
          setError("Error loading exercises");
        }
      }
    };

    fetchExercises();
  }, [navigate]);

  const toggleExpand = (id) => {
    setExpandedItems((prev) => ({ ...prev, [id]: !prev[id] }));
    if (!personalBests[id]) {
      setPersonalBests((prev) => ({
        ...prev,
        [id]: personalBestData[id] || "-",
      }));
    }
    if (!setsData[id]) {
      setSetsData((prev) => ({
        ...prev,
        [id]: [{ setno: 1, weights: 0, repetitions: 0 }],
      }));
    }
  };
  const handleLogout = async (e) => {
    e.preventDefault();
    try {
      await fetch("/logout", {
        method: "POST",
        credentials: "include",
      });
    } catch (err) {
      console.error("Logout failed:", err);
    }
    navigate("/login");
  };

  const addSet = (id) => {
    setSetsData((prev) => {
      const arr = prev[id] || [];
      return {
        ...prev,
        [id]: [...arr, { setno: arr.length + 1, weights: 0, repetitions: 0 }],
      };
    });
  };

  const handleInputChange = (id, idx, field, value) => {
    setSetsData((prev) => {
      const arr = [...(prev[id] || [])];
      arr[idx] = { ...arr[idx], [field]: Number(value) };
      return { ...prev, [id]: arr };
    });
  };

  const handleAddWorkout = async () => {
    setError("");
    const token = localStorage.getItem("token");
    if (!token) {
      setError("Not authorized. Please log in.");
      navigate("/login", { replace: true });
      return;
    }
    console.log("all exercise IDs in setsData:", Object.keys(setsData));
    const payload = {
      workouts: Object.entries(setsData).map(([exerciseid, sets]) => ({
        exerciseid: Number(exerciseid),
        sets: sets.map(({ setno, weights, repetitions }) => ({
          setno,
          weights,
          repetitions,
        })),
        created_at: new Date().toISOString(),
      })),
    };

    console.log(payload);
    console.log("→ Posting to /add-workout with headers/payload:", {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: payload,
    });

    setLoading(true);
    try {
      await axios.post("http://192.168.0.12:4000/add-workout", payload, {
        headers: {
          "Content-Type": "application/json",
          "session-token": `${token}`,
        },
        withCredentials: true,
      });
      navigate("/dashboard", { replace: true });
    } catch (err) {
      if (err.response?.status === 401) {
        localStorage.removeItem("token");
        setError("Session expired. Please log in again.");
      } else {
        setError(`Error adding workout: ${err.message}`);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="dashcontainer">
      <div className="topnav">
        <a href="/dashboard">Home</a>
        <a className="active" href="#">
          Add Workout
        </a>
        <a href="/profile" style={{ float: "right" }}>
          Profile
        </a>
        <a href="/login" onClick={handleLogout} style={{ float: "right" }}>
          Logout
        </a>
      </div>

      <div className="Logo">
        <h1 className="Page-Heading">Select Exercise</h1>
      </div>

      <div className="addContainer">
        <div className="selectExercise">
          <h2 className="exeListHead">Exercise List</h2>
          {error && <p className="errorMsg">{error}</p>}

          {items.map((item, idx) => {
            const id = item.exercise_id ?? item.id ?? idx;
            const name =
              item.name ?? item.title ?? item.exerciseName ?? `#${id}`;

            return (
              <div key={id}>
                <div
                  className="workouts"
                  onClick={() => toggleExpand(id)}
                  style={{ cursor: "pointer" }}
                >
                  {name}
                </div>

                {expandedItems[id] && (
                  <div className="expandedOptions">
                    <p className="pBest">
                      Personal Best: {personalBests[id] || "Loading..."}
                    </p>
                    <div className="workout-titles">
                      <p>Set</p>
                      <p>KG/s</p>
                      <p>Reps</p>
                    </div>
                    {setsData[id]?.map((set, sidx) => (
                      <div className="workout" key={sidx}>
                        <input
                          className="metrics"
                          type="number"
                          value={set.setno}
                          onChange={(e) =>
                            handleInputChange(id, sidx, "setno", e.target.value)
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
                              sidx,
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
                              sidx,
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

          <button
            className="addButton"
            onClick={handleAddWorkout}
            disabled={loading}
          >
            {loading ? "Saving..." : "Add Workout"}
          </button>
        </div>
      </div>
    </div>
  );
};

export default AddExercise;
