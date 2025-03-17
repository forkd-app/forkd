type Environment = {
  NODE_ENV: typeof process.env.NODE_ENV
  BACKEND_URL: string
  SESSION_SECRETS: string[]
  SESSION_COOKIE_NAME: string
}

export const environment: Environment = {
  NODE_ENV: process.env.NODE_ENV,
  BACKEND_URL: process.env.BACKEND_URL ?? "http://api:8000/query",
  SESSION_SECRETS: process.env.SESSION_SECRETS?.split(",")?.map((s) =>
    s.trim()
  ) ?? ["super-secret"],
  SESSION_COOKIE_NAME: process.env.SESSION_COOKIE_NAME ?? "__sesh",
}
