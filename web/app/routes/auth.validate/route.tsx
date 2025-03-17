import { ReactNode, useEffect } from "react"
import { redirect, useSearchParams, useSubmit } from "@remix-run/react"
import { getSDK } from "~/gql/client"
import { cookieSession, sessionWrapper } from "~/.server/session"

export const action = sessionWrapper(async ({ request }, session) => {
  try {
    const token = session.get("magicLinkToken")
    const code = await request.text()
    if (token && code) {
      const sdk = getSDK("http://localhost:8000/query")
      const res = await sdk.Login({
        token,
        code,
      })

      if (res.user?.login.token) {
        session.set("sessionToken", res.user.login.token)
        return redirect("/", {
          headers: {
            "Set-Cookie": await cookieSession.commitSession(session),
          },
        })
      }
    }
  } catch (err) {
    console.error(err)
  }
  return redirect("/auth/login")
})

export default function Validate(): ReactNode {
  const [searchParams] = useSearchParams()
  const code = searchParams.get("code")
  const submit = useSubmit()
  useEffect(() => {
    submit(code, { method: "POST", encType: "text/plain" })
  }, [code, submit])
  return null
}
