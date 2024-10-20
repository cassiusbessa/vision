/* eslint-disable import/prefer-default-export */

import { DefaultResponse } from './dtos/responses/default-response';
import { LoadedTag } from './dtos/responses/loaded-tag';

export async function loadTags(): Promise<DefaultResponse<LoadedTag[]>> {
  const tagsURL = process.env.NEXT_PUBLIC_VISION_TAG;

  const response = await fetch(`${tagsURL}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  let data;
  try {
    data = await response.json();
  } catch (error) {
    data = null;
  }

  return {
    ok: response.ok,
    status: response.status,
    data,
  };
}
