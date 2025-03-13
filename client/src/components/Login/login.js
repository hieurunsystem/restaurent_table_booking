import axios from "axios";
import { useEffect, useState } from "react";

const LoginPage = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/user_list")
      .then((response) => setUsers(response.data)) // Không cần .json()
      .catch((error) => console.error("Error fetching users:", error));
  }, []);
  console.log(users[0])
  return (
    <div>
      <h2>User List</h2>
      <ul>
        {Array.isArray(users) ? users.map((user) => (
          <li key={user.Id}>{user.Name} - {user.Email}</li>
        )) : <p>No users found</p>}
      </ul>
    </div>
  );
};

export default LoginPage;
