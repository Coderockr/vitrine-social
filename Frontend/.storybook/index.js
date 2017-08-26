import React from 'react';
import { storiesOf } from '@kadira/storybook';

import Button from './Button';

storiesOf('Button', module)
  .add('Default View', () => (
      <Button onClick={console.log} >
        teste
      </Button>
  ));
