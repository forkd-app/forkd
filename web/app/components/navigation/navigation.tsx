import { Grid, NavLink } from "@mantine/core"

export function Navigation() {
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
