import React, { useEffect } from "react";
import Login from "./pages/Login/Login";
import { Route, Routes, Navigate } from 'react-router-dom';
import Register from "./pages/Register/Register";
import HomePage from "./pages/Homepage/HomePage";
import AppLayout from "./components/Layout/AppLayout";
import BooksPage from "./pages/BooksPage/BooksPage";
import AccountPage from "./pages/AccountPage/AccountPage";
import NotFoundPage from "./pages/NotFoundPage/NotFoundPage";
import { useAuth } from "./context/AuthContext";
import BookDetailsPage from "./pages/BookDetailsPage/BookDetailsPage";
import ProtectedLayout from "./components/Layout/ProtectedLayout";
import { ToastContainer } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';

function App() {

  const { isAuthenticated } = useAuth();



  return (
    <>   
     <ToastContainer position="top-center" autoClose={3000} />
      <Routes>
        <Route element={<AppLayout />}>
          <Route element={<ProtectedLayout isAuthenticated={isAuthenticated} />}>
            <Route path="/" element={<HomePage />} />
            <Route path="/books" element={<BooksPage />} />
            {console.log(isAuthenticated)};
            <Route path="/account" element={<AccountPage />} />
            <Route path="/books/:id" element={<BookDetailsPage />} />
          </Route>
        </Route>
        <Route path="/login" element={isAuthenticated ? <Navigate to="/" replace /> : <Login />} />
        <Route path="/register" element={isAuthenticated ? <Navigate to="/" replace /> : <Register />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </>

  );
}

export default App;
