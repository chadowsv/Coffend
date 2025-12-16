import React, { useState, useEffect } from 'react';
import { isLoggedIn, logout } from '../auth';
import { isAdmin } from "../auth";
import { useNavigate } from 'react-router-dom';
import Button from './Button';

export default function Navbar() {

  const [adminUser, setAdminUser] = useState(false);
  const [logged, setLogged] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    setLogged(isLoggedIn());
    setAdminUser(isAdmin());
  }, []);

  const handleLogout = () => {
    logout();
    setLogged(false);
    navigate('/login');
  }

  return (
    <nav className="navbar">
  
      <a href="/" className="navbar-logo-img">
        <img src="/img/logo.webp" alt="logo" />
      </a>
      
      <div className="navbar-links">
        <a href="/menus">Menú</a>
        <a href="/sucursales">Sucursales</a>

        {adminUser ? (
          <>
            <a href="/mesas">Mesas</a>
          </>
        ) : null}

        {!logged ? (
          <>
            <a href="/login">Login</a>
            <a href="/Register">Registrarse</a>
          </>
        ) : (
          <Button type="button" text="Logout" onClick={handleLogout} className="btn-secondary" />
        )}
      </div>
    </nav>
  );
}

