"use client";
import { Message } from "../models/Message";

interface ChatWindowProps {
    messages: Message[];
    username: string;
}

export default function ChatWindow({ messages, username }: ChatWindowProps) {
    return (
        <main className="flex-1 p-6 overflow-y-auto">
            <div className="bg-white rounded-lg shadow-md p-4 min-h-[60vh] flex flex-col gap-3">
                {messages.length === 0 ? (
                    <p className="text-gray-500 text-center mt-20">
                        No messages yet. Start the conversation below!
                    </p>
                ) : (
                    messages.map((message) => (
                        <div
                            key={message.id ?? Math.random().toString()}
                            className={`flex ${message.from.username === username ? "justify-end" : "justify-start"
                                }`}
                        >
                            <div
                                className={`max-w-[70%] rounded-lg px-4 py-2 ${message.from.username === username
                                        ? "bg-indigo-600 text-white"
                                        : "bg-gray-200 text-gray-800"
                                    }`}
                            >
                                <p className={`text-sm font-semibold mb-1 ${message.from.username === username ? "text-white/80" : "text-gray-700"}`}>
                                    {message.from.username}
                                </p>
                                <p>{message.text}</p>
                                <span className="text-xs opacity-75">
                                    {new Date(message.time).toISOString().slice(0, 16).replace("T", " ")}
                                </span>
                            </div>
                        </div>
                    ))
                )}
            </div>
        </main>
    );
}