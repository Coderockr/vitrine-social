import React from 'react';
import { updateLocale } from 'moment';
import ptBr from 'moment/locale/pt-br';
import {
  Router,
  Route,
} from 'react-router-dom';
import ReactGA from 'react-ga';
import createHistory from 'history/createBrowserHistory';

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

const titleObject = {
  '/': 'Home',
  '/sobre': 'Sobre',
  '/esqueci-senha': 'Esqueci a Senha',
  '/contato': 'Contato',
  '/login': 'Login',
};

const getTitle = (pathname) => {
  const title = titleObject[pathname];
  if (title) {
    return title;
  } if (pathname.search('/complete-registration/') !== -1) {
    return 'Completar Cadastro';
  } if (pathname.search('/recover-password/') !== -1) {
    return 'Recuperar Senha';
  } if (pathname.search('/busca/') !== -1) {
    return 'Busca';
  } if (pathname.search('/entidade/') !== -1) {
    return 'Perfil da Organização';
  }
  return null;
};

const trackPageView = (location) => {
  ReactGA.set({ page: location.pathname });
  ReactGA.pageview(location.pathname, null, getTitle(location.pathname));
};

const initGA = (hist) => {
  ReactGA.initialize('UA-122417824-1');
  trackPageView(hist.location);
  hist.listen(trackPageView);
};

const history = createHistory();
initGA(history);

const App = () => (
  <Router history={history}>
    <div className="App">
      <Route exact path="/" component={Home} />
      <Route exact path="/sobre" component={About} />
      <Route exact path="/contato" component={Contact} />
      <Route exact path="/busca/:searchParams" component={Results} />
      <Route exact path="/entidade/:organizationId" component={OrganizationProfile} />
      <Route exact path="/detalhes/:requestId" component={Home} />
      <Route exact path="/login" component={Login} />
      <Route exact path="/esqueci-senha" component={ForgotPassword} />
      <Route exact path="/complete-registration/:token" component={ResetPassword} />
      <Route exact path="/recover-password/:token" component={ResetPassword} />
    </div>
  </Router>
);

export default App;
