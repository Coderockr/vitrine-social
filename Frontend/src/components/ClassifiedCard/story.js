import React from 'react';
import { storiesOf } from '@kadira/storybook';

import ClassifiedCard from './';

const organization = {
  name: 'Sony Dual shock 4 GamePad',
  link: '',
  image: 'http://placehold.it/128x80',
  categories: ['alimentos', 'brinquedos'],
  description: 'Non-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international'
}

storiesOf('ClassifiedCard', module)
  .add('Default View', () => (
    <ClassifiedCard organization={organization} />
  ));
