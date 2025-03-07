import { Grid } from "@mantine/core"

export function Footer() {
  return (
    <Grid style={styles.foot}>
      <Grid.Col span={{ base: 12, md: 12, lg: 6 }}>1</Grid.Col>
      <Grid.Col span={{ base: 12, md: 12, lg: 6 }}>2</Grid.Col>
    </Grid>
  )
}

const styles = {
  foot: {
    height: 80,
    background: "black",
    position: "fixed",
    bottom: 0,
    width: "100%",
    color: "white",
  },
}
