import { Container, Text, Button, Input } from "@mantine/core"
import { useState } from "react"
import { IconMail } from "@tabler/icons-react"
import { Form } from "@remix-run/react"

export default function SignUp() {
  const [email, setEmail] = useState<string>("")

  const handleSignUp = () => {}

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
      <Form style={{ width: 300 }} onSubmit={handleSignUp}>
        <Text style={{ marginBottom: 10 }}>SIGN UP</Text>
        <Input
          type="email"
          placeholder="Your email"
          leftSection={<IconMail size={16} />}
          value={email}
          onChange={(event) => setEmail(event.currentTarget.value)}
        />

        <Button style={{ width: "100%", marginTop: 15 }} type="submit">
          sign up
        </Button>
      </Form>
    </Container>
  )
}
