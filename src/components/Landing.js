import React, { useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Layout from "./Layout";
import Container from "@material-ui/core/Container";
import { Button, TextField } from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    height: "100vh",
  },
}));

function Landing() {
  const classes = useStyles();
  const [message, setMessage] = useState("");
  const [chats, setChats] = useState([]);
  const [connection, setConnection] = useState({});

  const handleSubmit = (e) => {
    e.preventDefault();
    connection.send(message);
    setMessage("");
  };

  const handleMessageChange = (e) => {
    setMessage(e.target.value);
  };

  useEffect(() => {
    let ws = new WebSocket(`ws://${window.location.host}/ws`);
    // let ws = new WebSocket("wss://echo.websocket.org");
    ws.onopen = (e) => {
      setConnection(ws);
    };
  }, []);

  connection.onmessage = (e) => {
    const chatLog = [e.data, ...chats];
    setChats(chatLog);
  };

  return (
    <Layout>
      <main className={classes.root}>
        <Container maxWidth="xs">
          <form noValidate onSubmit={handleSubmit}>
            <TextField
              variant="outlined"
              margin="normal"
              fullWidth
              id="message"
              label="Message"
              name="message"
              autoFocus
              value={message}
              onChange={handleMessageChange}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              className={classes.submit}
            >
              Sign In
            </Button>
            <strong>{connection.readyState && "Connected successfully"}</strong>
            <div>
              {chats.map((chat) => (
                <p>{chat}</p>
              ))}
            </div>
          </form>
        </Container>
      </main>
    </Layout>
  );
}

export default Landing;
