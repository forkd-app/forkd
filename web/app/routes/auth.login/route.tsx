import { Container, Text, Button, TextInput } from "@mantine/core"
import { IconMail } from "@tabler/icons-react"
import {
  Form,
  useActionData,
  useNavigation,
  useSubmit,
  ActionFunctionArgs,
  redirect,
} from "react-router"
import { isEmail, useForm } from "@mantine/form"
import { object, string, email, pipe, parse } from "valibot"
import { client } from "~/gql/client"
import { cookieSession, getSessionOrThrow } from "~/.server/session"
import { ClientError } from "graphql-request"

const NOT_REGISTERED_ERROR = "email not registered"

const loginForm = object({
  email: pipe(string(), email()),
})

export async function action(args: ActionFunctionArgs) {
  const session = await getSessionOrThrow(args)
  const body = await args.request.json()
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
      // request current user obect from gpl client
    }
  } catch (err) {
    if (err instanceof ClientError) {
      // TODO: Maybe revisit this and make it a little more resilient...
      // Lol, this is kinda whack.
      const msg = err.message.split(":", 2).slice(0, 2).join(":")
      if (msg.startsWith(NOT_REGISTERED_ERROR)) {
        return { errors: [{ field: "email", error: msg }] }
      }
    }
  }
  return null
}

export default function LogIn() {
  const navigation = useNavigation()
  const isSubmitting =
    navigation.formAction === "/auth/signup" &&
    navigation.state === "submitting"
  const data = useActionData<typeof action>()
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
        onSubmit={form.onSubmit((values) => {
          submit(values, { method: "POST", encType: "application/json" })
        })}
      >
        <Text style={{ marginBottom: 10 }}>LOGIN</Text>
        <TextInput
          withAsterisk
          disabled={isSubmitting}
          type="email"
          label="email"
          placeholder="user@forkd.food"
          leftSection={<IconMail size={16} />}
          key={form.key("email")}
          {...form.getInputProps("email")}
          error={
            !!data &&
            "errors" in data &&
            data.errors.find((err) => err.field === "email")?.error
          }
        />
        <Button
          loading={isSubmitting}
          style={{ width: "100%", marginTop: 15 }}
          type="submit"
        >
          login
        </Button>
      </Form>
    </Container>
  )
}
