import { PostStampBody } from './model';

export async function postStamp(url: string, body: PostStampBody) {
	const res = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(body)
	});

	const json = await res.json();
	result = JSON.stringify(json);

	return result;
}
