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
import Profile from "./Components/UserProfile/Profile.jsx";
import UserGoals from "./Components/UserGoals/usergoals.jsx";



import MeasurementsForm from "./Components/Measurements/MeasurementsForm.jsx";
import MeasurementsDisplay from "./Components/Measurements/MeasurementsDisplay.jsx";

import AddExercise from "./Components/AddExercise/AddExercise.jsx";
//  import Test from "./Components/test/test.jsx";

function App() {
  return (
    <div>
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/forgot-password" element={<ForgotPass />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/usergoals" element={<UserGoals />} />
          
      

          <Route path="/measurements" element={<MeasurementsForm />} />
          <Route path="/measurements/display" element={<MeasurementsDisplay />} />
          

          <Route path="/addexercise" element={<AddExercise />} />
          {/* <Route path="/test" element={<Test />} /> */}
          {/* Redirect the root path to the login page */}
          <Route path="/" element={<Navigate to="/login" replace />} />

          {/* Fallback route */}
          <Route path="*" element={<Login />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
