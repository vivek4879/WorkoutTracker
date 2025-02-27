import { Link, useNavigate } from "react-router-dom";
import "./Login.css";
const Login = () => {
  const navigate = useNavigate();

  const handleLogin = () => {
    // Normally, you'd validate credentials before navigating
    navigate("/dashboard");
  };

  const handleSignUp = () => {
    navigate("/signup");
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
            <input type="email" placeholder="Email" required />
          </div>
          <div className="input">
            <input type="password" placeholder="Password" required />
          </div>
        </div>

        <div className="forgot-password">
          <Link to="/forgot-password">Forgot Password?</Link>
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
