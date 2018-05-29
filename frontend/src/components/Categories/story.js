import React from 'react';
import { storiesOf } from '@kadira/storybook';

import Categories from './';

storiesOf('Categories', module)
  .add('Categories', () => (
    <Categories />
  ));
