import { Flex, Container } from "@mantine/core"
import { Outlet } from "react-router"

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
