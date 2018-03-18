import React from 'react';
import { storiesOf } from '@kadira/storybook';

import Dialog from './';

storiesOf('Dialog', module)
  .add('Default View', () => (
    <Dialog active>
      <h1>teste</h1>
    </Dialog>
  ));
