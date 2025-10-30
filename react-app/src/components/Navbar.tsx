export default function Navbar() {
  return (
    <nav className="navbar">
      <a href="/" className="navbar-logo">La Esquina del Verde</a>
      <div className="navbar-links">
        <a href="/login">Iniciar sesión</a>
        <a href="/register">Registrarse</a>
        <a href="/menu">Menú</a>
      </div>
    </nav>
  );
}