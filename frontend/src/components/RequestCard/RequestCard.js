import React from 'react';
import moment from 'moment';
import {
  Container,
  Media,
  MediaRight,
  Title,
} from 're-bulma';
import ItemIndicator from '../ItemIndicator';

import './style.css';

const RequestCard = ({ request }) => (
  <Container isFullwidth>
    <Media className="requestCard">
      <ItemIndicator request={request} />
      <div className="organizationContent">
        <Title size="is5">{request.item}</Title>
        <a href={request.organization.link} target="_blank">
          {request.organization.name}
        </a>
        <p>
          Atualizado em: {
            moment(request.data).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
          }
        </p>
      </div>
      <MediaRight className="interestedContent">
        <button className="button">
          MAIS DETALHES
        </button>
      </MediaRight>
    </Media>
  </Container>
);

export default RequestCard;
