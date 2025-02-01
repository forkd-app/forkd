import { Flex, Anchor } from "@mantine/core"

export function Categories() {
  return (
    <>
      <Flex
        style={styles.gridContainer}
        justify="center"
        align="center"
        gap="xl"
      >
        <Anchor c="gray" style={styles.link} href="#link1">
          Text 1
        </Anchor>
        <Anchor c="gray" style={styles.link} href="#link2">
          Text 2
        </Anchor>
        <Anchor c="gray" style={styles.link} href="#link3">
          Text 3
        </Anchor>
        <Anchor c="gray" style={styles.link} href="#link4">
          Text 4
        </Anchor>
        <Anchor c="gray" style={styles.link} href="#link5">
          Text 5
        </Anchor>
        <Anchor c="gray" style={styles.link} href="#link6">
          Text 6
        </Anchor>
        <Anchor c="gray" style={styles.link} href="#dropdown">
          Text 7
        </Anchor>
      </Flex>
    </>
  )
}

const styles: Record<string, React.CSSProperties> = {
  link: {
    textAlign: "center",
  },
  gridContainer: {
    flexWrap: "wrap",
    padding: 20,
  },
}
