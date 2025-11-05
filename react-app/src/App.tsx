import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Menu from "./pages/menu";
import Register from "./pages/register";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />          {/* Página principal */}
        <Route path="/login" element={<Login />} />    {/* Página de inicio de sesión */}
        <Route path="/menu" element={<Menu />} />    {/* Página de menu*/}
        <Route path="/register" element={<Register />} /> {/* Página de registro */}
      </Routes>
    </Router>
  );
}

export default App;