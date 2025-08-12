import { useEffect, useState } from "react"
import { useDisclosure } from "@mantine/hooks"
import {
  Grid,
  Flex,
  Autocomplete,
  Popover,
  Button,
  Text,
  Burger,
} from "@mantine/core"
import {
  IconChevronCompactDown,
  IconSearch,
  IconGrillFork,
} from "@tabler/icons-react"
import { Navigation } from "~/components/navigation/navigation"

export function MobileHeader() {
  const [opened, { toggle }] = useDisclosure()
  const [isVisible, setIsVisible] = useState(false)

  useEffect(() => {
    if (isVisible) {
      document.body.style.overflow = "hidden" // Disable scrolling
    } else {
      document.body.style.overflow = "" // Re-enable scrolling
    }

    // Clean up on component unmount
    return () => {
      document.body.style.overflow = ""
    }
  }, [isVisible])

  return (
    <Flex
      style={styles.flex}
      direction="column"
      justify="center"
      align="stretch"
    >
      <Grid style={styles.mobileGrid} justify="space-around">
        <Grid.Col
          style={[styles.mobileGrid, { justifyContent: "left" }]}
          span={{ base: 6, md: 6, lg: 6 }}
        >
          <IconGrillFork style={{ paddingRight: 5 }} />
          <Text>FORKD</Text>
        </Grid.Col>
        <Grid.Col
          style={[styles.mobileGrid, { justifyContent: "right" }]}
          span={{ base: 6, md: 6, lg: 6 }}
        >
          <Burger
            opened={opened}
            lineSize={3}
            onClick={() => {
              toggle()
              setIsVisible(!isVisible)
            }}
            aria-label="Toggle navigation"
          />
        </Grid.Col>
      </Grid>
      {isVisible && (
        <div style={styles.navContainer}>
          <Navigation />
        </div>
      )}
      <Autocomplete
        style={styles.autocomplete}
        placeholder="Search for a recipe"
        data={[
          "Vegan Mac and Cheese",
          "Chickpea Salad",
          "Peanut Butter Cookies",
          "Avocado Toast",
        ]}
        leftSection={<IconSearch size={16} stroke={1.5} />}
        rightSection={
          <Popover width={200} position="bottom" withArrow shadow="md">
            <Popover.Target>
              <IconChevronCompactDown />
            </Popover.Target>
            <Popover.Dropdown>
              <Flex direction="column">
                <Button variant="transparent" color="gray" size="xs">
                  Recipe Title
                </Button>
                <Button variant="transparent" color="gray" size="xs">
                  Author
                </Button>
                <Button variant="transparent" color="gray" size="xs">
                  # of Forks
                </Button>
                <Button variant="transparent" color="gray" size="xs">
                  Publish Date
                </Button>
              </Flex>
            </Popover.Dropdown>
          </Popover>
        }
      />
    </Flex>
  )
}

const styles: Record<string, React.CSSProperties> = {
  grid: {
    padding: 10,
    display: "flex",
    justifyContent: "space-evenly",
    alignItems: "center",
  },
  flex: {
    position: "static",
  },
  mobileGrid: {
    padding: 15,
    display: "flex",
    justifyContent: "space-evenly",
    alignItems: "center",
  },
  navContainer: {
    height: "100vh",
    zIndex: 1,
  },
  autocomplete: {
    width: "90%",
    margin: "auto",
  },
}
