import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Menu from "./pages/menu";
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
        <Route path="/menu" element={<Menu />} />    
        <Route path="/sucursales" element={<Sucursales />} /> 
        <Route path="/mesas" element={<Mesas />} /> 
        <Route path="/register" element={<Register />} /> 
        <Route path="/register/success" element={<RegisterSuccess />} /> 
      </Routes>
    </Router>
  );
}

export default App;