import "./App.css";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import Login from "./Components/Login/Login.jsx";
import SignUp from "./Components/SignUp/SignUp.jsx";
import ForgotPass from "./Components/ForgotPass/ForgotPass.jsx";
import Dashboard from "./Components/Dashboard/Dashboard.jsx";
// import SignUp from "./Components/LoginSignup/Signup.jsx";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/forgot-password" element={<ForgotPass />} />
        <Route path="/dashboard" element={<Dashboard />} />
        {/* Redirect the root path to the login page */}
        <Route path="/" element={<Navigate to="/login" replace />} />
        {/* Fallback route */}
        <Route path="*" element={<Login />} />
      </Routes>
    </Router>
  );
}

export default App;
