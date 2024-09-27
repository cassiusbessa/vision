export function setTokenLocalStorage(token: string) {
  localStorage.setItem('@token', token);
}

export function setTokenSessionStorage(token: string) {
  sessionStorage.setItem('@token', token);
}

export function getToken() {
  return localStorage.getItem('@token') || sessionStorage.getItem('@token');
}

export function removeTokenLocalStorage() {
  localStorage.removeItem('@token');
}

export function removeTokenSessionStorage() {
  sessionStorage.removeItem('@token');
}
