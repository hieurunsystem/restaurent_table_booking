import { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min";
import $ from "jquery";
import axios from "axios";

const LoginPage = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    import("bootstrap/dist/js/bootstrap.bundle.min");
    axios
      .get("http://localhost:8080/user_list")
      .then((response) => setUsers(response.data.users)) // Không cần .json()
      .catch((error) => console.error("Error fetching users:", error));
  }, []);

  return (
    <div className="d-flex">
      <div className="sidenav d-flex align-items-center justify-content-center text-white text-center">
        <div className="login-main-text">
          <h2>
            Kien's Restaurant
            <br /> Login Page
          </h2>
          <p>Login or register from here to access.</p>
          <div>
            <h2>User List</h2>
            <ul>
              {Array.isArray(users) ? (
                users.map((user) => (
                  <li key={user.Id}>
                    {user.Name} - {user.Email}
                  </li>
                ))
              ) : (
                <p>No users found</p>
              )}
            </ul>
          </div>
        </div>
      </div>
      <div className="main container d-flex align-items-center justify-content-center">
        <div className="col-md-6 col-sm-12">
          <div className="login-form">
            <form>
              <div className="form-group my-2">
                <label>User Name</label>
                <input
                  type="text"
                  className="form-control"
                  placeholder="User Name"
                />
              </div>
              <div className="form-group my-2">
                <label>Password</label>
                <input
                  type="password"
                  className="form-control"
                  placeholder="Password"
                />
              </div>
              <button type="submit" className="btn btn-black">
                Login
              </button>
              <button type="submit" className="btn btn-secondary mx-2">
                Register
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;

// CSS styles
const styles = `
  .sidenav {
    height: 100vh;
    width: 40%;
    background-color: #000;
    color: #fff;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .login-main-text {
    padding: 60px;
    text-align: center;
  }
  .main {
    width: 60%;
    padding: 20px;
  }
  .btn-black {
    background-color: #000 !important;
    color: #fff;
  }
  @media (max-width: 768px) {
    .sidenav {
      width: 100%;
      height: 40vh;
    }
    .main {
      width: 100%;
      padding-top: 20px;
    }
  }
`;

const styleSheet = document.createElement("style");
styleSheet.type = "text/css";
styleSheet.innerText = styles;
document.head.appendChild(styleSheet);
