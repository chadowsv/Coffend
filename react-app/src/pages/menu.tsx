//useState sirve para guardar datos dentro del componente
//useEffect sirve para ejecutar algo cuando el componente se muestra en pantalla
//menu es para las interfaces y decirle a ts como luce cada menu
import React, { useEffect, useState } from "react";
import { Menu } from "../interfaces/Menu";
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
      <h1>Menús</h1>
      {menus.map((menu) => (
        <div key={menu.menu_id} className="menu">
          <h2 className="menuName">{menu.name}</h2>
          <ul>
            {menu.foods.map((food) => (
              <li key={food.food_id}>
                {food.name} - ${food.price.toFixed(2)}
              </li>
            ))}
          </ul>
        </div>
      ))}
    </div>
  );
};

export default Menus;
