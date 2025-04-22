import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./ForgotPass.css";

const ForgotPass = () => {
  const [email, setEmail] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    if (email.length <= 0) {
      alert("Please enter a valid email");
      return;
    }
    console.log("Password reset link sent to:", email);
    alert("If this email is registered, you'll receive a password reset link.");
    navigate("/login"); // Redirect back to login page
  };

  return (
    <div className="container">
      <div className="Logo">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>
      <div className="ForgotPass">
        <div className="Header">
          <h2 className="Page-Heading">Forgot Password?</h2>
        </div>
        <div className="inputs">
          <div className="input">
            <input
              type="email"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
        </div>

        <div className="submit-container">
          <div className="submit-button" onClick={handleSubmit}>
            Submit
          </div>
        </div>
      </div>
    </div>
  );
};

export default ForgotPass;
