import Navbar from "../components/Navbar";
import "../styles/sucursales.css";

const Sucursales = () => {
    return (
        <>
            <Navbar />
            <div className="sucursales-container">
                <h1>Sucursales</h1>

                {/* MACHACHI */}
                <div className="sucursal-card">
                    <div className="sucursal-content">
                        <div className="sucursal-info">
                            <h2>Sucursal Machachi</h2>
                            <p>
                                Pérez Pareja, y, Machachi
                            </p>
                        </div>
                        <iframe
                            className="sucursal-map"
                            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3030.2581161774783!2d-78.56813174474907!3d-0.5103475755923768!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x91d5afea3c8dfa9f%3A0x7180d65cd7b8aa2d!2sLa%20Esquina%20del%20verde!5e0!3m2!1ses!2sec!4v1765320726350!5m2!1ses!2sec"
                            loading="lazy"
                            allowFullScreen
                            referrerPolicy="no-referrer-when-downgrade"
                        ></iframe>
                    </div>
                </div>

                {/* TAMBILLO */}
                <div className="sucursal-card">
                    <div className="sucursal-content">
                        <div className="sucursal-info">
                            <h2>Sucursal Tambillo</h2>
                            <p>
                                Carretera Panamericana y Paraíso
                            </p>
                        </div>
                        <iframe
                            className="sucursal-map"
                            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3989.7187221977083!2d-78.5442219241594!3d-0.4048438995911802!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x91d5a5006a7db46b%3A0x5d7a3158f47502b!2sLa%20esquina%20del%20verde!5e0!3m2!1ses!2sec!4v1765320829900!5m2!1ses!2sec"
                            loading="lazy"
                            allowFullScreen
                            referrerPolicy="no-referrer-when-downgrade"
                        ></iframe>
                    </div>
                </div>

                {/* LA MERCED */}
                <div className="sucursal-card">
                    <div className="sucursal-content">
                        <div className="sucursal-info">
                            <h2>Sucursal La Merced</h2>
                            <p>
                                César Enrique Balseca, y, Quito 170810
                            </p>
                        </div>
                        <iframe
                            className="sucursal-map"
                            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3989.765358743464!2d-78.40576752415934!3d-0.29521749970187866!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x91d5bdb63e1a80c7%3A0xc2fbbeab8d9cd797!2sLa%20Esquina%20del%20Verde%20La%20Merced!5e0!3m2!1ses!2sec!4v1765320867300!5m2!1ses!2sec"
                            loading="lazy"
                            allowFullScreen
                            referrerPolicy="no-referrer-when-downgrade"
                        ></iframe>
                    </div>
                </div>
            </div>
        </>
    );
};

export default Sucursales;
