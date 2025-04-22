import React from "react";
import { Link } from "react-router-dom";
import "../Dashboard/Dashboard.css";            // yellow / black navbar

function MeasurementsDisplay() {
  const measurements =
    JSON.parse(localStorage.getItem("userMeasurements")) || {};

  /* --------------------------------------------- */
  /* Navbar is identical to Dashboard.jsx          */
  /* --------------------------------------------- */
  const NavBar = () => (
    <div className="topnav">
      <a className="active" href="/dashboard">
        Home
      </a>
      <a href="/addexercise">Add Workout</a>

      {/* right‑aligned links */}
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
  );

  /* If no data found */
  if (Object.keys(measurements).length === 0) {
    return (
      <div className="container">
        <NavBar />

        <div className="Logo">
          <h1 className="Page-Heading">Gambare!</h1>
        </div>
        <div className="header">
          <h2 className="Page-Heading">Your Measurements</h2>
        </div>

        <div
          style={{
            width: "90%",
            maxWidth: "500px",
            margin: "0 auto 40px auto",
            background: "#fff",
            borderRadius: "8px",
            boxShadow: "0 4px 8px rgba(0,0,0,0.1)",
            padding: "20px",
            textAlign: "center",
            color: "#333",
          }}
        >
          No measurements found. Please enter them first.
        </div>
      </div>
    );
  }

  /* Show table of saved measurements */
  return (
    <div className="container">
      <NavBar />

      <div className="Logo">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>
      <div className="header">
        <h2 className="Page-Heading">Your Measurements</h2>
      </div>

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
        <table
          style={{
            width: "100%",
            borderCollapse: "collapse",
            color: "#333",
          }}
        >
          <thead>
            <tr style={{ background: "#f5f5f5" }}>
              <th
                style={{
                  textAlign: "left",
                  padding: "8px",
                  borderBottom: "1px solid #ddd",
                }}
              >
                Measurement
              </th>
              <th
                style={{
                  textAlign: "left",
                  padding: "8px",
                  borderBottom: "1px solid #ddd",
                }}
              >
                Value
              </th>
            </tr>
          </thead>
          <tbody>
            {Object.entries(measurements).map(([key, value]) => {
              /* Convert camelCase → Title Case (Right Bicep, Upper Abs, …) */
              const label =
                key.charAt(0).toUpperCase() +
                key
                  .slice(1)
                  .replace(/([A-Z])/g, " $1")
                  .replace("Upper Abs", "Upper Abs")
                  .replace("Lower Abs", "Lower Abs");
              return (
                <tr key={key}>
                  <td
                    style={{
                      padding: "8px",
                      borderBottom: "1px solid #ddd",
                    }}
                  >
                    {label}
                  </td>
                  <td
                    style={{
                      padding: "8px",
                      borderBottom: "1px solid #ddd",
                    }}
                  >
                    {value}
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default MeasurementsDisplay;
