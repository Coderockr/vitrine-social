import React from 'react';
import ReactDOM from 'react-dom';
import MetaTags from 'react-meta-tags';
import App from './App';
import registerServiceWorker from './registerServiceWorker';

const app = (
  <App>
    <MetaTags>
      <meta property="og:locale" content="pt_BR" />
      <meta property="og:url" content="https://www.vitrinesocial.org" />
      <meta property="og:image" content="http://static01.nyt.com/images/2015/02/19/arts/international/19iht-btnumbers19A/19iht-btnumbers19A-facebookJumbo-v2.jpg" />
    </MetaTags>
  </App>
);

ReactDOM.render(app, document.getElementById('root'));
registerServiceWorker();
