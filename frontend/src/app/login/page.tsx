
import { redirect } from "next/navigation";

async function handleLogin(formData: FormData) {
  'use server'
  const username = formData.get("username")?.toString() || "";
  const room = formData.get("room")?.toString() || "";

  if (!username || !room) {
    throw new Error("Missing username or room");
  }

  redirect(`/chat?username=${encodeURIComponent(username)}&room=${encodeURIComponent(room)}`);
}

export default function LoginPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <form
        action={handleLogin}
        className="bg-white rounded-xl shadow-lg p-8 w-full max-w-md flex flex-col gap-6"
      >
        <h1 className="text-2xl font-bold text-gray-800 text-center">
          Join Chat Room
        </h1>

        <div className="flex flex-col gap-4">
          <input
            type="text"
            name="username"
            placeholder="Enter your username"
            required
            className="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-gray-400"
          />
          <input
            type="text"
            name="room"
            placeholder="Enter room name"
            required
            className="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-gray-400"
          />
        </div>

        <button
          type="submit"
          className="w-full bg-indigo-600 text-white font-semibold py-2 rounded-lg hover:bg-indigo-700 transition duration-200"
        >
          Join Room
        </button>
      </form>
    </div>
  );
}