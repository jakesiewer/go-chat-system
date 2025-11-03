import { redirect } from "next/navigation";
import ChatContainer from "../components/ChatContainer";

export default async function ChatPage({ searchParams }: { searchParams: Promise<{ username?: string; room?: string }> }) {
  const params = await searchParams;
  const username = params.username;
  const room = params.room;

  if (!username || !room) {
    redirect("/login");
  }

  return <ChatContainer username={username} room={room} />;
}