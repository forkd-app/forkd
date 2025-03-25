import { redirect } from "@remix-run/react"
import { getSDK } from "~/gql/client"
import { cookieSession, getSessionOrThrow } from "~/.server/session"
import { LoaderFunctionArgs } from "@remix-run/node"
import { environment } from "~/.server/env"

export async function loader(args: LoaderFunctionArgs) {
  const session = await getSessionOrThrow(args, true)
  const token = session.get("sessionToken")
  if (token) {
    const sdk = getSDK(environment.BACKEND_URL, token)
    const res = await sdk.Logout()

    if (res.user?.logout) {
      return redirect("/", {
        headers: {
          "Set-Cookie": await cookieSession.destroySession(session),
        },
      })
    }
  }
  return redirect("/auth/login")
}
