import { useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/Navbar";
import Button from "../components/Button";
import "../styles/register.css"; 

const Register = () => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [phone, setPhone] = useState("");
  const [role, setRole] = useState("cliente");
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

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

      const data = await response.json();

      if (data) {
        alert("Registro exitoso. Ahora puedes iniciar sesión.");
        navigate("/login");
      }
    } catch (error: any) {
      console.error("Error al registrarse:", error);
      alert("No se pudo completar el registro. Verifica los datos e inténtalo nuevamente.");
    }
  };

  return (
    <div>
      <Navbar />
      <div className="register_container">
        <h1 className="titulo_registro">Crear Cuenta</h1>
        <div className="form_container">
          <form onSubmit={handleSubmit} className="register-form">
            <label htmlFor="firstName">Nombre</label>
            <input
              type="text"
              placeholder="Nombre"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
              required
            />

            <label htmlFor="lastName">Apellido</label>
            <input
              type="text"
              placeholder="Apellido"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
              required
            />

            <label htmlFor="email">Correo</label>
            <input
              type="email"
              placeholder="example@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />

            <label htmlFor="phone">Teléfono</label>
            <input
              type="tel"
              placeholder="0991234567"
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
              required
              pattern="[0-9]{10}"
            />

            <label htmlFor="password">Contraseña</label>
            <input
              type="password"
              placeholder="*********"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />

            <label htmlFor="role">Rol</label>
            <select value={role} onChange={(e) => setRole(e.target.value)}>
              <option value="cliente">Cliente</option>
              <option value="admin">Administrador</option>
            </select>

            <Button type="submit" text="Registrarse" />
          </form>
        </div>
      </div>
    </div>
  );
};

export default Register;
