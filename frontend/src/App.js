import axios from "axios";
import React, { useEffect, useState } from "react";

const App = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    axios
      .get("http://127.0.0.1:8000/api/v1/users")
      .then((response) => {
        console.log(response?.data?.data);
        setUsers(response?.data?.data);
      })
      .catch((err) => {
        console.log(`Error in fetching users ${err}`);
      });
  }, []);

  return (
    <>
      {/* NavBar */}
      <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
        <div className="container-fluid">
          <a className="navbar-brand" href="#">
            Docker Compose Tutorial
          </a>
        </div>
      </nav>

      {/* Content */}
      <div className="container mt-3">
        {users?.map((user) => (
          <div className="d-flex flex-column mb-3">
            <div className="card" key={user?.id}>
              <div className="card-body">
                <h5 className="card-title">{user?.firstName}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{user?.email}</h6>
                <p className="card-text">{user?.description}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default App;
