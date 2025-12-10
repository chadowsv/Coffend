import React, { useEffect, useState } from "react";
import Navbar from "../components/Navbar";
import "../styles/mesas.css";

interface Table {
  table_id: number;
  number_guests: number;
  status: boolean;
  created_at: string;
  updated_at: string;
}

export default function Mesas() {
  const [tables, setTables] = useState<Table[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchTables = async () => {
      setLoading(true);
      setError("");

      try {
        const token = localStorage.getItem("token");
        if (!token) {
          setError("No hay token disponible. Por favor, inicia sesión.");
          setLoading(false);
          return;
        }

        const response = await fetch("http://localhost:8000/tables", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.status === 401) {
          setError("No autorizado. Token inválido o expirado.");
          setLoading(false);
          return;
        }

        const data = await response.json();
        setTables(Array.isArray(data) ? data : []);
      } catch (err) {
        console.error(err);
        setError("Error al conectarse con el servidor.");
      } finally {
        setLoading(false);
      }
    };

    fetchTables();
  }, []);

  const toggleTableStatus = async (table: Table) => {
    const token = localStorage.getItem("token");
    if (!token) return;

    try {
      const response = await fetch(`http://localhost:8000/tables/${table.table_id}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ status: !table.status }),
      });

      if (response.ok) {
        const updatedTable = await response.json();
        setTables(prev =>
          prev.map(t => (t.table_id === table.table_id ? updatedTable : t))
        );
      } else {
        console.error("Error al actualizar la mesa");
      }
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <Navbar />
      <div className="mesas-container">
        <h1 className="title">Mesas</h1>

        {loading && <p>Cargando mesas...</p>}
        {error && <p style={{ color: "red" }}>{error}</p>}
        {!loading && !error && tables.length === 0 && <p>No hay mesas disponibles.</p>}

        <div className="grid">
          {tables.map((table) => (
            <div
              key={table.table_id}
              className={`mesa-card ${table.status ? "ocupada" : "libre"}`}
              onClick={() => toggleTableStatus(table)}
              style={{ cursor: "pointer" }}
            >
              <h3>Mesa #{table.table_id}</h3>

              <div className="mesa-vista">
                <div className="chair chair-top"></div>
                <div className="mesa-middle">
                  <div className="chair chair-left"></div>
                  <div className="table"></div>
                  <div className="chair chair-right"></div>
                </div>
                <div className="chair chair-bottom"></div>
              </div>

              <p>
                <strong>Capacidad:</strong>{" "}
                <span className="cap-num">{table.number_guests} personas</span>
              </p>

              <p>
                <strong>Estado:</strong>{" "}
                {table.status ? (
                  <span className="estado-ocupado">Ocupada</span>
                ) : (
                  <span className="estado-libre">Libre</span>
                )}
              </p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
