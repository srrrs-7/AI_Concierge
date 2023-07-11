const url = 'http://localhost:8880';

export async function get() {
  try {
    const result = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' }
    });
    return new Response(JSON.stringify(result), {
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
