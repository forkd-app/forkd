import { IconSearch, IconGrillFork } from "@tabler/icons-react"
// import { useDisclosure } from "@mantine/hooks"
import { Grid, Button, Autocomplete, Text } from "@mantine/core"
import { Categories } from "../categoriesList/categories"

// const links = [
//   { link: '/about', label: 'Features' },
//   { link: '/pricing', label: 'Pricing' },
//   { link: '/learn', label: 'Learn' },
//   { link: '/community', label: 'Community' },
// ];

export function Header() {
  // const [opened, { toggle }] = useDisclosure(false);

  // const items = links.map((link) => (
  //   <a
  //     key={link.label}
  //     href={link.link}
  //     className={classes.link}
  //     onClick={(event) => event.preventDefault()}
  //   >
  //     {link.label}
  //   </a>
  // ));

  return (
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
            {" "}
            My Recipes{" "}
          </Button>
          <Button variant="transparent" color="gray">
            {" "}
            Add Recipe{" "}
          </Button>
        </Grid.Col>
        <Grid.Col
          style={[styles.grid, { justifyContent: "space-evenly" }]}
          span={{ base: 12, md: 6, lg: 6 }}
        >
          <Autocomplete
            style={{ width: "90%" }}
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
        </Grid.Col>
        <Grid.Col style={styles.grid} span={{ base: 12, md: 6, lg: 2 }}>
          <Button variant="transparent" color="gray">
            {" "}
            Log In{" "}
          </Button>
          <Button variant="" color="gray">
            {" "}
            Create Account{" "}
          </Button>
        </Grid.Col>
      </Grid>
      <Categories />
    </div>
  )
}

const styles = {
  grid: {
    padding: 10,
    display: "flex",
    justifyContent: "space-evenly",
    alignItems: "center",
  },
}
