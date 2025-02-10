import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./ForgotPass.css";

const ForgotPass = () => {
  const [email, setEmail] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Password reset link sent to:", email);
    alert("If this email is registered, you'll receive a password reset link.");
    navigate("/login"); // Redirect back to login page
  };

  return (
    <div className="LoginSignup">
      <div className="total-box">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>
      <div className="header">
        <div className="text">Forgot Password</div>
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
        <button className="submit-button" onClick={handleSubmit}>
          Submit
        </button>
      </div>
    </div>
  );
};

export default ForgotPass;
