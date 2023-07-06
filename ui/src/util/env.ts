type Env = {
  CALENDAR_API_ENDPOINT: string | undefined;
};

export const env: Env = {
  CALENDAR_API_ENDPOINT: process.env.CALENDAR_API_ENDPOINT
};
