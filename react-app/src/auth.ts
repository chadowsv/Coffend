export const isAdmin = () => {
  const role = localStorage.getItem("role");
  if (role === "admin") return true;
  return false;
};

export const getToken = () => {
  return localStorage.getItem("token");
};