import { Header } from "./header"
import { Footer } from "./footer"
import { Outlet } from "@remix-run/react"

export default function Index() {
  return (
    <>
      <Header />
      <Outlet />
      <Footer />
    </>
  )
}
