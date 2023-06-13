const host = 'http://localhost:8880';
export function geBasicAuthUrl(): string {
  return host + '/register';
}

type BasicAuth = {
  id: string;
  password: string;
};
export async function postBasicAuth(url: string, body: BasicAuth) {
  const res = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(body)
  });

  const json = await res.json();
  const result = JSON.stringify(json);

  return result;
}
