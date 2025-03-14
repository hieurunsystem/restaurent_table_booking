import axios from "axios";

const HandleLogout = async () => {
  try {
    await axios.post(
      "http://localhost:8080/logout",
      {},
      { withCredentials: true }
    );
    window.location.href = "/";
  } catch (error) {
    console.error("Logout failed:", error);
  }
};

export default HandleLogout;
