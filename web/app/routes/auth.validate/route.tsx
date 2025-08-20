import { ReactNode, useEffect } from "react"
import {
  redirect,
  useSearchParams,
  useSubmit,
  ActionFunctionArgs,
} from "react-router"
import { getSDK } from "~/gql/client"
import { cookieSession, getSessionOrThrow } from "~/.server/session"
import { environment } from "~/.server/env"

export async function action(args: ActionFunctionArgs) {
  const session = await getSessionOrThrow(args)
  const token = session.get("magicLinkToken")
  const code = await args.request.text()
  if (token && code) {
    const sdk = getSDK(environment.BACKEND_URL)
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
  return redirect("/auth/login")
}

export default function Validate(): ReactNode {
  const [searchParams] = useSearchParams()
  const code = searchParams.get("code")
  const submit = useSubmit()
  useEffect(() => {
    submit(code, { method: "POST", encType: "text/plain" })
  }, [code, submit])
  return null
}
