import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Menu from "./pages/menu";
import Register from "./pages/register";
import RegisterSuccess from "./pages/RegisterSuccess";
import Sucursales from "./pages/sucursales";
import Mesas from "./pages/mesas";
import Menu2 from "./pages/menu2";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />        
        <Route path="/login" element={<Login />} />    
        <Route path="/menu" element={<Menu />} />    
        <Route path="/sucursales" element={<Sucursales />} /> 
        <Route path="/mesas" element={<Mesas />} /> 
        <Route path="/register" element={<Register />} /> {/* Página de registro */}
        <Route path="/register/success" element={<RegisterSuccess />} /> {/* Página de registro exitoso */}
        <Route path="/menu2" element={<Menu2 />} /> 
      </Routes>
    </Router>
  );
}

export default App;