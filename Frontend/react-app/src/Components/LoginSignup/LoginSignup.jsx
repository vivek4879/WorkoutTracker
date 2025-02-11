import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./LoginSignup.css";

const LoginSignup = () => {
  const [action, setAction] = useState("Login");
  const navigate = useNavigate();

  const handleLogin = () => {
    // Normally, you'd validate credentials before navigating
    navigate("/dashboard");
  };

  return (
    <div className="LoginSignup">
      <div className="total-box">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>
      <div className="header">
        <div className="text">{action}</div>
      </div>
      <div className="inputs">
        {action === "Login" ? (
          <div></div>
        ) : (
          <>
            <div className="input">
              <input type="text" placeholder="First Name" required />
            </div>
            <div className="input">
              <input type="text" placeholder="Last Name" required />
            </div>
          </>
        )}
        <div className="input">
          <input type="email" placeholder="Email" required />
        </div>
        <div className="input">
          <input type="password" placeholder="Password" required />
        </div>
      </div>
      {action === "Login" ? (
        <div className="forgot-password">
          <Link to="/forgot-password">Forgot Password?</Link>
        </div>
      ) : null}

      <div className="submit-container">
        <div
          className={
            action === "Sign Up" ? "submit-button gray" : "submit-button"
          }
          onClick={() => setAction("Sign Up")}
        >
          Sign Up
        </div>
        {action === "Login" && (
          <div className="submit-button" onClick={handleLogin}>
            Login
          </div>
        )}
      </div>
    </div>
  );
};

export default LoginSignup;
