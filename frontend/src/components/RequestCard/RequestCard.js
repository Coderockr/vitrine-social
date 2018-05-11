import React from 'react';
import moment from 'moment';
import { Row, Col } from 'antd';
import ItemIndicator from '../ItemIndicator';

import './style.css';

const RequestCard = ({ request }) => (
  <Row>
    <Col>
      <div className="requestCard">
        <ItemIndicator request={request} />
        <div className="organizationContent">
          <h2>{request.item}</h2>
          <a href={request.organization.link} target="_blank">
            {request.organization.name}
          </a>
          <p>
            Atualizado em: {
              moment(request.data).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
            }
          </p>
        </div>
        <div className="interestedContent">
          <button className="button">
            MAIS DETALHES
          </button>
        </div>
      </div>
    </Col>
  </Row>
);

export default RequestCard;
