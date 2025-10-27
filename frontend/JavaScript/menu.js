// script.js
const container = document.getElementById('menus-container');

// URL de tu API
const API_URL = 'http://localhost:8080/menu';

// Función para traer los menús
async function fetchMenus() {
  try {
    const response = await fetch(API_URL);
    const menus = await response.json();
    
    // Limpiar container
    container.innerHTML = '';

    // Recorrer menús y crear HTML
    menus.forEach(menu => {
      const menuDiv = document.createElement('div');
      menuDiv.innerHTML = `
        <h2>${menu.menuname} (${menu.menu_status ? 'Activo' : 'Inactivo'})</h2>
        <ul>
          ${menu.foods_details.map(food => `<li>${food.name} - $${food.price}</li>`).join('')}
        </ul>
      `;
      container.appendChild(menuDiv);
    });

  } catch (err) {
    console.error('Error fetching menus:', err);
  }
}

// Ejecutar la función
fetchMenus();


