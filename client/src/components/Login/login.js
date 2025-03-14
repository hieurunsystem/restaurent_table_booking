import { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min";
import $ from "jquery";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const LoginPage = () => {
  const navigate = useNavigate(); // Dùng để chuyển hướng trang
  let [Email, setEmail] = useState("");
  let [Password, setPassword] = useState("");

  useEffect(() => {
    import("bootstrap/dist/js/bootstrap.bundle.min");
  }, []);

  async function HandleLogin(e) {
    e.preventDefault(); // Ngăn chặn hành vi reload trang mặc định

    try {
      const response = await axios.post("http://localhost:8080/login", {
        Email,
        Password,
      });

      if (response.status === 200) {
        toast.success("Login successful!");
        console.log("User data:", response.data);

        const { role, tokens } = response.data;

        // Lưu vào localStorage
        localStorage.setItem("role", role);
        localStorage.setItem("token", tokens);

        // Wait for 2 seconds before navigating
        setTimeout(() => {
          // Điều hướng dựa vào role
          if (role === "admin") {
            navigate("/admin"); // Chuyển sang trang admin
          } else {
            navigate("/"); // Chuyển sang trang user thông thường
          }
        }, 1000);
      } else {
        toast.error("Login failed. Please check your credentials.");
      }
    } catch (error) {
      console.error("Login error:", error);
      toast.error("Error during login. Please try again.");
    }
  }

  return (
    <div className="d-flex">
      <ToastContainer />
      <div className="sidenav d-flex align-items-center justify-content-center text-white text-center">
        <div className="login-main-text">
          <h1>Restaurant</h1>
          <h3>Login Page</h3>
          <p>Login or register from here to access.</p>
        </div>
      </div>
      <div className="main container d-flex align-items-center justify-content-center">
        <div className="col-md-6 col-sm-12">
          <div className="login-form">
            <form onSubmit={HandleLogin}>
              <div className="form-group my-2">
                <label className="form-label">Email</label>{" "}
                {/* Added Bootstrap class "form-label" */}
                <input
                  type="text"
                  className="form-control"
                  placeholder="Email"
                  value={Email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div className="form-group my-2">
                <label className="form-label">Password</label>{" "}
                {/* Added Bootstrap class "form-label" */}
                <input
                  type="password"
                  className="form-control"
                  placeholder="Password"
                  value={Password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              <div className="mt-4">
                <button type="submit" className="btn btn-dark w-100 mb-2">
                  Login
                </button>
                <div className="d-flex align-items-center my-2">
                  <hr className="flex-grow-1" />
                  <span className="mx-2">or</span>
                  <hr className="flex-grow-1" />
                </div>{" "}
                {/* Added line dividers */}
                <button type="submit" className="btn btn-light w-100">
                  Register
                </button>
              </div>
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
    background-color: #343a40;
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
    background-color: #f8f9fa;
  }
  .btn-dark {
    background-color: #343a40 !important;
    color: #fff;
  }
  .btn-light {
    background-color: #f8f9fa !important;
    color: #343a40;
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
