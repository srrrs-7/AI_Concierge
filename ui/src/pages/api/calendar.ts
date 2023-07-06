import { env } from '../../util/env';

export async function get({ params }) {
  const result = await fetch(`${HOST}/user`);
  const data = result.json();
  return new Response(data, { status: 200 });
}

export async function post({ request }) {
  if (request.headers.get('Content-Type') === 'application/json') {
    const body = await request.json();
    const url = env.CALENDAR_API_ENDPOINT + '/calendar';
    const result = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    });
    const json = await result.json();
    return new Response(json, { status: 200 });
  }
  return new Response('invalid content-type', { status: 404 });
}
