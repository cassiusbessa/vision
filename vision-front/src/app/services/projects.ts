import Project from './dtos/requests/project';
import { DefaultResponse, LoadedProjects, Message } from './dtos/responses/default-response';
import { getToken } from './token';

export async function loadProjectsByProfileId(id: string):
Promise<DefaultResponse<LoadedProjects>> {
  const profileURL = process.env.NEXT_PUBLIC_VISION_PROJECT;

  const response = await fetch(`${profileURL}/profile/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  let data;
  try {
    data = await response.json() as LoadedProjects;
  } catch (error) {
    data = null;
  }

  return {
    ok: response.ok,
    status: response.status,
    data,
  };
}

export async function createProject(project: Project): Promise<DefaultResponse<Message>> {
  const projectURL = process.env.NEXT_PUBLIC_VISION_PROJECT;

  const response = await fetch(`${projectURL}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: getToken() || '',
    },
    body: project.toJSON(),
  });

  let data;
  try {
    data = await response.json() as Message;
  } catch (error) {
    data = null;
  }

  return {
    ok: response.ok,
    status: response.status,
    data,
  };
}
