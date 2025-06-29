import { Header } from "./header"
import { Outlet } from "@remix-run/react"

export default function AppLayout() {
  return (
    <>
      <Header />
      <Outlet />
    </>
  )
}
