import { Header } from "./_app/header"
import { Footer } from "./_app/footer"
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
