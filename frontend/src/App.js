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

import Layout from './containers/Layout';

import './App.css';
import reducers from './reducers';

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
        <Route exact path="/" component={Layout} />
      </div>
    </Provider>
  </Router>
);

export default App;
