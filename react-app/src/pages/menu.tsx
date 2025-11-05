//useState sirve para guardar datos dentro del componente
//useEffect sirve para ejecutar algo cuando el componente se muestra en pantalla
//menu es para las interfaces y decirle a ts como luce cada menu
import React, { useEffect, useState } from "react";
import { Menu } from "../interfaces/Menu";
import Navbar from "../components/Navbar";
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
))}
    </div>
  );
};

export default Menus;
