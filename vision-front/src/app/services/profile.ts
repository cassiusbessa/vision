/* eslint-disable import/prefer-default-export */
import Profile from './dtos/requests/profile';
import { DefaultResponse, LoadedProfile, Message } from './dtos/responses/default-response';
import { getToken } from './token';

export async function createProfile(profile: Profile): Promise<DefaultResponse<Message>> {
  const profileURL = process.env.NEXT_PUBLIC_VISION_PROFILE;

  const response = await fetch(`${profileURL}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: getToken() || '',
    },
    body: profile.toJSON(),
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

export async function loadProfileByUrl(url: string): Promise<DefaultResponse<LoadedProfile>> {
  const profileURL = process.env.NEXT_PUBLIC_VISION_PROFILE;

  const response = await fetch(`${profileURL}/link/${url}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  let data;
  try {
    data = await response.json() as LoadedProfile;
  } catch (error) {
    data = null;
  }

  return {
    ok: response.ok,
    status: response.status,
    data,
  };
}
