type Command =
  | 'message'
  | 'join'
  | 'leave'
  | 'list-users'
  | 'list-rooms';

interface ClientInfo {
  username: string;
  id: string;
}

export interface Message {
  id?: string;
  from: ClientInfo;
  text: string;
  room: string;
  type: Command;
  time: string;
}