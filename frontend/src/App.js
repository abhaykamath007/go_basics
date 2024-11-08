import React from "react";
import Login from "./pages/Login/Login";
import { Route, Routes, Navigate } from 'react-router-dom';
import Register from "./pages/Register/Register";
import HomePage from "./pages/Homepage/HomePage";
import AppLayout from "./components/Layout/AppLayout";
import BooksPage from "./pages/BooksPage/BooksPage";
import AccountPage from "./pages/AccountPage/AccountPage";
import NotFoundPage from "./pages/NotFoundPage/NotFoundPage";
import { useAuth } from "./context/AuthContext";

function App() {

  const {isAuthenticated} = useAuth();

  return (
      <Routes>
          <Route element={<AppLayout/>}>
            <Route path="/"  element={isAuthenticated ? <HomePage/> : <Navigate to='/login'/>} />
            <Route path="/books" element={isAuthenticated ? <BooksPage/> : <Navigate to='/login'/>}/>
            <Route path="/account" element={isAuthenticated ? <AccountPage/> : <Navigate to='/login'/>}/>
          </Route>

          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="*" element={<NotFoundPage/>} />

          <Route path="*" element={<NotFoundPage />} />
      </Routes>
  );
}

export default App;
