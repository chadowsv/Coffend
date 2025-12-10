export default function Navbar() {
  return (
    <nav className="navbar">
      <a href="/" className="navbar-logo">La Esquina del Verde</a>
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
