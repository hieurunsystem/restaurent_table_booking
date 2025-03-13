import HomePage from "../../pages/HomePage";
import BookingPage from "../../pages/BookingPage";
import ConfirmedBookingPage from "../../pages/ConfirmedBookingPage";
import { Routes, Route } from "react-router-dom";

import Login from "../Login/login";

const Main = () => {
  return (
    <>
      {/* ROUTES */}
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/bookings" element={<BookingPage />} />
        <Route path="/confirmed" element={<ConfirmedBookingPage />} />
        <Route path="/login" element={<Login />} />
      </Routes>
    </>
  );
};

export default Main;
