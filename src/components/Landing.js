import React from "react";
import { Link as RouteLink } from "react-router-dom";
import { makeStyles } from "@material-ui/core/styles";
import Layout from "./Layout";
import Container from "@material-ui/core/Container";
import {
  Typography,
  Button,
  Grid,
  Box,
  TextField,
  Link,
} from "@material-ui/core";
import { GitHub } from "@material-ui/icons";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    height: "100vh",
  },
  action: {
    marginTop: "20px",
  },
}));

function Landing() {
  const classes = useStyles();

  return (
    <Layout>
      <main className={classes.root}>
        <Box
          minHeight="100vh"
          display="flex"
          justifyContent="center"
          alignItems="center"
        >
          <Container maxWidth="md">
            <Box my={5}>
              <Grid
                container
                direction="column"
                justify="center"
                alignItems="center"
              >
                <Grid item xs={12}>
                  <GitHub fontSize="large" />
                </Grid>
                <Grid item xs={12}>
                  <Typography align="center" variant="body1">
                    <Link href="https://github.com/cyberorgnizm/mole">
                      View Source
                    </Link>
                  </Typography>
                </Grid>
              </Grid>
            </Box>
            <Box mb={4}>
              <Typography variant="h5">
                A real-time video streaming service that lets you stay in touch
                with your loved ones
              </Typography>
            </Box>
            <Grid container justify="center" alignItems="center" spacing={3}>
              <Grid item sm={6}>
                <form>
                  <TextField
                    variant="outlined"
                    margin="normal"
                    fullWidth
                    id="stream_id"
                    label="Create or Join"
                    placeholder="Enter stream ID here to connect"
                    name="stream_id"
                    autoFocus
                    size="medium"
                  />
                </form>
              </Grid>
              <Grid item sm={3}>
                <Box mt={1}>
                  <RouteLink to="/test">
                    <Button color="primary" variant="contained" size="large">
                      Get started
                    </Button>
                  </RouteLink>
                </Box>
              </Grid>
            </Grid>
          </Container>
        </Box>
      </main>
    </Layout>
  );
}

export default Landing;
