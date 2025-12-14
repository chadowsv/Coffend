import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Menus from "./pages/menu";
import Register from "./pages/register";
import RegisterSuccess from "./pages/RegisterSuccess";
import Sucursales from "./pages/sucursales";
import Mesas from "./pages/mesas";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />        
        <Route path="/login" element={<Login />} />    
        <Route path="/menus" element={<Menus />} />    
        <Route path="/sucursales" element={<Sucursales />} /> 
        <Route path="/mesas" element={<Mesas />} /> 
        <Route path="/register" element={<Register />} /> {/* Página de registro */}
        <Route path="/register/success" element={<RegisterSuccess />} /> {/* Página de registro exitoso */}
      </Routes>
    </Router>
  );
}

export default App;