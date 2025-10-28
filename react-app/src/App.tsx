import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />          {/* Página principal */}
        <Route path="/login" element={<Login />} />    {/* Página de inicio de sesión */}
      </Routes>
    </Router>
  );
}

export default App;