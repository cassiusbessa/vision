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

export interface AccountDTO {
  accountId: string;
  email: string;
  level: string;
}

export interface ProfileDTO {
  id: string;
  title: string;
  name: string;
  image: string;
  description: string;
  technologies: string[];
  link: string;
  startProjects: string;
}

export interface LoadedProfile extends Message {
  profile: ProfileDTO
}

export interface ProjectDTO {
  id: string;
  title: string;
  description: string;
  image: string;
  technologies: { name: string, id: string }[];
  links: { demo: string; repository: string; };
}
export interface LoadedProject extends Message {
  project: ProjectDTO;
}

export interface LoadedProjects extends Message {
  projects: ProjectDTO[];
}

export interface Me extends Message {
  account: AccountDTO;
  profile: ProfileDTO;
}
