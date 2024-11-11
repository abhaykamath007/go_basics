
import { Outlet, Navigate } from 'react-router-dom';

const ProtectedLayout = ({ isAuthenticated }) => {
  return isAuthenticated ? <Outlet /> : <Navigate to="/login" />;
};

export default ProtectedLayout;