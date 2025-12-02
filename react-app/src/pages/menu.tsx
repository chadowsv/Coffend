//useState sirve para guardar datos dentro del componente
//useEffect sirve para ejecutar algo cuando el componente se muestra en pantalla
//menu es para las interfaces y decirle a ts como luce cada menu
import React, { useEffect, useState } from "react";
import { Menu } from "../interfaces/Menu";
import Navbar from "../components/Navbar";
import Card from '../components/Card';
import { isAdmin } from "../auth";
import "../styles/global.css";
import "../styles/menu.css";

//Se crea un componente con React llamado Menus
const Menus: React.FC = () => {

  const [menus, setMenus] = useState<Menu[]>([]);

  useEffect(() => {
    fetch("http://localhost:8000/menus")
      .then((res) => res.json())
      .then((data) =>
        setMenus(
          data.map((menu: any) => ({
            ...menu,
            foods: menu.foods || [], // asegura que siempre haya un array
          }))
        ))
      .catch((error) => console.error("Error fetching menus:", error));
  }, []);

  return (
    <div>
      <Navbar />
      <h1>Menú</h1>
      {isAdmin() && (
        <button className="admin_button">Agregar nuevo menú</button>
      )}
      <div className="container">
        <section className="menu-grid">
          {menus.map((menu) => (
            <Card key={menu.menu_id} className="menu-item">
              {/* Contenido de texto */}
              <div className="menu-content">
                <h2>{menu.name}</h2>

                {(menu.menu_id === 2 || menu.menu_id === 4 || menu.menu_id === 8) && (
                  <ul>
                    {menu.foods.map((food) => (
                      <li key={food.food_id}>
                        <strong>{food.name}</strong> - ${food.price.toFixed(2)}
                        {food.description && <p>{food.description}</p>}
                      </li>
                    ))}
                  </ul>
                )}

                {menu.menu_id !== 2 && menu.menu_id !== 4 && menu.menu_id !== 8 && (
                  <ul>
                    {menu.foods.map((food) => (
                      <li key={food.food_id}>
                        {food.name} - ${food.price.toFixed(2)}
                      </li>
                    ))}
                  </ul>
                )}
              </div>

              {/* Imagen */}
              <div className="menu-header">
                {menu.menu_id === 1 && <img src="/img/bolones.jpg" alt="Bolones" className="menu-image" />}
                {menu.menu_id === 2 && <img src="/img/tigrillos.jpg" alt="Tigrillos" className="menu-image" />}
                {menu.menu_id === 3 && <img src="/img/empanadas.jpg" alt="Empanadas" className="menu-image" />}
                {menu.menu_id === 4 && <img src="/img/canoademaduro.jpg" alt="Especialidades" className="menu-image" />}
                {menu.menu_id === 5 && <img src="/img/bebidascalientes.jpg" alt="Bebidas" className="menu-image" />}
                {menu.menu_id === 6 && <img src="/img/jugos.png" alt="Jugos" className="menu-image" />}
                {menu.menu_id === 7 && <img src="/img/batidos.jpg" alt="Batidos" className="menu-image" />}
                {menu.menu_id === 8 && <img src="/img/combos.jpg" alt="Combos" className="menu-image" />}
                {menu.menu_id === 9 && <img src="/img/extras.jpg" alt="Extras" className="menu-image" />}
              </div>
            </Card>
          ))}
        </section>
      </div>
    </div>
  );
};

export default Menus;