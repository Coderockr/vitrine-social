import React from 'react';
import { updateLocale } from 'moment';
import ptBr from 'moment/locale/pt-br';
import {
  BrowserRouter as Router,
  Route,
} from 'react-router-dom';

import Home from './containers/Home';
import About from './containers/About';
import Results from './containers/Results';
import Login from './containers/Login';
import OrganizationProfile from './containers/OrganizationProfile';

import './utils/styles/global.module.scss';

updateLocale('pt-br', ptBr);

const App = () => (
  <Router>
    <div className="App">
      <Route exact path="/" component={Home} />
      <Route exact path="/about" component={About} />
      <Route exact path="/search/:searchParams" component={Results} />
      <Route exact path="/organization/:organizationId" component={OrganizationProfile} />
      <Route exact path="/login" component={Login} />
    </div>
  </Router>
);

export default App;
