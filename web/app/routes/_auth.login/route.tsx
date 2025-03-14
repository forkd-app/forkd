import { Container, Text, Input, Button } from "@mantine/core"
import { IconMail } from "@tabler/icons-react"
import { Form } from "@remix-run/react"
import { useState } from "react"

export default function LogIn() {
  const [email, setEmail] = useState<string>("")

  const handleLogIn = () => {}

  return (
    <Container
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
        flexDirection: "column",
      }}
      fluid
    >
      <Form style={{ width: 300 }} onSubmit={handleLogIn}>
        <Text style={{ marginBottom: 10 }}>LOGIN</Text>
        <Input
          type="email"
          placeholder="Your email"
          leftSection={<IconMail size={16} />}
          value={email}
          onChange={(event) => setEmail(event.currentTarget.value)}
        />

        <Button style={{ width: "100%", marginTop: 15 }} type="submit">
          login
        </Button>
      </Form>
    </Container>
  )
}
