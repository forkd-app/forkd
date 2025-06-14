import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
} from "@remix-run/react"
import { MantineProvider } from "@mantine/core"
import "@mantine/core/styles.css"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { environment } from "~/.server/env"
import { useGlobals } from "~/stores/global"
import { LoaderFunctionArgs } from "@remix-run/node"
import { ClientError } from "graphql-request"

export async function loader(args: LoaderFunctionArgs) {
  const session = await getSessionOrThrow(args, false)
  console.log("Session Token: ", session.get("sessionToken"))
  const auth = session.get("sessionToken")
  const sdk = getSDK(`${environment.BACKEND_URL}`, auth)
  try {
    const data = await sdk.CurrentUser()
    return data?.user?.current ?? null
  } catch (err) {
    if (err instanceof ClientError && err.message.includes("missing auth")) {
      return null
    }
    throw err
  }
}

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" data-mantine-color-scheme="light">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        <MantineProvider>{children}</MantineProvider>
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  )
}

export default function App() {
  const data = useLoaderData<typeof loader>()
  useGlobals.getInitialState().setUser(data)
  console.log(useGlobals.getState().user, "user signing in")

  return <Outlet />
}
