import Account from './dtos/account';

export default async function createAccount(account: Account) {
  const accountURL = process.env.NEXT_PUBLIC_VISION_ACCOUNT;

  const response = await fetch(`${accountURL}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: account.toJSON(),
  });

  const data = await response.json();
  console.log(data);
}
