import { Link, useNavigate } from "react-router-dom";
import "./Login.css";
import axios from "axios";
import { useState } from "react";

const Login = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [token, setToken] = useState("");
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  const handleLogin = async () => {
    // navigate("/dashboard");
    if (error === "") {
      try {
        const response = await axios.post(
          "http://192.168.0.12:4000/authenticate",
          {
            email,
            password,
          },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );
        // setToken(response.data.token); // Save JWT token
        localStorage.setItem("token", response.data.token); // Store in localStorage
        alert("Login successful");
      } catch (error) {
        alert(error.response?.data || error);
      }
    }
  };
  const handleSignUp = () => {
    navigate("/signup");
  };
  const handleForgotPass = () => {
    navigate("/forgot-password");
  };
  const handleEmailChange = (e) => {
    const value = e.target.value;
    setEmail(value);

    // Validate email
    if (!emailRegex.test(value)) {
      setError("Please enter a valid email address.");
    } else {
      setError("");
    }
  };
  const handleChangePassword = (e) => {
    const value = e.target.value;
    setPassword(value);
  };
  return (
    <div className="container">
      <div className="Logo">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>

      <div className="Login">
        <div className="header">
          <h2 className="Page-Heading">Login</h2>
        </div>
        <div className="inputs">
          <div className="input">
            <input
              type="email"
              value={email}
              onChange={handleEmailChange}
              placeholder="Email"
              required
            />
          </div>
          <div>{error && <p style={{ color: "red" }}>{error}</p>}</div>
          <div className="input">
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={handleChangePassword}
              required
            />
          </div>
        </div>
        <div>{token && <p style={{ color: "red" }}>{token}</p>}</div>
        <div className="forgot-password" onClick={handleForgotPass}>
          <div className="header">
            <Link>Forgot Password?</Link>
          </div>
        </div>

        <div className="submit-container">
          <div className="submit-button" onClick={handleLogin}>
            Login
          </div>
          <div className="submit-button gray" onClick={handleSignUp}>
            Sign Up
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
