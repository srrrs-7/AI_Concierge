export function getNow(): string {
	const now = new Date();

	const year: string = now.getFullYear();
	const month: string = now.getMonth() + 1;
	const date: string = now.getDate();

	const current = year + '-' + month + '-' + date; // e.g 2006-01-02
	return current;
}
