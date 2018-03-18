import React from 'react';
import { storiesOf } from '@kadira/storybook';

import Header from './';

storiesOf('Header', module)
  .add('Default View', () => (
    <Header />
  ))
  .add('Header Active', () => (
    <Header />
  ));
