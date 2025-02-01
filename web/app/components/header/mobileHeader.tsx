import { Grid, Text, Burger, Autocomplete, Flex } from "@mantine/core"
import { IconGrillFork, IconSearch } from "@tabler/icons-react"
import { useDisclosure } from "@mantine/hooks"
import { useState, useEffect } from "react"
import { Navigation } from "../navigation/navigation"
import { Categories } from "../categoriesList/categories"

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
      <Grid style={styles.grid} justify="space-around">
        <Grid.Col
          style={[styles.grid, { justifyContent: "left" }]}
          span={{ base: 6, md: 6, lg: 6 }}
        >
          <IconGrillFork style={{ paddingRight: 5 }} />
          <Text>FORKD</Text>
        </Grid.Col>
        <Grid.Col
          style={[styles.grid, { justifyContent: "right" }]}
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
        label=""
        placeholder="Search for a recipe"
        data={[
          "Vegan Mac and Cheese",
          "Chickpea Salad",
          "Peanut Butter Cookies",
          "Avocado Toast",
        ]}
        leftSection={<IconSearch size={16} stroke={1.5} />}
      />
      <Categories />
    </Flex>
  )
}

const styles: Record<string, React.CSSProperties> = {
  flex: {
    position: "static",
  },
  grid: {
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
