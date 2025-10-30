import React, { useEffect, useState } from "react";
import { Menu } from "../interfaces/Menu";

const Menus: React.FC = () => {
  const [menus, setMenus] = useState<Menu[]>([]);

  useEffect(() => {
    fetch("http://localhost:8000/menus")
      .then((res) => res.json())
      .then((data) => setMenus(data))
      .catch((error) => console.error("Error fetching menus:", error));
  }, []);

  return (
    <div>
      <h1>Menús</h1>
      {menus.map((menu) => (
        <div key={menu.menu_id}>
          <h2>{menu.name}</h2>
          <p>Estado: {menu.menu_status ? "Activo" : "Inactivo"}</p>
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
