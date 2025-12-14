import React from "react";
import { Link } from "react-router-dom";
import Navbar from "../components/Navbar";
import Card from "../components/Card";
import "../styles/auth_forms.css";

const RegisterSuccess = () => {
  return (
    <div>
      <Navbar />

      <div className="container center" style={{ minHeight: "60vh" }}>
        <Card className="stack" style={{ maxWidth: 520 }}>
          <div className="success-container">
            <div className="success-box">
              <h1>Registro Exitoso</h1>
              <p>Tu cuenta ha sido creada correctamente.</p>

              <Link to="/login" className="success-btn">
                Ir a iniciar sesión
              </Link>
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default RegisterSuccess;
