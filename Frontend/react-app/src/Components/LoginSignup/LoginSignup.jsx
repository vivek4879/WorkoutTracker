import { useState } from "react";
import { Link } from "react-router-dom";
import "./LoginSignup.css";

const LoginSignup = () => {
  const [action, setAction] = useState("Login");
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
            {" "}
            <div className="input">
              <input type="text" placeholder="First Name" required="true" />
            </div>
            <div className="input">
              <input type="text" placeholder="Last Name" required="true" />
            </div>
          </>
        )}
        <div className="input">
          <input type="email" placeholder="Email" required="true" />
        </div>
        <div className="input">
          <input type="password" placeholder="Password" required="true" />
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
            action === "Login" ? "submit-button gray" : "submit-button"
          }
          onClick={() => {
            setAction("Sign Up");
          }}
        >
          Sign Up
        </div>
        <div
          className={
            action === "Sign Up" ? "submit-button gray" : "submit-button"
          }
          onClick={() => {
            setAction("Login");
          }}
        >
          Login
        </div>
      </div>
    </div>
  );
};

export default LoginSignup;
