import { useState } from "react";
import { Link } from "react-router-dom";

import "../Dashboard/Dashboard.css";
  // yellow / black navbar styles
import "./Profile.css";     // page‑specific styles

const Profile = () => {
  // simple local state (replace with real data source later)
  const [form, setForm] = useState({
    firstName: "John",
    lastName: "Doe",
    email: "johndoe@example.com",
    password: "",
  });

  const handleChange = (e) =>
    setForm({ ...form, [e.target.name]: e.target.value });

  const handleSave = () => {
    // stub – replace with API call / context update
    console.log("Profile saved", form);
    alert("Profile information updated.");
  };

  return (
    <div className="container">
      <div className="container">
        {/* ───────── navbar identical to dashboard ───────── */}
        <div className="topnav">
          <a href="/dashboard">Home</a>
          <a href="/addexercise">Add Workout</a>

          {/* right‑aligned links */}
          <a href="/profile" style={{ float: "right" }}>
            {/* same SVG user icon */}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 -960 960 960"
              width="23"
              height="23"
              fill="#000000"
            >
              <path d="M235.2-279.59q51-38.52 113.52-60.54 62.52-22.02 131.33-22.02t131.64 22.38q62.83 22.38 113.11 60.42 34.29-41.24 53.31-91.92 19.02-50.69 19.02-108.73 0-131.57-92.78-224.35T480-797.13q-131.57 0-224.35 92.78T162.87-480q0 57.8 18.9 108.49 18.9 50.68 53.43 91.92ZM480-437.85q-59.96 0-100.93-40.86-40.98-40.86-40.98-100.81 0-59.96 40.98-100.94 40.97-40.97 100.93-40.97 59.96 0 100.93 40.97 40.98 40.98 40.98 100.94 0 59.95-40.98 100.81-40.97 40.86-100.93 40.86Zm-.02 365.98q-84.65 0-159.09-32.1-74.43-32.1-129.63-87.29-55.19-55.2-87.29-129.65-32.1-74.46-32.1-159.11 0-84.65 32.1-159.09 32.1-74.43 87.29-129.63 55.2-55.19 129.65-87.29 74.46-32.1 159.11-32.1 84.65 0 159.09 32.1 74.43 32.1 129.63 87.29 55.19 55.2 87.29 129.65 32.1 74.46 32.1 159.11 0 84.65-32.1 159.09-32.1 74.43-87.29 129.63-55.2 55.19-129.65 87.29-74.46 32.1-159.11 32.1Z" />
            </svg>
          </a>
          <Link style={{ float: "right" }} to="/measurements">
            Measurements
          </Link>
          <Link style={{ float: "right" }} to="/login">
            Logout
          </Link>
        </div>
        {/* ───────── end navbar ───────── */}

        <div className="header">
          <h2 className="Page-Heading">Edit Profile</h2>
        </div>

        {/* ───────── profile edit card ───────── */}
        <div className="profile-container">
          <div className="profile-card">
            <form
              className="profile-form"
              onSubmit={(e) => {
                e.preventDefault();
                handleSave();
              }}
            >
              <label>
                First Name
                <input
                  type="text"
                  name="firstName"
                  value={form.firstName}
                  onChange={handleChange}
                  required
                />
              </label>

              <label>
                Last Name
                <input
                  type="text"
                  name="lastName"
                  value={form.lastName}
                  onChange={handleChange}
                  required
                />
              </label>

              <label>
                Email
                <input
                  type="email"
                  name="email"
                  value={form.email}
                  onChange={handleChange}
                  required
                />
              </label>

              <label>
                Password
                <input
                  type="password"
                  name="password"
                  value={form.password}
                  onChange={handleChange}
                  required
                />
              </label>

              <button type="submit" className="save-btn">
                Save
              </button>
            </form>
          </div>
        </div>
        {/* ───────── end card ───────── */}
      </div>
    </div>
  );
};

export default Profile;
