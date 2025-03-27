import "./Profile.css";

const Profile = () => {
  return (
    <div className="container">
      <div className="container">
        <div className="topnav">
          <a  href="/dashboard">
            Home
          </a>
          <a href="#">Add Workout</a>
          <a className="active" style={{ float: "right" }} href="#">
            Profile
          </a>
        </div>
        <div className="header">
          <h2 className="Page-Heading">User Profile!</h2>
        </div>
        <div className="profile-container">
      <div className="profile-card">
        <div className="profile-header">
          <div className="profile-image-container">
            <img src="https://fortune.com/img-assets/wp-content/uploads/2024/08/GettyImages-1253261584-e1723845288674.jpg?w=1440&q=75" alt="Profile" className="profile-image" />
            <button className="edit-profile-btn">Edit</button>
          </div>
          <div className="profile-info">
            <h2 className="profile-name">John Doe</h2>
           
            <p className="profile-details">johndoe@example.com</p>
            <p className="profile-details">Age - 26 Years</p>
            <p className="profile-details">Gender-Male</p>
            
          </div>
        </div>
      </div>
    </div>
      </div>
    </div>
  );
};

export default Profile;
