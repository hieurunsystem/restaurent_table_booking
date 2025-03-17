import React, { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const Personal = () => {
  const [user, setUser] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    axios
      .get("http://localhost:8080/me", { withCredentials: true })
      .then((res) => {
        setUser(res.data);
      })
      .catch((err) => {
        console.log("Error fetching user data", err);
      });
  }, []);

  return (
    <div>
      <h1>Personal Information</h1>
      {user ? (
        <>
          <p>Name: {user.Name}</p>
          <p>Role: {user.Role}</p>
          {user.Role === "customer" && (
            <>
              <p>Email: {user.Email}</p>
              <p>Phone: {user.Phone}</p>
              <p>Other ID: {user.orther_id}</p>
              <a href="/reset-password">Reset Password</a>
            </>
          )}
          {user.Role === "owner" && (
            <>
              <p>Email: {user.Email}</p>
              <p>Phone: {user.Phone}</p>
              <p>Other ID: {user.orther_id}</p>
              <p>Owner-specific information here</p>
            </>
          )}
          {user.Role === "staff" && (
            <>
              <p>Email: {user.Email}</p>
              <p>Phone: {user.Phone}</p>
              <p>Other ID: {user.orther_id}</p>
              <p>Staff-specific information here</p>
            </>
          )}
          {user.Role === "admin" && (
            <>
              <p>Email: {user.Email}</p>
              <p>Phone: {user.Phone}</p>
              <p>Other ID: {user.orther_id}</p>
              <p>Admin-specific information here</p>
            </>
          )}
        </>
      ) : (
        <>
          <p>Chưa đăng nhập</p>
          <button onClick={() => navigate("/login")}>Login</button>
        </>
      )}
    </div>
  );
};

export default Personal;
