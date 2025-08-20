import { Header } from "./header"
import { Footer } from "./footer"
import { Outlet } from "react-router"

export default function AppLayout() {
  return (
    <>
      <Header />
      <Outlet />
      <Footer />
    </>
  )
}
