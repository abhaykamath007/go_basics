
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
import AdminDashboard from "./pages/admin/AdminDashboard/AdminDashboard";
import AdminLayout from "./components/Layout/AdminLayout";
import AdminWrapper from "./components/Layout/AdminWrapper";
import ManageBooks from "./pages/admin/ManageBooks/ManageBooks";
import ManageUsers from "./pages/admin/ManageUsers/ManageUsers";

function App() {

  const { isAuthenticated } = useAuth();
  const role = localStorage.getItem("Role");

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


        <Route element={<AdminWrapper />}>
          <Route element={<AdminLayout />}>
            <Route path="/admin" element={<AdminDashboard />} />
            <Route path="/admin/manage-books" element={<ManageBooks />} />
            <Route path="/admin/manage-users" element={<ManageUsers />} />
          </Route>
        </Route>


        <Route path="/login" element={
          isAuthenticated
            ? role === "admin"
              ? <Navigate to="/admin" replace />
              : <Navigate to="/" replace />
            : <Login />
        } />

        <Route path="/register" element={
          isAuthenticated ? role === "admin"
            ? <Navigate to="/admin" replace />
            : <Navigate to="/" replace />
            : <Register />
        } />

        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </>

  );
}

export default App;
