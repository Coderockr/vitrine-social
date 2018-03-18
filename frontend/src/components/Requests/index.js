import React from 'react';
import {
  Container,
  Column,
  Columns,
  Section,
} from 're-bulma';
import moment from 'moment';
import RequestCard from '../../components/RequestCard';

import './style.css';

const organization = {
  name: 'Lar Abdon batista',
  link: 'http://coderockr.com/',
  category: 'voluntarios',
  data: moment().subtract(2, 'days'),
  item: '10 voluntarios para ler para criancinhas felizes',
  description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
};

const Requests = () => (
  <Section className="requests">
    <Container>
      <Columns>
        <Column>
          <h2 className="containerTitle">
            <span>SOLICITAÇÕES RECENTES</span>
          </h2>
        </Column>
      </Columns>

      <Columns className="row">
        <RequestCard organization={organization} />
      </Columns>
      <Columns className="row">
        <RequestCard organization={organization} />
      </Columns>
      <Columns className="row">
        <RequestCard organization={organization} />
      </Columns>
      <Columns className="row">
        <RequestCard organization={organization} />
      </Columns>
      <Columns className="row">
        <RequestCard organization={organization} />
      </Columns>
      <Columns className="row">
        <RequestCard organization={organization} />
      </Columns>

    </Container>
  </Section>
);

export default Requests;
