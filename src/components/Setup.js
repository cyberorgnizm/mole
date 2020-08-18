import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Layout from "./Layout";
import Container from "@material-ui/core/Container";
import { Button, TextField, Grid, Box } from "@material-ui/core";
import { useWebRTC } from "../hooks";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    height: "100vh",
  },
  video: {
    width: "100%",
    height: "600px",
  },
}));

function Setup() {
  const classes = useStyles();

  const [name, setName] = useState("");
  const streamRef = useWebRTC();

  const handleNameChange = (e) => {
    setName(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
  };

  return (
    <Layout>
      <main className={classes.root}>
        <Grid container spacing={2} justify="center" alignItems="center">
          <Grid item xs={12}>
            <Box mt={2}>
              <video
                autoPlay
                playsInline
                width="100%"
                className={classes.video}
                ref={streamRef}
              />
            </Box>
          </Grid>
        </Grid>
        <Box mt={4}>
          <Container maxWidth="xs">
            <form noValidate onSubmit={handleSubmit}>
              <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                id="name"
                label="Set your display name"
                name="message"
                autoFocus
                value={name}
                onChange={handleNameChange}
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
              >
                Join Stream
              </Button>
            </form>
          </Container>
        </Box>
      </main>
    </Layout>
  );
}

export default Setup;
