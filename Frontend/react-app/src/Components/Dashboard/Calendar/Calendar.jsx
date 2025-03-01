import "./Calendar.css";
function Calendar() {
  const date = new Date();
  const days = date.getDate();
  const newDate = new Date(date.getTime() - days * 24 * 60 * 60 * 1000);
  var offset = newDate.getDay();
  const daysOfWeek = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];
  const week1 = [1, 2, 3, 4, 5, 6, 7];
  return (
    <div className="calContainer">
      <div className="calendar">
        <div className="week">
          {daysOfWeek.map((day) => (
            <div className="calBox dayBox" key={day}>
              <div>
                <p>{day}</p>
              </div>
            </div>
          ))}
        </div>
        <div className="week">
          {week1.map((num) => (
            <div className="calBox" key={num}>
              {num - offset > 0 && (
                <div>
                  <p>{num - offset}</p>
                </div>
              )}
            </div>
          ))}
        </div>
        <div className="week">
          {week1.map((num) => (
            <div className="calBox" key={num + 7}>
              {num + 7 - offset > 0 && (
                <div>
                  <p>{num + 7 - offset}</p>
                </div>
              )}
            </div>
          ))}
        </div>
        <div className="week">
          {week1.map((num) => (
            <div className="calBox" key={num + 14}>
              {num + 14 - offset > 0 && (
                <div>
                  <p>{num + 14 - offset}</p>
                </div>
              )}
            </div>
          ))}
        </div>
        <div className="week">
          <div className="calBox">
            {22 - offset > 0 && (
              <div>
                <p>{22 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {23 - offset > 0 && (
              <div>
                <p>{23 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {24 - offset > 0 && (
              <div>
                <p>{24 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {25 - offset > 0 && (
              <div>
                <p>{25 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {26 - offset > 0 && (
              <div>
                <p>{26 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {27 - offset > 0 && (
              <div>
                <p>{27 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {28 - offset > 0 && (
              <div>
                <p>{28 - offset}</p>
              </div>
            )}
          </div>
        </div>
        <div className="week">
          <div className="calBox">
            {29 - offset > 0 && (
              <div>
                <p>{29 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {30 - offset > 0 && (
              <div>
                <p>{30 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {31 - offset > 0 && (
              <div>
                <p>{31 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {32 - offset > 0 && (
              <div>
                <p>{32 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {33 - offset > 0 && (
              <div>
                <p>{33 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {34 - offset > 0 && (
              <div>
                <p>{34 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {35 - offset > 0 && (
              <div>
                <p>{35 - offset}</p>
              </div>
            )}
          </div>
        </div>
        <div className="week">
          <div className="calBox">
            {36 - offset > 0 && (
              <div>
                <p>{36 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {37 - offset > 0 && (
              <div>
                <p>{37 - offset}</p>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default Calendar;
