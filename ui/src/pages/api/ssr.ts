const HOST = 'http://host.docker.internal:3000';

type User = {
  id: number;
  name: string;
};

export async function getUser() {
  const result = await fetch(`${HOST}/user`);
  return result.json() as User;
}
