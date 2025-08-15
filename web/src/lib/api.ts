const API_BASE = import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";
export async function fetchJSON<T>(path: string): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`);
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  return res.json();
}