import { cookies } from "next/headers";
import { ClientInfo } from "../models/ClientInfo";

export async function saveClientInfo(info: ClientInfo) {
  const cookieStore = await cookies();
  cookieStore.set("clientInfo", JSON.stringify(info), {
    httpOnly: true, 
    sameSite: "lax",
    path: "/",
  });
}
export async function getClientInfo(): Promise<ClientInfo | null> {
  const cookieStore = await cookies();
  const cookie = cookieStore.get("clientInfo");
  if (cookie) {
    return JSON.parse(cookie.value);
  }
  return null;
}