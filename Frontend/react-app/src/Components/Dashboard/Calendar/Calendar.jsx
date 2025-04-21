import "./Calendar.css";
function Calendar() {
  const date = new Date();
  const days = date.getDate();
  const newDate = new Date(date.getTime() - days * 24 * 60 * 60 * 1000);
  var offset = newDate.getDay();
  const daysOfWeek = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];
  const week1 = [...Array(7).keys()].map((x) => x + 1);
  // console.log(squares);
  // const week1 = [1, 2, 3, 4, 5, 6, 7];
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
          {week1.map((num) => (
            <div className="calBox" key={num + 21}>
              {num + 21 - offset > 0 && (
                <div>
                  <p>{num + 21 - offset}</p>
                </div>
              )}
            </div>
          ))}
        </div>
        <div className="week">
          {week1.map((num) => (
            <div className="calBox streak" key={num + 28}>
              {31 >= num + 28 - offset > 0 && (
                <div>
                  <p>{num + 28 - offset}</p>
                </div>
              )}
            </div>
          ))}
        </div>
        <div className="week">
          <div className="calBox today">
            {31 >= 36 - offset > 0 && (
              <div>
                <p>{36 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox">
            {31 >= 37 - offset > 0 && (
              <div>
                <p>{37 - offset}</p>
              </div>
            )}
          </div>
          <div className="calBox"></div>
          <div className="calBox"></div>
          <div className="calBox"></div>
          <div className="calBox"></div>
          <div className="calBox"></div>
        </div>
      </div>
    </div>
  );
}

export default Calendar;
