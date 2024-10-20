export interface DefaultResponse<T> {
  ok: boolean;
  status: number;
  data: T | null;
}

export interface Message {
  message: string;
}

export interface Token extends Message {
  token: string;
}
