import { request } from 'playwright';
import { env } from 'src/domain/calendar/util/env';

const host = 'http://localhost:8880';

export async function post({ request }: any) {
  try {
    const body = await request.json();
    const url = `${host}'/calendar'`;
    const res = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    });
    return new Response(JSON.stringify(res), {
      status: 200,
      headers: { 'Content-Type': 'application/json' }
    });
  } catch (error) {
    return new Response(JSON.stringify(error), {
      status: 404,
      headers: { 'Content-Type': 'application/json' }
    });
  }
}
