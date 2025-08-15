import { useEffect, useState } from "react";
import { fetchJSON } from "./lib/api";

type Profile = { full_name: string; title?: string; email?: string; about?: string };

export default function App() {
  const [profile, setProfile] = useState<Profile | null>(null);
  useEffect(() => {
    fetchJSON<Profile>("/api/profile").then(setProfile).catch(console.error);
  }, []);

  return (
    <div className="mx-auto max-w-3xl p-6">
      <header className="flex items-center gap-4">
        <div className="h-16 w-16 rounded-full bg-gray-200" />
        <div>
          <h1 className="text-2xl font-bold">{profile?.full_name ?? "Your Name"}</h1>
          <p className="text-gray-600">{profile?.title ?? "Title"}</p>
        </div>
      </header>

      <section className="mt-8 space-y-2">
        <h2 className="text-xl font-semibold">About</h2>
        <p className="text-gray-700">{profile?.about ?? "Short bio goes here."}</p>
      </section>

      <section className="mt-8">
        <h2 className="text-xl font-semibold">Publications</h2>
        {/* Add list rendering by calling /api/publications */}
      </section>
    </div>
  );
}