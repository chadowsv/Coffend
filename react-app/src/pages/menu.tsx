import React, { useEffect, useState } from "react";
import { Menu, Food } from "../interfaces/Menu";
import Navbar from "../components/Navbar";
import Card from "../components/Card";
import { isAdmin, getToken } from "../auth";
import "../styles/global.css";
import "../styles/menu.css";

interface FoodForm {
  name: string;
  description: string;
  price: number;
}

const Menus: React.FC = () => {
  const [menus, setMenus] = useState<Menu[]>([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [adminUser, setAdminUser] = useState(false);
  const [newMenuName, setNewMenuName] = useState("");
  const [foods, setFoods] = useState<FoodForm[]>([
    { name: "", description: "", price: 0 },
  ]);
  const token = getToken();

  // Verificar si es admin
  useEffect(() => {
    if (isAdmin()) {
      setAdminUser(true);
    }
  }, []);

  // Cargar menús del backend
  useEffect(() => {
    const fetchMenus = async () => {
      try {
        const response = await fetch("http://localhost:8000/menus");
        if (!response.ok) {
          throw new Error("Error fetching menus");
        }
        const data = await response.json();
        setMenus(data || []);
      } catch (error) {
        console.error("Error al cargar menús:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchMenus();
  }, []);

  // Manejar cambio en el nombre del menú
  const handleMenuNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewMenuName(e.target.value);
  };

  // Manejar cambios en los alimentos
  const handleFoodChange = (
    index: number,
    field: keyof FoodForm,
    value: any
  ) => {
    const newFoods = [...foods];
    newFoods[index] = { ...newFoods[index], [field]: value };
    setFoods(newFoods);
  };

  // Agregar un nuevo campo de alimento
  const addFoodField = () => {
    setFoods([...foods, { name: "", description: "", price: 0 }]);
  };

  // Eliminar un campo de alimento
  const removeFoodField = (index: number) => {
    setFoods(foods.filter((_, i) => i !== index));
  };

  // Crear nuevo menú con sus alimentos
  const handleCreateMenu = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!newMenuName.trim()) {
      alert("Por favor ingresa un nombre para el menú");
      return;
    }

    if (foods.some((food) => !food.name.trim())) {
      alert("Por favor completa el nombre de todos los alimentos");
      return;
    }

    try {
      // Crear el menú
      const menuResponse = await fetch("http://localhost:8000/menus", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ name: newMenuName, menu_status: true }),
      });

      if (!menuResponse.ok) {
        throw new Error("Error creando menú");
      }

      const newMenu = await menuResponse.json();

      // Crear los alimentos para el menú
      for (const food of foods) {
        if (food.name.trim()) {
          await fetch("http://localhost:8000/foods", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
              name: food.name,
              description: food.description,
              price: parseFloat(food.price.toString()),
              menu_id: newMenu.menu_id,
            }),
          });
        }
      }

      // Actualizar la lista de menús
      const updatedMenus = await fetch("http://localhost:8000/menus");
      const updatedData = await updatedMenus.json();
      setMenus(updatedData || []);

      // Limpiar formulario y cerrar modal
      setNewMenuName("");
      setFoods([{ name: "", description: "", price: 0 }]);
      setShowModal(false);

      alert("Menú creado exitosamente");
    } catch (error) {
      console.error("Error al crear menú:", error);
      alert("Error al crear el menú");
    }
  };

  // Eliminar menú
  const handleDeleteMenu = async (menuId: number) => {
    if (!window.confirm("¿Estás seguro de que deseas eliminar este menú?")) {
      return;
    }

    try {
      // Obtener el menú con sus alimentos
      const menu = menus.find((m) => m.menu_id === menuId);

      // Eliminar todos los alimentos del menú primero
      if (menu && menu.foods && menu.foods.length > 0) {
        for (const food of menu.foods) {
          await fetch(`http://localhost:8000/foods/${food.food_id}`, {
            method: "DELETE",
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
        }
      }

      // Luego eliminar el menú
      const response = await fetch(`http://localhost:8000/menus/${menuId}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error("Error eliminando menú");
      }

      // Actualizar la lista de menús
      setMenus(menus.filter((menu) => menu.menu_id !== menuId));
      alert("Menú eliminado exitosamente");
    } catch (error) {
      console.error("Error al eliminar menú:", error);
      alert("Error al eliminar el menú");
    }
  };

  if (loading) {
    return (
      <div>
        <Navbar />
        <div className="menu-container">
          <p className="laoding-state">Cargando menús...</p>
        </div>
      </div>
    );
  }

  return (
    <div>
      <Navbar />
      <div className="menu-container">
        <div className="menu-header">
          <h1>Menús</h1>
          {adminUser && (
            <button
              className="btn-create-menu"
              onClick={() => setShowModal(true)}
            >
              Crear Menú
            </button>
          )}
        </div>

        {menus.length === 0 ? (
          <p>No hay menús disponibles</p>
        ) : (
          <div className="menus-grid">
            {menus.map((menu) => (
              <div key={menu.menu_id} className="menu-card">
                <div className="menu-card-header">
                  <h2>{menu.name}</h2>
                  {adminUser && (
                    <button
                      className="btn-delete-menu"
                      onClick={() => handleDeleteMenu(menu.menu_id)}
                    >
                      ✕
                    </button>
                  )}
                </div>

                <div className="foods-list">
                  {menu.foods && menu.foods.length > 0 ? (
                    menu.foods.map((food: Food) => (
                      <Card
                        key={food.food_id}
                        name={food.name}
                        description={food.description}
                        price={food.price}
                      />
                    ))
                  ) : (
                    <p className="no-foods">No hay alimentos en este menú</p>
                  )}
                </div>
              </div>
            ))}
          </div>
        )}

        {/* Modal para crear menú */}
        {showModal && (
          <div className="modal-overlay" onClick={() => setShowModal(false)}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
              <div className="modal-header">
                <h2>Crear Nuevo Menú</h2>
                <button
                  className="modal-close"
                  onClick={() => setShowModal(false)}
                >
                  ✕
                </button>
              </div>

              <form onSubmit={handleCreateMenu} className="modal-form">
                <div className="form-group">
                  <label htmlFor="menu-name">Nombre del Menú</label>
                  <input
                    id="menu-name"
                    type="text"
                    placeholder="Ej: Desayuno, Almuerzo, Cena"
                    value={newMenuName}
                    onChange={handleMenuNameChange}
                    required
                  />
                </div>

                <div className="form-section">
                  <div className="section-header">
                    <h3>Alimentos</h3>
                    <button
                      type="button"
                      className="btn-add-food"
                      onClick={addFoodField}
                    >
                      + Agregar Alimento
                    </button>
                  </div>

                  <div className="foods-form">
                    {foods.map((food, index) => (
                      <div key={index} className="food-form-group">
                        <input
                          type="text"
                          placeholder="Nombre del alimento"
                          value={food.name}
                          onChange={(e) =>
                            handleFoodChange(index, "name", e.target.value)
                          }
                          required
                        />
                        <input
                          type="text"
                          placeholder="Descripción"
                          value={food.description}
                          onChange={(e) =>
                            handleFoodChange(
                              index,
                              "description",
                              e.target.value
                            )
                          }
                        />
                        <input
                          type="number"
                          placeholder="Precio"
                          step="0.01"
                          value={food.price}
                          onChange={(e) =>
                            handleFoodChange(index, "price", e.target.value)
                          }
                          required
                        />
                        {foods.length > 1 && (
                          <button
                            type="button"
                            className="btn-remove-food"
                            onClick={() => removeFoodField(index)}
                          >
                            ✕
                          </button>
                        )}
                      </div>
                    ))}
                  </div>
                </div>

                <div className="modal-actions">
                  <button
                    type="button"
                    className="btn-cancel"
                    onClick={() => setShowModal(false)}
                  >
                    Cancelar
                  </button>
                  <button type="submit" className="btn-submit">
                    Crear Menú
                  </button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Menus;