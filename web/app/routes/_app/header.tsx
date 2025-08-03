import { IconSearch, IconGrillFork, IconFilter } from "@tabler/icons-react"
import {
  Grid,
  Button,
  Autocomplete,
  Text,
  Flex,
  Popover,
  ActionIcon,
} from "@mantine/core"
import { useMediaQuery } from "@mantine/hooks"
import { Link } from "@remix-run/react"
import { MobileHeader } from "./mobileHeader"
import { useSelector } from "react-redux"
import type { RootState } from "~/stores/global"

export function Header() {
  const user = useSelector((state: RootState) => state.user.value)
  const isMobile = useMediaQuery("(max-width: 1199px)")

  return isMobile ? (
    <MobileHeader />
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
            rightSection={
              <Popover width={200} position="bottom" withArrow shadow="md">
                <Popover.Target>
                  <ActionIcon variant="transparent">
                    <Text>
                      <IconFilter color="gray" size="18" />
                    </Text>
                  </ActionIcon>
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
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 2 }}>
          {user ? (
            <>
              <Button
                component={Link}
                to="/auth/logout"
                variant="transparent"
                color="gray"
              >
                Logout
              </Button>
              <Button component={Link} to="" color="gray">
                Manage Account
              </Button>
            </>
          ) : (
            <>
              <Button
                component={Link}
                to="/auth/login"
                variant="transparent"
                color="gray"
              >
                Login
              </Button>
              <Button component={Link} to="/auth/signup" color="gray">
                Create Account
              </Button>
            </>
          )}
        </Grid.Col>
      </Grid>
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
