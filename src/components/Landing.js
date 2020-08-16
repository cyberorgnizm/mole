import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Layout from "./Layout";
import Container from "@material-ui/core/Container";
import { Typography, Button, Grid, Box, TextField } from "@material-ui/core";

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
          <Container maxWidth="sm">
            <Typography variant="h4">
              A real-time video streaming service
            </Typography>
            <Grid justify="center" container className={classes.action}>
              <Button color="primary" variant="contained">
                Get started
              </Button>
            </Grid>
            <Grid container>
              <Grid item sm={12}>
                <form>
                  <TextField
                    variant="outlined"
                    margin="normal"
                    fullWidth
                    id="stream_id"
                    label="Join"
                    placeholder="Copy and paste stream ID here to connect"
                    name="stream_id"
                    autoFocus
                    disabled
                  />
                </form>
              </Grid>
            </Grid>
          </Container>
        </Box>
      </main>
    </Layout>
  );
}

export default Landing;
