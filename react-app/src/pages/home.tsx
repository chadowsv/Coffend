import React from "react";
import { Link } from "react-router-dom";
import "../styles/home.css";

const Home: React.FC = () => {
  return (
    <div>
        <nav className="home-nav">
          <Link to="/login">Iniciar sesión</Link>
          <Link to="/register">Registrarse</Link>
        </nav>
      <main className="home-container">
        <header className="home-header">
          <h1>La Esquina del Verde</h1>
          <p>Gestiona pedidos, facturas y mucho más desde una sola plataforma.</p>
        </header>

        <footer className="home-footer">
          <p>© {new Date().getFullYear()} La Esquina del Verde — Todos los derechos reservados.</p>
        </footer>
      </main>
    </div>
  );
};

export default Home;
