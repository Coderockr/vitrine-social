import React from 'react';
import { Row, Col } from 'antd';
import moment from 'moment';
import RequestCard from '../../components/RequestCard';
import './style.css';

const request = {
  organization: {
    name: 'Lar Abdon Batista',
    link: 'http://coderockr.com/',
  },
  category: 'voluntarios',
  data: moment().subtract(2, 'days'),
  item: '10 voluntarios para ler para criancinhas felizes',
  description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
};

const Requests = () => (
  <div className="requests">
    <Row>
      <Col span={20} offset={2}>
        <h2 className="containerTitle">
          <span>SOLICITAÇÕES RECENTES</span>
        </h2>
      </Col>
    </Row>
    <Row>
      <Col lg={{ span: 14, offset: 5 }} md={{ span: 20, offset: 2 }} sm={{ span: 20, offset: 2 }} xs={{ span: 22, offset: 1 }} className="row">
        <RequestCard request={request} />
      </Col>
      <Col lg={{ span: 14, offset: 5 }} md={{ span: 20, offset: 2 }} sm={{ span: 20, offset: 2 }} xs={{ span: 22, offset: 1 }} className="row">
        <RequestCard request={request} />
      </Col>
      <Col lg={{ span: 14, offset: 5 }} md={{ span: 20, offset: 2 }} sm={{ span: 20, offset: 2 }} xs={{ span: 22, offset: 1 }} className="row">
        <RequestCard request={request} />
      </Col>
      <Col lg={{ span: 14, offset: 5 }} md={{ span: 20, offset: 2 }} sm={{ span: 20, offset: 2 }} xs={{ span: 22, offset: 1 }} className="row">
        <RequestCard request={request} />
      </Col>
      <Col lg={{ span: 14, offset: 5 }} md={{ span: 20, offset: 2 }} sm={{ span: 20, offset: 2 }} xs={{ span: 22, offset: 1 }} className="row">
        <RequestCard request={request} />
      </Col>
    </Row>
  </div>
);

export default Requests;
