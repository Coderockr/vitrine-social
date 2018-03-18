import React from 'react';
import { storiesOf } from '@kadira/storybook';

import Requests from './';

storiesOf('Requests', module)
  .add('Requests', () => (
    <Requests />
  ));
