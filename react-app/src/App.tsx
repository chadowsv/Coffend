import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Menu from "./pages/menu";
import Register from "./pages/register";
import RegisterSuccess from "./pages/RegisterSuccess";
<<<<<<< HEAD
import Sucursales from "./pages/sucursales";
=======
import Mesas from "./pages/mesas";
>>>>>>> c3506b7 (Mesas)

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />          {/* Página principal */}
        <Route path="/login" element={<Login />} />    {/* Página de inicio de sesión */}
        <Route path="/menu" element={<Menu />} />    {/* Página de menu*/}
<<<<<<< HEAD
        <Route path="/sucursales" element={<Sucursales />} /> {/* Página de sucursales */}
=======
        <Route path="/mesas" element={<Mesas />} /> 
>>>>>>> c3506b7 (Mesas)
        <Route path="/register" element={<Register />} /> {/* Página de registro */}
        <Route path="/register/success" element={<RegisterSuccess />} /> {/* Página de registro exitoso */}
      </Routes>
    </Router>
  );
}

export default App;