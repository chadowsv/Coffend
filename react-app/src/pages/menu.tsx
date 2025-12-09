//useState sirve para guardar datos dentro del componente
//useEffect sirve para ejecutar algo cuando el componente se muestra en pantalla
//menu es para las interfaces y decirle a ts como luce cada menu

import React, { useEffect, useState } from "react";
import { Menu } from "../interfaces/Menu";
import Navbar from "../components/Navbar";
import Card from "../components/Card";
import { isAdmin , getToken } from "../auth";
import "../styles/global.css";
import "../styles/menu.css";

const Menus: React.FC = () => {
  const [menus, setMenus] = useState<Menu[]>([]);
  const [showMenuModal, setShowMenuModal] = useState(false);

  const [newMenu, setNewMenu] = useState({
    name: "",
  });


  useEffect(() => {
    fetch("http://localhost:8000/menus")
      .then((res) => res.json())
      .then((data) =>
        setMenus(
          data.map((menu: any) => ({
            ...menu,
            foods: menu.foods || [],
          }))
        )
      );
  }, []);

  // Crear menú (POST)
  const handleCreateMenu = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = getToken();
    const res = await fetch("http://localhost:8000/menus", {
      method: "POST",
      headers: { "Content-Type": "application/json",
        Authorization: `Bearer ${token}`
       },
      body: JSON.stringify(newMenu),
    });

    const created = await res.json();

    setMenus((prev) => [...prev, { ...created, foods: [] }]);
    setShowMenuModal(false);
    setNewMenu({ name: "" });
  };

  return (
    <div>
      <Navbar />
      <h1>Menú</h1>

      {isAdmin() && (
        <button className="admin_button" onClick={() => setShowMenuModal(true)}>
          Agregar nuevo menú
        </button>
      )}

      {showMenuModal && (
        <div className="modal-overlay">
          <div className="modal">
            <h2>Crear Nuevo Menú</h2>

            <form onSubmit={handleCreateMenu}>
              <label>Nombre:</label>
              <input
                type="text"
                value={newMenu.name}
                onChange={(e) =>
                  setNewMenu({ ...newMenu, name: e.target.value })
                }
                required
              />

              <div className="modal-buttons">
                <button type="submit" className="save-btn">
                  Guardar
                </button>
                <button
                  type="button"
                  className="cancel-btn"
                  onClick={() => setShowMenuModal(false)}
                >
                  Cancelar
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* LISTADO */}
      <div className="container">
        <section className="menu-grid">
          {menus.map((menu) => (
            <Card key={menu.menu_id} className="menu-item">
              {/* Contenido de texto */}
              <div className="menu-content">
                <h2>{menu.name}</h2>

                {(menu.menu_id === 2 ||
                  menu.menu_id === 4 ||
                  menu.menu_id === 8) && (
                    <ul>
                      {menu.foods.map((food) => (
                        <li key={food.food_id}>
                          <strong>{food.name}</strong> - $
                          {food.price.toFixed(2)}
                          {food.description && <p>{food.description}</p>}
                        </li>
                      ))}
                    </ul>
                  )}

                {menu.menu_id !== 2 &&
                  menu.menu_id !== 4 &&
                  menu.menu_id !== 8 && (
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
                {menu.menu_id === 1 && (
                  <img
                    src="/img/bolones.jpg"
                    alt="Bolones"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 2 && (
                  <img
                    src="/img/tigrillos.jpg"
                    alt="Tigrillos"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 3 && (
                  <img
                    src="/img/empanadas.jpg"
                    alt="Empanadas"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 4 && (
                  <img
                    src="/img/canoademaduro.jpg"
                    alt="Especialidades"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 5 && (
                  <img
                    src="/img/bebidascalientes.jpg"
                    alt="Bebidas"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 6 && (
                  <img
                    src="/img/jugos.png"
                    alt="Jugos"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 7 && (
                  <img
                    src="/img/batidos.jpg"
                    alt="Batidos"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 8 && (
                  <img
                    src="/img/combos.jpg"
                    alt="Combos"
                    className="menu-image"
                  />
                )}
                {menu.menu_id === 9 && (
                  <img
                    src="/img/extras.jpg"
                    alt="Extras"
                    className="menu-image"
                  />
                )}
              </div>
            </Card>
          ))}
        </section>
      </div>
    </div>
  );
};

export default Menus;
