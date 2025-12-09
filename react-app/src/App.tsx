import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Menu from "./pages/menu";
import Register from "./pages/register";
import RegisterSuccess from "./pages/RegisterSuccess";
import Sucursales from "./pages/sucursales";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />          {/* Página principal */}
        <Route path="/login" element={<Login />} />    {/* Página de inicio de sesión */}
        <Route path="/menu" element={<Menu />} />    {/* Página de menu*/}
        <Route path="/sucursales" element={<Sucursales />} /> {/* Página de sucursales */}
        <Route path="/register" element={<Register />} /> {/* Página de registro */}
        <Route path="/register/success" element={<RegisterSuccess />} /> {/* Página de registro exitoso */}
      </Routes>
    </Router>
  );
}

export default App;