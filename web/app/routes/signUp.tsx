import type { MetaFunction } from "@remix-run/node"
import { Text, Flex } from "@mantine/core"

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
    <Flex
      style={styles.contain}
      justify="center"
      align="center"
      direction={"column"}
    >
      <Text> Sign Up Page </Text>
    </Flex>
  )
}

const styles = {
  contain: {
    height: "100vh",
    width: "100%",
  },
}
