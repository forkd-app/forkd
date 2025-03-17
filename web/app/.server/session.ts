import {
  ActionFunction,
  ActionFunctionArgs,
  createCookieSessionStorage,
  data,
  LoaderFunction,
  LoaderFunctionArgs,
} from "@remix-run/node"
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

export async function getSession(req: Request) {
  return cookieSession.getSession(req.headers.get("Cookie"))
}

/**
 * Wraps the session to reduce boilerplate when writing actions and loaders that need access to the session
 * Also includes a simple default for when the session is required
 */
export function sessionWrapper<
  T extends ActionFunctionArgs | LoaderFunctionArgs,
  U extends T extends ActionFunctionArgs ? ActionFunction : LoaderFunction,
>(
  cb: (
    args: T,
    session: Awaited<ReturnType<typeof getSession>>
  ) => ReturnType<U>,
  required?: boolean
): (args: T) => Promise<ReturnType<U>> {
  return async (args) => {
    const session = await getSession(args.request)
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
    return cb(args, session)
  }
}
