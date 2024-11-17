import Account from './dtos/requests/account';
import Credentials from './dtos/requests/credentials';
import {
  Message, DefaultResponse, Token, Me,
} from './dtos/responses/default-response';

export async function createAccount(account: Account): Promise<DefaultResponse<Message>> {
  const accountURL = process.env.NEXT_PUBLIC_VISION_ACCOUNT;

  const response = await fetch(`${accountURL}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: account.toJSON(),
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

export async function loginAccount(account: Credentials): Promise<DefaultResponse<Token>> {
  const accountURL = process.env.NEXT_PUBLIC_VISION_ACCOUNT;

  const response = await fetch(`${accountURL}/auth`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: account.toJSON(),
  });

  let data;
  try {
    data = await response.json() as Token;
  } catch (error) {
    data = null;
  }

  return {
    ok: response.ok,
    status: response.status,
    data,
  };
}

export async function loadMe(token: string): Promise<DefaultResponse<Me>> {
  const accountURL = process.env.NEXT_PUBLIC_VISION_ACCOUNT;

  const response = await fetch(`${accountURL}/me`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: token,
    },
  });

  let data;
  try {
    data = await response.json() as Me;
  } catch (error) {
    data = null;
  }

  return {
    ok: response.ok,
    status: response.status,
    data,
  };
}
