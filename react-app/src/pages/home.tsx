import React from "react";
import Navbar from "../components/Navbar";
import "../styles/home.css";

const Home: React.FC = () => {
  return (
    <div>
      <Navbar />
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
