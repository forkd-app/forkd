import {
  ActionFunctionArgs,
  createCookieSessionStorage,
  data,
  LoaderFunctionArgs,
} from "react-router"
import { environment } from "app/.server/env"

export const cookieSession = createCookieSessionStorage<
  { sessionToken: string },
  { magicLinkToken: string }
>({
  cookie: {
    name: environment.SESSION_COOKIE_NAME,
    httpOnly: true,
    secure: environment.NODE_ENV === "production",
    secrets: environment.SESSION_SECRETS,
  },
})

async function getSessionFromRequest(req: Request) {
  return cookieSession.getSession(req.headers.get("Cookie"))
}

/**
 * Wraps the session to reduce boilerplate when writing actions and loaders that need access to the session
 * Also includes a simple default for when the session is required
 */
export async function getSessionOrThrow(
  args: ActionFunctionArgs | LoaderFunctionArgs,
  required?: boolean
): ReturnType<typeof getSessionFromRequest> {
  const session = await getSessionFromRequest(args.request)
  if (required && !session.has("sessionToken")) {
    throw data(
      {
        error: "unauthorized",
      },
      {
        status: 401,
      }
    )
  }
  return session
}
