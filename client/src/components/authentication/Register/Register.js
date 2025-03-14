import { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { motion } from "framer-motion"; // Import animation

const RegisterPage = () => {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [phone, setPhone] = useState("");
  const [role, setRole] = useState("user");
  const [confirmPassword, setConfirmPassword] = useState(""); // New state for confirm password
  const [errors, setErrors] = useState({});

  useEffect(() => {
    import("bootstrap/dist/js/bootstrap.bundle.min");
  }, []);

  const handleRegister = async (e) => {
    e.preventDefault();
    let validationErrors = {};

    // Validation checks
    if (!name) {
      validationErrors.name = "Name is required.";
    }
    if (!email) {
      validationErrors.email = "Email is required.";
    }
    if (!phone) {
      validationErrors.phone = "Phone number is required.";
    }
    if (!password) {
      validationErrors.password = "Password is required.";
    }
    if (password !== confirmPassword) {
      validationErrors.confirmPassword = "Passwords do not match.";
    }

    if (Object.keys(validationErrors).length > 0) {
      setErrors(validationErrors);
      return;
    }

    try {
      const res = await axios.post("http://localhost:8080/register", {
        Name: name,
        Email: email,
        Password: password,
        Phone: phone,
        Role: role,
      });

      if (res.status === 201) {
        toast.success("Registration successful! Redirecting to login...");
        setTimeout(() => navigate("/login"), 2000);
      }
    } catch (error) {
      console.log(error);
      toast.error("Registration failed. Please try again.");
    }
  };

  return (
    <div>
      <motion.div
        className="d-flex"
        initial={{ opacity: 0, x: -100 }} // Bắt đầu từ bên trái
        animate={{ opacity: 1, x: 0 }} // Di chuyển vào giữa
        exit={{ opacity: 0, x: 100 }} // Rời khỏi sang phải
        transition={{ duration: 0.5 }}
      >
        <ToastContainer />
        <div className="main container d-flex align-items-center justify-content-center">
          <div className="col-md-6 col-sm-12">
            <div className="register-form">
              <form onSubmit={handleRegister}>
                <div className="form-group my-2">
                  <label className="form-label">Name</label>
                  <input
                    type="text"
                    className="form-control"
                    placeholder="Full Name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                  />
                  {errors.name && (
                    <small className="text-danger">{errors.name}</small>
                  )}
                </div>
                <div className="form-group my-2">
                  <label className="form-label">Email</label>
                  <input
                    type="email"
                    className="form-control"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                  />
                  {errors.email && (
                    <small className="text-danger">{errors.email}</small>
                  )}
                </div>
                <div className="form-group my-2">
                  <label className="form-label">Password</label>
                  <input
                    type="password"
                    className="form-control"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                  />
                  {errors.password && (
                    <small className="text-danger">{errors.password}</small>
                  )}
                </div>
                <div className="form-group my-2">
                  <label className="form-label">Confirm Password</label>
                  <input
                    type="password"
                    className="form-control"
                    placeholder="Confirm Password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                  />
                  {errors.confirmPassword && (
                    <small className="text-danger">
                      {errors.confirmPassword}
                    </small>
                  )}
                </div>
                <div className="form-group my-2">
                  <label className="form-label">Phone</label>
                  <input
                    type="text"
                    className="form-control"
                    placeholder="Phone Number"
                    value={phone}
                    onChange={(e) => setPhone(e.target.value)}
                  />
                  {errors.phone && (
                    <small className="text-danger">{errors.phone}</small>
                  )}
                </div>
                <div className="form-group my-2">
                  <label className="form-label">Role</label>
                  <select
                    className="form-control"
                    value={role}
                    onChange={(e) => setRole(e.target.value)}
                  >
                    <option value="user">User</option>
                    <option value="admin">Admin</option>
                  </select>
                </div>
                <div className="mt-4">
                  <button type="submit" className="btn btn-dark w-100 mb-2">
                    Register
                  </button>
                  <div className="d-flex align-items-center my-2">
                    <hr className="flex-grow-1" />
                    <span className="mx-2">or</span>
                    <hr className="flex-grow-1" />
                  </div>
                  <button
                    onClick={() => navigate("/login")}
                    className="btn btn-light w-100"
                  >
                    Login
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
        <div className="sidenav d-flex align-items-center justify-content-center text-white text-center">
          <div className="register-main-text">
            <h1>Restaurant</h1>
            <h3>Register Page</h3>
            <p>Create an account to get started.</p>
          </div>
        </div>
      </motion.div>
    </div>
  );
};

export default RegisterPage;
