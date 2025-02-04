import type { MetaFunction } from "@remix-run/node"
import { Text, Flex, Container } from "@mantine/core"

export const meta: MetaFunction = () => {
  return [
    { title: "Log In" },
    {
      name: "join in on the fun and start creating recipes",
      content: "Welcome to Forkd!",
    },
  ]
}

export default function LogIn() {
  return (
    <Container
    >
    <Flex
      justify="center"
      align="center"
      style={styles.contain}
      direction={"column"}
    >
      <Text> Log In Page </Text>
    </Flex>
    </Container>
  )
}

const styles = {
  contain: {
    height: "100vh",
  },
}
