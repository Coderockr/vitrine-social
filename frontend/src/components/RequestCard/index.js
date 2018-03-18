import React from 'react';
import moment from 'moment';
import {
  Container,
  Media,
  MediaLeft,
  MediaRight,
  Title,
} from 're-bulma';

import Icon from '../Icons';
import ProgressCircle from '../ProgressCircle';

import './style.css';

const RequestCard = ({ organization }) => (
  <Container isFullwidth>
    <Media className="requestCard">
      <MediaLeft className="requestIcon">
        <Icon icon={organization.category} size={60} color="#FF974B" />
        <div className="progress-circle-container">
          <ProgressCircle progress={30} />
          <div className="laste-qtd">
            <p>
              Faltam 4
            </p>
          </div>
        </div>
      </MediaLeft>
      <div className="organizationContent">
        <Title size="is5">{organization.item}</Title>
        <a href={organization.link} target="_blank">
          {organization.name}
        </a>
        <p>
          Atualizado em: {
            moment(organization.data).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
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
