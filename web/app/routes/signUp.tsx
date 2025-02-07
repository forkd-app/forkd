import type { MetaFunction } from "@remix-run/node"
import { Text, Flex, Container } from "@mantine/core"

export const meta: MetaFunction = () => {
  return [
    { title: "Sign Up" },
    {
      name: "join in on the fun and start creating recipes",
      content: "Welcome to Forkd!",
    },
  ]
}

export default function SignUp() {
  return (
    <Container>
      <Flex
        style={styles.contain}
        justify="center"
        align="center"
        direction={"column"}
      >
        <Text> Sign Up Page </Text>
      </Flex>
    </Container>
  )
}

const styles = {
  contain: {
    height: "100vh",
  },
}
