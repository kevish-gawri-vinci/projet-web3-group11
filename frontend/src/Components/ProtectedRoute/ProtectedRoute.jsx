// src/Components/ProtectedRoute.jsx
import React, { useContext } from 'react';
import { UserContext } from '../Context/UserContext';  // Assurez-vous que le chemin est correct

const ProtectedRoute = ({ role, children }) => {
    const { user } = useContext(UserContext);

    if (!user || user.role !== role) {
        return <p>Access Forbidden</p>;
    }

    return children;
};

export default ProtectedRoute;
