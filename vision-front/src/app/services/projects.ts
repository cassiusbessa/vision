/* eslint-disable import/prefer-default-export */
import { DefaultResponse, LoadedProject } from './dtos/responses/default-response';

export async function loadProjectsByProfileId(id: string):
Promise<DefaultResponse<LoadedProject[]>> {
  const profileURL = process.env.NEXT_PUBLIC_VISION_PROJECT;

  const response = await fetch(`${profileURL}/profile/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  let data;
  try {
    data = await response.json() as LoadedProject[];
  } catch (error) {
    data = null;
  }

  return {
    ok: response.ok,
    status: response.status,
    data,
  };
}
