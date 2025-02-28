import "./Calendar.css";

const Calendar = () => {
  return (
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
          <div className="calBox">
            <p>1</p>
          </div>
          <div className="calBox">
            <p>2</p>
          </div>
          <div className="calBox">
            <p>3</p>
          </div>
          <div className="calBox">
            <p>4</p>
          </div>
          <div className="calBox">
            <p>5</p>
          </div>
          <div className="calBox">
            <p>6</p>
          </div>
          <div className="calBox">
            <p>7</p>
          </div>
        </div>
        <div className="week">
          <div className="calBox">
            <p>8</p>
          </div>
          <div className="calBox">
            <p>9</p>
          </div>
          <div className="calBox">
            <p>10</p>
          </div>
          <div className="calBox">
            <p>11</p>
          </div>
          <div className="calBox">
            <p>12</p>
          </div>
          <div className="calBox">
            <p>13</p>
          </div>
          <div className="calBox">
            <p>14</p>
          </div>
        </div>
        <div className="week">
          <div className="calBox">
            <p>15</p>
          </div>
          <div className="calBox">
            <p>16</p>
          </div>
          <div className="calBox">
            <p>17</p>
          </div>
          <div className="calBox">
            <p>18</p>
          </div>
          <div className="calBox">
            <p>19</p>
          </div>
          <div className="calBox">
            <p>20</p>
          </div>
          <div className="calBox">
            <p>21</p>
          </div>
        </div>
        <div className="week">
          <div className="calBox">
            <p>22</p>
          </div>
          <div className="calBox">
            <p>23</p>
          </div>
          <div className="calBox">
            <p>24</p>
          </div>
          <div className="calBox">
            <p>25</p>
          </div>
          <div className="calBox">
            <p>26</p>
          </div>
          <div className="calBox">
            <p>27</p>
          </div>
          <div className="calBox">
            <p>28</p>
          </div>
        </div>
        <div className="week">
          <div className="calBox">
            <p>29</p>
          </div>
          <div className="calBox">
            <p>30</p>
          </div>
          <div className="calBox">
            <p>31</p>
          </div>
          <div className="calBox">
            <p></p>
          </div>
          <div className="calBox">
            <p></p>
          </div>
          <div className="calBox">
            <p></p>
          </div>
          <div className="calBox">
            <p></p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Calendar;
