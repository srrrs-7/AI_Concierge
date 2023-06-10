type Env = {
	STAMP_API_ENDPOINT: string;
};

export const env: Env = {
	STAMP_API_ENDPOINT: process.env.STAMP_API_ENDPOINT
};
