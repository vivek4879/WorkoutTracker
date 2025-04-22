import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "../Dashboard/Dashboard.css";          // yellow / black navbar & pale‑blue body

function MeasurementsForm() {
  const [measurements, setMeasurements] = useState({});
  const navigate = useNavigate();

  /* ------------- handlers ------------- */
  const handleChange = (e) =>
    setMeasurements({ ...measurements, [e.target.name]: e.target.value });

  const handleSubmit = () => {
    localStorage.setItem("userMeasurements", JSON.stringify(measurements));
    /* go to the table page */
    navigate("/measurements/display");
  };

  /* ------------- jsx ------------- */
  return (
    <div className="container">
      {/* -------- navbar (copied from dashboard) -------- */}
      <div
        className="topnav"
        /* inline style forces the bar to stay yellow even if another rule overrides */
        style={{ backgroundColor: "#ffeb00" }}
      >
        <a href="/dashboard" className="active">
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
        <a href="/addexercise">Add Workout</a>

        <Link style={{ float: "right" }} to="/login">
          Logout
        </Link>
        <Link style={{ float: "right" }} to="/profile">
          Profile
        </Link>
        <Link style={{ float: "right" }} to="/measurements">
          Measurements
        </Link>
      </div>

      {/* -------- headings -------- */}
      <div className="Logo">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>
      <div className="header">

      {/* 3) A subheading or page title, if you want it */}
      <div className="Header">
        <h2 className="Page-Heading">Enter Measurements</h2>
      </div>

      {/* -------- white card -------- */}
      <div
        style={{
          width: "90%",
          maxWidth: "500px",
          margin: "0 auto 40px auto",
          background: "#fff",
          borderRadius: "8px",
          boxShadow: "0 4px 8px rgba(0,0,0,0.1)",
          padding: "20px",
        }}
      >
        <div style={{ display: "flex", flexDirection: "column", gap: "10px" }}>
          {[
            ["weight", "Weight (kg)"],
            ["neck", "Neck (cm)"],
            ["shoulders", "Shoulders (cm)"],
            ["chest", "Chest (cm)"],
            ["leftBicep", "Left Bicep (cm)"],
            ["rightBicep", "Right Bicep (cm)"],
            ["upperAbs", "Upper Abs (cm)"],
            ["lowerAbs", "Lower Abs (cm)"],
            ["waist", "Waist (cm)"],
            ["hips", "Hips (cm)"],
            ["leftThigh", "Left Thigh (cm)"],
            ["rightThigh", "Right Thigh (cm)"],
            ["leftCalf", "Left Calf (cm)"],
            ["rightCalf", "Right Calf (cm)"],
          ].map(([name, placeholder]) => (
            <input
              key={name}
              type="number"
              name={name}
              placeholder={placeholder}
              onChange={handleChange}
              style={{ padding: "8px", borderRadius: "4px" }}
            />
          ))}
        </div>

        <button
          onClick={handleSubmit}
          style={{
            marginTop: "20px",
            padding: "10px 15px",
            background: "#323232",
            color: "#fff",
            borderRadius: "4px",
            border: "none",
            cursor: "pointer",
          }}
        >
          Update
        </button>
      </div>
    </div>
  );
}

export default MeasurementsForm;
