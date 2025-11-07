//useState sirve para guardar datos dentro del componente
//useEffect sirve para ejecutar algo cuando el componente se muestra en pantalla
//menu es para las interfaces y decirle a ts como luce cada menu
import React, { useEffect, useState } from "react";
import { Menu } from "../interfaces/Menu";
import Navbar from "../components/Navbar";
import "../styles/menu.css"
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
      {menus.map((menu) => (
        <div key={menu.menu_id} className="menu">
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
            {/* Imagen: Bolones de verde 
            Fuente: Pixabay (https://pixabay.com/photos/bolon-de-verde-ecuador)
            */}
            {menu.menu_id === 1 && <img src="/img/bolones.jpg" alt="Bolones de verde - Fuente: Pixabay" className="menu-image" />}

            {/* Imagen: Tigrillos 
            Fuente: Unsplash (https://unsplash.com/photos/tigrillos-ecuador)
            */}
            {menu.menu_id === 2 && <img src="/img/tigrillos.jpg" alt="Tigrillos - Fuente: Unsplash" className="menu-image" />}

            {/* Imagen: Empanadas 
            Fuente: Pixabay (https://pixabay.com/photos/empanadas)
            */}
            {menu.menu_id === 3 && <img src="/img/empanadas.jpg" alt="Empanadas - Fuente: Pixabay" className="menu-image" />}

            {/* Imagen: Canoa de maduro 
            Fuente: Pixabay (https://pixabay.com/photos/canoa-de-maduro-ecuador)
            */}
            {menu.menu_id === 4 && <img src="/img/canoademaduro.jpg" alt="Canoa de maduro - Fuente: Pixabay" className="menu-image" />}

            {/* Imagen: Bebidas calientes 
            Fuente: Unsplash (https://unsplash.com/photos/hot-drinks)
            */}
            {menu.menu_id === 5 && <img src="/img/bebidascalientes.jpg" alt="Bebidas calientes - Fuente: Unsplash" className="menu-image" />}

            {/* Imagen: Jugos naturales 
            Fuente: Pixabay (https://pixabay.com/photos/fruit-juice)
            */} 
            {menu.menu_id === 6 && <img src="/img/jugos.png" alt="Jugos naturales - Fuente: Pixabay" className="menu-image" />}

            {/* Imagen: Batidos 
            Fuente: Unsplash (https://unsplash.com/photos/smoothies)
            */}
            {menu.menu_id === 7 && <img src="/img/batidos.jpg" alt="Batidos - Fuente: Unsplash" className="menu-image" />}

            {/* Imagen: Combos 
            Fuente: Pixabay (https://pixabay.com/photos/food-combo)
            */}
            {menu.menu_id === 8 && <img src="/img/combos.jpg" alt="Combos - Fuente: Pixabay" className="menu-image" />}

            {/* Imagen: Extras o acompañamientos 
            Fuente: Pixabay (https://pixabay.com/photos/extras-food)
            */}
            {menu.menu_id === 9 && <img src="/img/extras.jpg" alt="Extras - Fuente: Pixabay" className="menu-image" />}

          </div>
        </div>

      ))}
    </div>
  );
};

export default Menus;
