import { useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/Navbar";
import Button from "../components/Button"
import "../styles/login.css"

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8000/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        throw new Error("Credenciales incorrectas o error del servidor");
      }

      const data = await response.json();

      if (data.token) {
        localStorage.setItem("token", data.token);
        localStorage.setItem("user", JSON.stringify(data.user));
        localStorage.setItem("role", data.user.role);
        navigate("/menu");
      } else {
        throw new Error("No se recibió el token JWT del servidor");
      }

    } catch (error: any) {
      console.error("Error al iniciar sesión:", error);
    }
  };

  return (
    <div>
      <Navbar />
      <div className="login_container">
        <h1 className="titulo_inicio_sesion">Inicio de Sesión</h1>
        <div className="form_container">
          <form onSubmit={handleSubmit} className="login-form">
            <label htmlFor="email">Correo</label>
            <input
              type="email"
              placeholder="example@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              />
            <label htmlFor="password">Contraseña</label>
            <input
              type="password"
              placeholder="*********"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              />
              <Button type="submit" text="Ingresar"/>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
