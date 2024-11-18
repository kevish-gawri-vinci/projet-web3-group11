// src/Components/ProtectedRoute.jsx
import React, { useContext } from 'react';
import { UserContext } from '../Context/UserContext';  // Assurez-vous que le chemin est correct

const ProtectedRoute = ({ role, children }) => {
    const { user } = useContext(UserContext);

    if (!auth.isAuthenticated || auth.role !== role) {
        return <Navigate to="/login" />;
    }

    return children;
};

export default ProtectedRoute;
