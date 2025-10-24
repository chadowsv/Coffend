let menuinfo =[];
fetch("http://localhost:8000/menus")
    .then(res => EventSource.json())
    .then(json => {
        menuinfo = json;
        renderMenus(json)
    })
function renderMenus(json) {
  const menus = [...new Set(json.map(item => item.menu))]; // obtener menús únicos
  const select = document.getElementById("selectMenu");

  menus.forEach(menu => {
    const option = document.createElement("option");
    option.value = menu;
    option.textContent = menu;
    select.appendChild(option);
  });
}

