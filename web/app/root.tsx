import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
  LoaderFunctionArgs,
} from "react-router"
import { MantineProvider } from "@mantine/core"
import "@mantine/core/styles.css"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { environment } from "~/.server/env"
import { getStore } from "~/stores/global"
import { Provider } from "react-redux"
import { ClientError } from "graphql-request"
import { useMemo } from "react"

export async function loader(args: LoaderFunctionArgs) {
  const session = await getSessionOrThrow(args, false)
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
  const data = useLoaderData<typeof loader>()
  const store = useMemo(() => {
    return getStore({ user: { value: data } })
  }, [data])
  console.log(data)

  return (
    <html lang="en" data-mantine-color-scheme="light">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        <MantineProvider>
          <Provider store={store}>{children} </Provider>
        </MantineProvider>
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  )
}

export default function App() {
  return <Outlet />
}
