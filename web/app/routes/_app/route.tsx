import { Header } from "./header"
import { Footer } from "./footer"
import { Outlet } from "@remix-run/react"

export default function AppLayout() {
  return (
    <>
      <Header />
      <Outlet />
      <Footer />
    </>
  )
}
