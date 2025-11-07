import { Link } from "react-router-dom";
import Navbar from "../components/Navbar";
import "../styles/register.css";

const RegisterSuccess = () => {
  return (
    <div>
      <Navbar />

      <div className="success_container">
        <div className="success_box">
          <h1>Registro Exitoso</h1>
          <p>Tu cuenta ha sido creada correctamente.</p>

          <Link to="/login" className="success_btn">
            Ir a iniciar sesión
          </Link>
        </div>
      </div>
    </div>
  );
};

export default RegisterSuccess;
