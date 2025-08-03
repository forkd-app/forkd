import { Center, Text } from "@mantine/core"
import { IconHeartFilled, IconBellFilled } from "@tabler/icons-react"

export function Footer() {
  return (
    <Center>
      <Text size="1.25rem">
        Made with <IconHeartFilled size={"1rem"} color="red" /> in Philly{" "}
        <IconBellFilled size={"1rem"} color="gold" />
      </Text>
    </Center>
  )
}
