import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/Navbar";
import Button from "../components/Button";
import Card from "../components/Card";
import "../styles/global.css";
import "../styles/register.css";

const Register = () => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [repeatPassword, setRepeatPassword] = useState("");
  const [phone, setPhone] = useState("");
  const [role] = useState("cliente");

  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (password !== repeatPassword) {
      alert("Las contraseñas no coinciden.");
      return;
    }

    try {
      const response = await fetch("http://localhost:8000/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          email,
          password,
          phone,
          role,
        }),
      });

      if (!response.ok) {
        throw new Error("Error al registrar usuario");
      }

      navigate("/register/success"); // ✅ nueva página
    } catch (error: any) {
      console.error("Error al registrarse:", error);
      alert("No se pudo completar el registro.");
    }
  };

  return (
    <div>
      <Navbar />
      <div className="register_container">
        <h1 className="titulo_registro">Crear Cuenta</h1>

        <div className="form_container">
          <form onSubmit={handleSubmit} className="register-form">
            
            <label>Nombre</label>
            <input
              type="text"
              placeholder="Nombre"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
              required
            />

            <label>Apellido</label>
            <input
              type="text"
              placeholder="Apellido"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
              required
            />

            <label>Correo</label>
            <input
              type="email"
              placeholder="example@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />

            <label>Teléfono</label>
            <input
              type="tel"
              placeholder="0991234567"
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
              required
              pattern="[0-9]{10}"
            />

            <label>Contraseña</label>
            <input
              type="password"
              placeholder="*********"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />

            <label>Repetir contraseña</label>
            <input
              type="password"
              placeholder="*********"
              value={repeatPassword}
              onChange={(e) => setRepeatPassword(e.target.value)}
              required
            />

            <label>Rol</label>
            <select 
              value={role}
              onChange={() => {}}
              className="select-disabled"
            >
              <option value="cliente">Cliente</option>
            </select>

            <Button type="submit" text="Registrarse" />
          </form>
        </div>
      </div>
    </div>
  );
};

export default Register;
