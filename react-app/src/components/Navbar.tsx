export default function Navbar() {
  return (
    <nav className="navbar">
      {/* Imagen totalmente a la izquierda */}
      <a href="/" className="navbar-logo-img">
        <img src="/img/logo.webp" alt="logo" />
      </a>

      {/* Todo lo demás permanece igual */}
      <a href="/" className="navbar-logo">La Esquina Del Verde</a>

      <div className="navbar-links">
        <a href="/menu">Menú</a>
        <a href="/sucursales">Sucursales</a>
        <a href="/mesas">Mesas</a>
        <a href="/login">Login</a>
        <a href="/Register">Registrarse</a>
      </div>
    </nav>
  );
}

