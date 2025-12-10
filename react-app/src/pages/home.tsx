import React from "react";
import Navbar from "../components/Navbar";
import "../styles/global.css";
import "../styles/home.css"

const Home: React.FC = () => {
  return (
    <div>
      <Navbar />
      <div className="container">
        <section id="quienes-somos" className="sections-home">
          <h2 className="section_title">Quiénes Somos</h2>
          <p className="section_text">
            La Esquina del Verde nació en 2010 en el corazón de Manabí, cuando un grupo de amigos apasionados por la comida tradicional manaba decidió compartir los sabores de su tierra con la comunidad. 
            Comenzamos con una pequeña esquina donde la gente podía disfrutar de platos frescos y llenos de sabor, elaborados con ingredientes locales y recetas heredadas de nuestras familias. 
            Hoy, seguimos manteniendo esa esencia y orgullo de nuestras raíces, ofreciendo un espacio acogedor donde cada plato cuenta una historia.
          </p>
          
        </section>

        <img src="/img/cocina.jpg" alt="imagen_de_los_cocineros" className="imagen_cocineros" />

        <div className="sections-union">
          <section id="vision" className="sections-home">
            <h2 className="section_title">Visión</h2>
            <p className="section_text">
              Ser el restaurante referente de la cocina manaba en Ecuador, reconocidos por nuestra autenticidad, calidad y compromiso con la comunidad. 
              Queremos que cada cliente que visite La Esquina del Verde sienta la calidez de nuestra tierra y la tradición en cada bocado.
            </p>
          </section>

          <section id="mision" className="sections-home">
            <h2 className="section_title">Misión</h2>
            <p className="section_text">
              Nuestra misión es ofrecer platos manabitas deliciosos y frescos, utilizando productos locales de calidad, respetando la tradición y fomentando la cultura gastronómica de Manabí. 
              Buscamos crear experiencias memorables para cada comensal y apoyar a los productores locales, promoviendo el desarrollo sostenible de nuestra región.
            </p>
          </section>
        </div>

        <section id="contacto" className="sections-home">
          <h2 className="section_title">Contacto</h2>
          <p className="section_text">
            Visítanos en nuestra esquina en Manabí o contáctanos vía correo electrónico: <a href="mailto:contacto@laesquinadelverde.com">contacto@laesquinadelverde.com</a>
          </p>
        </section>

        <footer className="footer">
          <p className="footer_text">&copy; 2025 La Esquina del Verde. Todos los derechos reservados.</p>
        </footer>
      </div>
    </div>
  );
};

export default Home;