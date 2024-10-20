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

export interface LoadedProfile extends Message {
  profile: {
    title: string;
    name: string;
    image: string;
    description: string;
    technologies: string[];
    link: string;
  };
}
