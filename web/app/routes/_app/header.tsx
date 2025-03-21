import { IconSearch, IconGrillFork } from "@tabler/icons-react"
import { Grid, Button, Autocomplete, Text, Flex, Burger } from "@mantine/core"
import { Categories } from "../../components/categoriesList/categories"
import { useMediaQuery, useDisclosure } from "@mantine/hooks"
import { useEffect, useState } from "react"
import { Navigation } from "../../components/navigation/navigation"
import { Link } from "@remix-run/react"

export function Header() {
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

  const isMobile = useMediaQuery("(max-width: 1199px)")

  return isMobile ? (
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
      />
      <Categories />
    </Flex>
  ) : (
    <div style={{ position: "static" }}>
      <Grid style={styles.grid} justify="space-around">
        <Grid.Col
          style={[styles.grid, { justifyContent: "center" }]}
          span={{ base: 12, md: 6, lg: 2 }}
        >
          <IconGrillFork style={{ paddingRight: 5 }} />
          <Text>FORKD</Text>
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 2 }}>
          <Button variant="transparent" color="gray">
            My Recipes
          </Button>
          <Button variant="transparent" color="gray">
            Add Recipe
          </Button>
        </Grid.Col>
        <Grid.Col
          style={[styles.grid, { justifyContent: "space-evenly" }]}
          span={{ base: 12, md: 6, lg: 6 }}
        >
          <Autocomplete
            style={{ width: "90%" }}
            placeholder="Search for a recipe"
            data={[
              "Vegan Mac and Cheese",
              "Chickpea Salad",
              "Peanut Butter Cookies",
              "Avocado Toast",
            ]}
            leftSection={<IconSearch size={16} stroke={1.5} />}
          />
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 2 }}>
          <Button
            component={Link}
            to="/auth/login"
            variant="transparent"
            color="gray"
          >
            Login
          </Button>
          <Button
            component={Link}
            to="/auth/logout"
            variant="transparent"
            color="gray"
          >
            Logout
          </Button>

          <Button component={Link} to="/auth/signup" color="gray">
            Create Account
          </Button>
        </Grid.Col>
      </Grid>
      <Categories />
    </div>
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
