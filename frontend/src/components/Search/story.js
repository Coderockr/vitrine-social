import React from 'react';
import { storiesOf } from '@kadira/storybook';

import Search from './';

storiesOf('Search', module)
  .add('Default View', () => (
    <Search />
  ));
