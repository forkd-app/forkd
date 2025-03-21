import { Container, Text, Button, TextInput } from "@mantine/core"
import { IconMail } from "@tabler/icons-react"
import { Form, useSubmit } from "@remix-run/react"
import { isEmail, useForm } from "@mantine/form"
import { redirect } from "@remix-run/node"
import { object, string, email, pipe, parse } from "valibot"
import { client } from "~/gql/client"
import { cookieSession, sessionWrapper } from "~/.server/session"

const loginForm = object({
  email: pipe(string(), email()),
})

export const action = sessionWrapper(async ({ request }, session) => {
  const body = await request.json()
  const formdata = parse(loginForm, body)
  try {
    const res = await client.RequestMagicLink(formdata)
    if (res.user?.requestMagicLink) {
      session.flash("magicLinkToken", res.user.requestMagicLink)
      return redirect("/auth/magic-link", {
        headers: {
          "Set-Cookie": await cookieSession.commitSession(session),
        },
      })
    }
  } catch (err) {
    console.error(err)
  }
  return null
})

export default function LogIn() {
  const form = useForm({
    name: "login",
    mode: "controlled",
    initialValues: {
      email: "",
    },
    validate: {
      email: isEmail("invalid email address"),
    },
  })

  const submit = useSubmit()

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
      <Form
        style={{ width: 300 }}
        method="POST"
        onSubmit={form.onSubmit((values) =>
          submit(values, { method: "POST", encType: "application/json" })
        )}
      >
        <Text style={{ marginBottom: 10 }}>LOGIN</Text>
        <TextInput
          withAsterisk
          type="email"
          label="email"
          placeholder="user@forkd.food"
          leftSection={<IconMail size={16} />}
          key={form.key("email")}
          {...form.getInputProps("email")}
        />
        <Button style={{ width: "100%", marginTop: 15 }} type="submit">
          login
        </Button>
      </Form>
    </Container>
  )
}
