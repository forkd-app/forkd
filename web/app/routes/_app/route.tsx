import { Header } from "./header"
import { Footer } from "./footer"
import { Outlet } from "@remix-run/react"
import { sessionWrapper } from "~/.server/session"

export const loader = sessionWrapper((_args, session) => {
  console.log("Session Token: ", session.get("sessionToken"))
  return null
})

export default function AppLayout() {
  return (
    <>
      <Header />
      <Outlet />
      <Footer />
    </>
  )
}
