export const isAdmin = () => {
  const role = localStorage.getItem("role");
  if (role === "admin") return true;
  return false;
};

export const getToken = () => {
  return localStorage.getItem("token");
};

export const isLoggedIn = () => {
  const token = getToken();
  return token !== null;
};

export const logout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("user");
  localStorage.removeItem("role");
}