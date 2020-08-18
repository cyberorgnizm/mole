import React from "react";
import { makeStyles } from "@material-ui/core/styles";

import { Box } from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    minHeight: "100vh",
  },
  body: {
    minHeight: "100vh",
  },
}));

function Layout({ children }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Box height="100%" width="100%" className={classes.body}>
        {children}
      </Box>
    </div>
  );
}

export default Layout;
