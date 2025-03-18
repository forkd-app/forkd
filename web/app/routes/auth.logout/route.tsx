import { ReactNode, useEffect } from "react"
import { redirect, useSearchParams, useSubmit } from "@remix-run/react"
import { getSDK } from '~/gql/client'
import { cookieSession, sessionWrapper } from "~/.server/session"


export const loader = sessionWrapper(async (_, session) => {

  try {
    const token = session.get("sessionToken")
   if (token) {
      const sdk = getSDK("http://localhost:8000/query", token)
      const res = await sdk.Logout()

      if (res.user?.logout) {
        return redirect("/", {
          headers: {
            "Set-Cookie": await cookieSession.destroySession(session),
          },
        })
      }
    }
  } catch (err) {
    console.error(err)
  }
  return redirect("/auth/login")
})
