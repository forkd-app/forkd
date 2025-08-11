import { Grid, NavLink, Button } from "@mantine/core"
import { useGlobals } from "~/stores/global"
import { Link } from "react-router"

export function Navigation() {
  const { user } = useGlobals()

  return (
    <>
      <Grid style={styles.gridContainer} justify="center">
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
          <NavLink variant="subtle" color="gray" label="Text" />
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
          <NavLink variant="subtle" color="gray" label="Text" />
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
          <NavLink variant="subtle" color="gray" label="Text" />
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
          <NavLink variant="subtle" color="gray" label="Text" />
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
          <NavLink
            variant="subtle"
            href="#required-for-focus"
            label="Text"
            childrenOffset={28}
          >
            <NavLink label="Child Link 1" href="#required-for-focus" />
            <NavLink label="Child Link 2" href="#required-for-focus" />
            <NavLink label="Child Link 3" href="#required-for-focus" />
          </NavLink>
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
          <NavLink variant="subtle" color="gray" label="Text" />
          {user ? (
            <>
              <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
                <Button
                  component={Link}
                  to="/auth/logout"
                  variant="transparent"
                  color="gray"
                >
                  Logout
                </Button>
              </Grid.Col>
              <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
                <Button component={Link} to="" color="gray">
                  Manage Account
                </Button>
              </Grid.Col>
            </>
          ) : (
            <>
              <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
                <Button
                  component={Link}
                  to="/auth/login"
                  variant="transparent"
                  color="gray"
                >
                  Login
                </Button>
              </Grid.Col>
              <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 1 }}>
                <Button component={Link} to="/auth/signup" color="gray">
                  Create Account
                </Button>
              </Grid.Col>
            </>
          )}
        </Grid.Col>
      </Grid>
    </>
  )
}

const styles = {
  gridContainer: {
    padding: 10,
  },
  grid: {
    padding: 0,
    justifyContent: "space-evenly",
  },
}
