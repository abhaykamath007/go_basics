import React from 'react'
import { Navigate, Outlet } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';

const AdminLayout = ({ children }) => {
    
    const {isAuthenticated} = useAuth();

    const role = localStorage.getItem("Role");
    if (!isAuthenticated || role !== "admin") {
        return <Navigate to="/" replace />;
    }
    return <Outlet />;
}

export default AdminLayout
