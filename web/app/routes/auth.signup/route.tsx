import { Container, Text, Button, TextInput } from "@mantine/core"
import { IconMail } from "@tabler/icons-react"
import { Form, useActionData, useNavigation, useSubmit } from "@remix-run/react"
import { isEmail, isNotEmpty, useForm } from "@mantine/form"
import { ActionFunctionArgs, redirect } from "@remix-run/node"
import { object, string, email, pipe, parse, InferOutput } from "valibot"
import { client } from "~/gql/client"
import { cookieSession, getSessionOrThrow } from "~/.server/session"

const signupForm = object({
  email: pipe(string(), email()),
  displayName: pipe(string()),
})

type SignupForm = InferOutput<typeof signupForm>

export async function action(args: ActionFunctionArgs) {
  const session = await getSessionOrThrow(args, false)
  const body = await args.request.json()
  const formdata = parse(signupForm, body)
  const userExists = await client.CheckUserSignup(formdata)
  const errors: { field: keyof SignupForm; error: string }[] = []

  if (userExists.user?.byEmail?.email) {
    errors.push({
      field: "email",
      error: `Email already in use: ${formdata.email}`,
    })
  }

  if (userExists.user?.byDisplayName?.displayName) {
    errors.push({
      field: "displayName",
      error: `Display Name already in use: ${formdata.displayName}`,
    })
  }

  if (errors.length > 0) {
    return { errors }
  }

  const res = await client.Signup(formdata)
  if (res.user?.signup) {
    session.flash("magicLinkToken", res.user.signup)
    return redirect("/auth/magic-link", {
      headers: {
        "Set-Cookie": await cookieSession.commitSession(session),
      },
    })
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
    name: "signup",
    mode: "controlled",
    initialValues: {
      email: "",
      displayName: "",
    },
    validate: {
      email: isEmail("invalid email address"),
      displayName: isNotEmpty(),
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
        <Text style={{ marginBottom: 10 }}>SIGNUP</Text>
        <TextInput
          disabled={isSubmitting}
          withAsterisk
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
        <TextInput
          disabled={isSubmitting}
          withAsterisk
          type="text"
          label="display name"
          placeholder="user"
          leftSection={<IconMail size={16} />}
          key={form.key("displayName")}
          {...form.getInputProps("displayName")}
          error={
            !!data &&
            "errors" in data &&
            data.errors.find((err) => err.field === "displayName")?.error
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
