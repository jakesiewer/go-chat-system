"use client";
import { useState, useEffect } from "react";
import ChatBanner from "./ChatBanner";
import ChatWindow from "./ChatWindow";
import ChatBar from "./ChatBar";
import { Message } from "../models/Message";

export default function ChatContainer({ username, room }: { username: string; room: string }) {
    const [messages, setMessages] = useState<Message[]>([]);
    const [ws, setWs] = useState<WebSocket | null>(null);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const socket = new WebSocket(`${process.env.NEXT_PUBLIC_CHAT_API_UR}?username=${username}&room=${room}`);
        setWs(socket);

        socket.onmessage = (event) => {
            const receivedMessage = JSON.parse(event.data);
            setMessages((prev) => [...prev, receivedMessage]);
        };

        return () => socket.close();
    }, []);

    const handleSendMessage = (message: string) => {
        const newMessage: Message = {
            from: { username, id: "" },
            text: message,
            room,
            type: "message",
            time: new Date().toISOString(),
        };

        if (!ws || ws.readyState !== WebSocket.OPEN) {
            setError("Failed to send message: connection not open.");
            return;
        }

        try {
            ws.send(newMessage.text);
            setError(null);
        } catch (err) {
            console.error("Error sending message:", err);
            setError("Failed to send message.");
        }

        setError(null);
    };

    return (
        <div className="min-h-screen flex flex-col bg-gray-100">
            {error && <div className="bg-red-500 text-white p-2 text-center">{error}</div>}
            <ChatBanner username={username} room={room} />
            <ChatWindow messages={messages} username={username} />
            <ChatBar onSendMessage={handleSendMessage} />
        </div>
    );
}