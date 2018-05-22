import React from 'react';
import { Row, Col, Radio } from 'antd';
import moment from 'moment';
import RequestCard from '../../components/RequestCard';
import styles from './styles.module.scss';

const allRequests = [
  {
    organization: {
      name: 'Lar Abdon Batista',
      link: 'http://coderockr.com/',
    },
    category: 'voluntarios',
    date: moment().subtract(2, 'days'),
    item: '10 voluntarios para ler para criancinhas felizes',
    description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
  },
  {
    organization: {
      name: 'Lar Abdon Batista',
      link: 'http://coderockr.com/',
    },
    category: 'voluntarios',
    date: moment().subtract(2, 'days'),
    item: '10 voluntarios para ler para criancinhas felizes',
    description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
  },
  {
    organization: {
      name: 'Lar Abdon Batista',
      link: 'http://coderockr.com/',
    },
    category: 'voluntarios',
    date: moment().subtract(2, 'days'),
    item: '10 voluntarios para ler para criancinhas felizes',
    description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
  },
  {
    organization: {
      name: 'Lar Abdon Batista',
      link: 'http://coderockr.com/',
    },
    category: 'voluntarios',
    date: moment().subtract(2, 'days'),
    item: '10 voluntarios para ler para criancinhas felizes',
    description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
  },
  {
    organization: {
      name: 'Lar Abdon Batista',
      link: 'http://coderockr.com/',
    },
    category: 'voluntarios',
    date: moment().subtract(2, 'days'),
    item: '10 voluntarios para ler para criancinhas felizes',
    description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
  },
];

const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;

class Requests extends React.Component {
  renderRequests(requests) {
    return (
      requests.map(request => (
        <div className={styles.requestWrapper}>
          <RequestCard request={request} isOrganization={this.props.isOrganization} />
        </div>
      ))
    );
  }

  render() {
    return (
      <div className={styles.requests} >
        <Row>
          <Col span={20} offset={2}>
            <h2 className={styles.containerTitle}>
              <span>SOLICITAÇÕES RECENTES</span>
            </h2>
          </Col>
        </Row>
        <Row>
          <Col
            lg={{ span: 14, offset: 5 }}
            md={{ span: 20, offset: 2 }}
            sm={{ span: 20, offset: 2 }}
            xs={{ span: 22, offset: 1 }}
            className={styles.row}
          >
            <div className={styles.actionWrapper} hidden={!this.props.isOrganization}>
              <RadioGroup defaultValue="Ativas">
                <RadioButton value="Ativas">ATIVAS</RadioButton>
                <RadioButton value="Inativas">INATIVAS</RadioButton>
              </RadioGroup>
              <button
                className={styles.button}
                onClick={this.props.onClick}
              >
                NOVA SOLICITAÇÃO
              </button>
            </div>
            {this.renderRequests(allRequests)}
          </Col>
        </Row>
      </div>
    );
  }
}

export default Requests;
