import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Landing from "./components/Landing";
import Setup from "./components/Setup";

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/test" component={Setup} />
        <Route path="/" component={Landing} />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
