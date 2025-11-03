import { redirect } from "next/navigation";
import { getClientInfo } from "./actions/clientActions";

export default async function Home() {
  const clientInfo = await getClientInfo();
  if (!clientInfo) {
    redirect('/login');
  }

    return (
    <div>Welcome to the Chat Application!</div>
  );
}