import React from 'react';
import { updateLocale } from 'moment';
import ptBr from 'moment/locale/pt-br';
import {
  BrowserRouter as Router,
  Route,
} from 'react-router-dom';

import Home from './containers/Home';
import About from './containers/About';
import Contact from './containers/Contact';
import Results from './containers/Results';
import Login from './containers/Login';
import ForgotPassword from './containers/ForgotPassword';
import ResetPassword from './containers/ResetPassword';
import OrganizationProfile from './containers/OrganizationProfile';

import './utils/styles/global.module.scss';

updateLocale('pt-br', ptBr);

const App = () => (
  <Router>
    <div className="App">
      <Route exact path="/" component={Home} />
      <Route exact path="/about" component={About} />
      <Route exact path="/contact" component={Contact} />
      <Route exact path="/search/:searchParams" component={Results} />
      <Route exact path="/organization/:organizationId" component={OrganizationProfile} />
      <Route exact path="/login" component={Login} />
      <Route exact path="/forgot-password" component={ForgotPassword} />
      <Route exact path="/complete-registration/:token" component={ResetPassword} />
      <Route exact path="/recover-password/:token" component={ResetPassword} />
    </div>
  </Router>
);

export default App;
