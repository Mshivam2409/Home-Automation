import React from 'react';
import SignIn from "pages/SignIn"
import './App.css';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/" component={SignIn} />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
