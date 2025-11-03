"use client";

interface ChatBannerProps {
  username: string;
  room: string;
}

export default function ChatBanner({ username, room }: ChatBannerProps) {
  return (
    <header className="bg-indigo-600 text-white py-4 px-6 shadow-md flex justify-between items-center">
      <h1 className="text-xl font-semibold">Room: {room}</h1>
      <span className="text-sm">Logged in as: {username}</span>
    </header>
  );
}