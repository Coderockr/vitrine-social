import React from 'react';
import { storiesOf } from '@kadira/storybook';
import moment from 'moment';

import RequestCard from './';

const request = {
  organization: {
    name: 'Lar Abdon batista',
    link: 'http://coderockr.com/',
  },
  category: 'voluntarios',
  data: moment().subtract(2, 'days'),
  item: '10 voluntarios para ler para criancinhas felizes',
  description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
};

storiesOf('RequestCard', module)
  .add('Default View', () => (
    <RequestCard request={request} />
  ));
