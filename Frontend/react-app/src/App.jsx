import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginSignup from "./Components/LoginSignup/LoginSignup.jsx";
import ForgotPass from "./Components/ForgotPass/ForgotPass.jsx";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<LoginSignup />} />
        <Route path="/forgot-password" element={<ForgotPass />} />
        <Route path="*" element={<LoginSignup />} /> {/* Default route */}
      </Routes>
    </Router>
  );
}

export default App;
