const HOST = 'http://host.docker.internal:8880';

export async function getUser() {
  const result = await fetch(`${HOST}/user`);
  const data = result.json();
  return data;
}
