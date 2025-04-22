import { useNavigate } from "react-router-dom";
import { useState } from "react";
import "./SignUp.css";
const SignUp = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [emailError, setError] = useState("");
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  const navigate = useNavigate();
  const passwordRegex =
    /^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*\W)(?!.* ).{8,16}$/;
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

    // Validate email
    if (!passwordRegex.test(value)) {
      setError("Please enter a valid password.");
    } else {
      setError("");
    }
  };
  const handleSignUp = () => {
    // Normally, you'd validate the inputs and create an account before navigating
    navigate("/login");
  };
  // const handleLogin = () => {
  //   // Normally, you'd validate credentials before navigating
  //   navigate("/login");
  // };

  return (
    <div className="container">
      <div className="Logo">
        <h1 className="Page-Heading">Gambare!</h1>
      </div>
      <div className="SignUp">
        <div className="Header">
          <h2 className="Page-Heading">Sign Up</h2>
        </div>
        <div className="inputs">
          <div className="input">
            <input type="text" placeholder="First Name" required />
          </div>
          <div></div>
          <div className="input">
            <input type="text" placeholder="Last Name" required />
          </div>
          <div></div>
          <div className="input">
            <input
              type="email"
              value={email}
              onChange={handleEmailChange}
              placeholder="Email"
              required
            />
          </div>
          <div>
            {emailError && <p style={{ color: "red" }}>{emailError}</p>}
          </div>
          <div className="input">
            <input
              type="password"
              value={password}
              onChange={handleChangePassword}
              placeholder="Password"
              required
            />
          </div>
          <div></div>
        </div>
        <div className="submit-container">
          <div className="submit-button" onClick={handleSignUp}>
            Sign Up
          </div>
          {/* <div className="submit-button gray" onClick={handleLogin}>
            Login
          </div> */}
        </div>
      </div>
    </div>
  );
};

export default SignUp;
