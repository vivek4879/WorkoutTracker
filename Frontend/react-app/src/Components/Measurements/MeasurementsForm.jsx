
// src/Components/Measurements/MeasurementsForm.jsx
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
// Import the same CSS your Dashboard uses, so the yellow line & styles match
import "../Dashboard/Dashboard.css"; // Adjust the relative path as needed

function MeasurementsForm() {
  const [measurements, setMeasurements] = useState({});
  const navigate = useNavigate();

  const handleChange = (e) => {
    setMeasurements({
      ...measurements,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = () => {
    localStorage.setItem("userMeasurements", JSON.stringify(measurements));
    navigate("/measurements/display");
  };

  return (
    // Container class to get the same pale background color
    <div className="dashcontainer">
      {/* 
        1) Same topnav bar used by your Dashboard (with black left area & yellow).
           This gives you the "yellow line" across the top.
      */}
      <div className="topnav">
        <a className="active" href="#">
          Home
        </a>
        <a href="#">Add Workout</a>
        {/* Right side links (float: right) */}
        <Link style={{ float: "right" }} to="/login">
          Logout
        </Link>
        <Link style={{ float: "right" }} to="/profile">
          Profile
        </Link>
        <Link style={{ float: "right" }} to="/usergoals">
          UserGoal
        </Link>
        <Link style={{ float: "right" }} to="/measurements">
          Measurements
        </Link>
      </div>

      {/* 
        2) The "Gambare!" heading (like your home page).
        This appears under the yellow line, matching your Dashboard style.
      */}
      <div className="Logo">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>

      {/* 3) A subheading or page title, if you want it */}
      <div className="Header">
        <h2 className="Page-Heading">Enter Measurements</h2>
      </div>

      {/* 
        4) The white box in the center for your measurements inputs.
        We’re using inline styles for the box—feel free to move them to CSS.
      */}
      <div
        style={{
          width: "90%",
          maxWidth: "500px",
          margin: "0 auto 40px auto",
          backgroundColor: "#fff",
          borderRadius: "8px",
          boxShadow: "0 4px 8px rgba(0,0,0,0.1)",
          padding: "20px",
        }}
      >
        {/* The input fields, each with a bit of styling */}
        <div style={{ display: "flex", flexDirection: "column", gap: "10px" }}>
          <input
            type="number"
            name="weight"
            placeholder="Weight (kg)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="neck"
            placeholder="Neck (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="shoulders"
            placeholder="Shoulders (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="chest"
            placeholder="Chest (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="leftBicep"
            placeholder="Left Bicep (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="rightBicep"
            placeholder="Right Bicep (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="upperAbs"
            placeholder="Upper Abs (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="lowerAbs"
            placeholder="Lower Abs (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="waist"
            placeholder="Waist (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="hips"
            placeholder="Hips (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="leftThigh"
            placeholder="Left Thigh (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="rightThigh"
            placeholder="Right Thigh (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="leftCalf"
            placeholder="Left Calf (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
          <input
            type="number"
            name="rightCalf"
            placeholder="Right Calf (cm)"
            onChange={handleChange}
            style={{ padding: "8px", borderRadius: "4px" }}
          />
        </div>

        <button
          onClick={handleSubmit}
          style={{
            marginTop: "20px",
            padding: "10px 15px",
            backgroundColor: "#333",
            color: "#fff",
            borderRadius: "4px",
            cursor: "pointer",
            border: "none",
          }}
        >
          Update
        </button>
      </div>
    </div>
  );
}

export default MeasurementsForm;
