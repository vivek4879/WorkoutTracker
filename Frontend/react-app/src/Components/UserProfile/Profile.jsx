import "./Profile.css";

const Profile = () => {
  return (
    <div className="container">
      <div className="container">
        <div className="topnav">
          <a href="/dashboard">Home</a>
          <a href="/profile" className="active" style={{ float: "right" }}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              height="23px"
              viewBox="0 -960 960 960"
              width="23px"
              fill="#000000"
            >
              <path d="M235.2-279.59q51-38.52 113.52-60.54 62.52-22.02 131.33-22.02t131.64 22.38q62.83 22.38 113.11 60.42 34.29-41.24 53.31-91.92 19.02-50.69 19.02-108.73 0-131.57-92.78-224.35T480-797.13q-131.57 0-224.35 92.78T162.87-480q0 57.8 18.9 108.49 18.9 50.68 53.43 91.92ZM480-437.85q-59.96 0-100.93-40.86-40.98-40.86-40.98-100.81 0-59.96 40.98-100.94 40.97-40.97 100.93-40.97 59.96 0 100.93 40.97 40.98 40.98 40.98 100.94 0 59.95-40.98 100.81-40.97 40.86-100.93 40.86Zm-.02 365.98q-84.65 0-159.09-32.1-74.43-32.1-129.63-87.29-55.19-55.2-87.29-129.65-32.1-74.46-32.1-159.11 0-84.65 32.1-159.09 32.1-74.43 87.29-129.63 55.2-55.19 129.65-87.29 74.46-32.1 159.11-32.1 84.65 0 159.09 32.1 74.43 32.1 129.63 87.29 55.19 55.2 87.29 129.65 32.1 74.46 32.1 159.11 0 84.65-32.1 159.09-32.1 74.43-87.29 129.63-55.2 55.19-129.65 87.29-74.46 32.1-159.11 32.1Zm.02-91q51.8 0 97.37-14.9 45.56-14.9 84.56-42.95-39.47-28.28-84.44-43.06-44.97-14.79-97.49-14.79-52.52 0-97.37 14.79-44.85 14.78-84.33 43.06 39 28.05 84.45 42.95 45.45 14.9 97.25 14.9Zm0-358.56q25.04 0 41.57-16.53 16.52-16.52 16.52-41.56 0-25.05-16.52-41.69-16.53-16.64-41.57-16.64t-41.57 16.64q-16.52 16.64-16.52 41.69 0 25.04 16.52 41.56 16.53 16.53 41.57 16.53Zm0-58.09Zm.24 358.8Z" />
            </svg>
          </a>
        </div>
        <div className="header">
          <h2 className="Page-Heading">User Profile!</h2>
        </div>
        <div className="profile-container">
          <div className="profile-card">
            <div className="profile-header">
              <div className="profile-image-container">
                <img
                  src="https://fortune.com/img-assets/wp-content/uploads/2024/08/GettyImages-1253261584-e1723845288674.jpg?w=1440&q=75"
                  alt="Profile"
                  className="profile-image"
                />
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
