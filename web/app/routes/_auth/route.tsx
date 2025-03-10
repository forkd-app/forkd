import type { MetaFunction } from "@remix-run/node"
import { Flex, Container } from "@mantine/core"
import { Outlet } from "@remix-run/react"

export const meta: MetaFunction = () => {
  return [
    { title: "Sign Up" },
    {
      name: "join in on the fun and start creating recipes",
      content: "Welcome to Forkd!",
    },
  ]
}

export default function AuthLayout() {
  return (
    <Container>
      <Flex
        style={styles.contain}
        justify="center"
        align="center"
        direction={"column"}
      >
        <Outlet />
      </Flex>
    </Container>
  )
}

const styles = {
  contain: {
    height: "100vh",
  },
}
