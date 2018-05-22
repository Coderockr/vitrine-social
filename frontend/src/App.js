import React from 'react';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import { updateLocale } from 'moment';
import ptBr from 'moment/locale/pt-br';
import thunk from 'redux-thunk';
import {
  BrowserRouter as Router,
  Route,
} from 'react-router-dom';
import reducers from './reducers';

import Home from './containers/Home';
import OrganizationProfile from './containers/OrganizationProfile';

import './utils/styles/global.module.scss';

updateLocale('pt-br', ptBr);

const store = createStore(
  reducers,
  {},
  applyMiddleware(thunk),
);

const App = () => (
  <Router>
    <Provider store={store}>
      <div className="App">
        <Route exact path="/" component={Home} />
        <Route exact path="/organization" component={OrganizationProfile} />
      </div>
    </Provider>
  </Router>
);

export default App;
