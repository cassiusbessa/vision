import Account from './dtos/account';
import Credentials from './dtos/credentials';

export async function createAccount(account: Account) {
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

export async function loginAccount(account: Credentials) {
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
