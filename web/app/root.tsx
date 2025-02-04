import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react"
import { MantineProvider, Image } from "@mantine/core"
import "@mantine/core/styles.css"
import { Header } from "./components/header/header"
import { MobileHeader } from "./components/header/mobileHeader"
import { useMediaQuery } from "@mantine/hooks"

export function Layout({ children }: { children: React.ReactNode }) {
  const isMobile = useMediaQuery("(max-width: 1199px)")
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        <MantineProvider>
          {isMobile ? <MobileHeader /> : <Header />}
          {children}
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
