import React from 'react'
import { Navigate, Outlet } from 'react-router-dom';

const AdminLayout = ({ children }) => {

    const role = localStorage.getItem("Role");
    if (!role || role !== "admin") {
        return <Navigate to="/" replace />;
    }
    return <Outlet />;
}

export default AdminLayout
