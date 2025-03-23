import { Header } from "./header"
import { Footer } from "./footer"
import { Outlet } from "@remix-run/react"
import { getSessionOrThrow } from "~/.server/session"
import { LoaderFunctionArgs } from "@remix-run/node"

export async function loader(args: LoaderFunctionArgs) {
  const session = await getSessionOrThrow(args)
  console.log("Session Token: ", session.get("sessionToken"))
  return null
}

export default function AppLayout() {
  return (
    <>
      <Header />
      <Outlet />
      <Footer />
    </>
  )
}
