const host = 'http://localhost:8880';
export function getUrl(): string {
  return host + '/register';
}

export async function postBasicAuth(url: string, body: any) {
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
