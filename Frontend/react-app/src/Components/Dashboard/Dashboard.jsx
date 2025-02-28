import "./Dashboard.css";

const Dashboard = () => {
  return (
    <div className="container">
      <div className="container">
        <div className="topnav">
          <a className="active" href="#">
            Home
          </a>
          <a href="#">Add Workout</a>
          <a style={{ float: "right" }} href="#">
            Profile
          </a>
        </div>
        <div className="header">
          <h2 className="Page-Heading">Your Streak!</h2>
        </div>
        <div className="calContainer">
          <div className="calendar">
            <div className="week">
              <div className="calBox dayBox">Mon</div>
              <div className="calBox dayBox">Tue</div>
              <div className="calBox dayBox">Wed</div>
              <div className="calBox dayBox">Thu</div>
              <div className="calBox dayBox">Fri</div>
              <div className="calBox dayBox">Sat</div>
              <div className="calBox dayBox today">Sun</div>
            </div>
            <div className="week">
              <div className="calBox">1</div>
              <div className="calBox">2</div>
              <div className="calBox">3</div>
              <div className="calBox">4</div>
              <div className="calBox">5</div>
              <div className="calBox">6</div>
              <div className="calBox">7</div>
            </div>
            <div className="week">
              <div className="calBox">8</div>
              <div className="calBox">9</div>
              <div className="calBox">10</div>
              <div className="calBox">11</div>
              <div className="calBox">12</div>
              <div className="calBox">13</div>
              <div className="calBox">14</div>
            </div>
            <div className="week">
              <div className="calBox">15</div>
              <div className="calBox">16</div>
              <div className="calBox">17</div>
              <div className="calBox">18</div>
              <div className="calBox">19</div>
              <div className="calBox">20</div>
              <div className="calBox">21</div>
            </div>
            <div className="week">
              <div className="calBox">22</div>
              <div className="calBox">23</div>
              <div className="calBox">24</div>
              <div className="calBox">25</div>
              <div className="calBox">26</div>
              <div className="calBox today">27</div>
              <div className="calBox">28</div>
            </div>
            <div className="week">
              <div className="calBox">29</div>
              <div className="calBox">30</div>
              <div className="calBox">31</div>
              <div className="calBox"></div>
              <div className="calBox"></div>
              <div className="calBox"></div>
              <div className="calBox"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
