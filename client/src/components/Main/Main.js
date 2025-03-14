import { Routes, Route } from "react-router-dom";
import HomePage from "../../pages/HomePage";
import BookingPage from "../../pages/BookingPage";
import ConfirmedBookingPage from "../../pages/ConfirmedBookingPage";
import Login from "../authentication/Login/Login";
import RegisterPage from "../authentication/Register/Register";
import Admin from "../Admin/Admin";
const Main = () => {
  return (
    <Routes future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
      <Route path="/" element={<HomePage />} />
      <Route path="/bookings" element={<BookingPage />} />
      <Route path="/confirmed" element={<ConfirmedBookingPage />} />
      <Route path="/login" element={<Login />} />
      <Route path="/admin" element={<Admin />} />
      <Route path="/register" element={<RegisterPage />} />
    </Routes>
  );
};

export default Main;
