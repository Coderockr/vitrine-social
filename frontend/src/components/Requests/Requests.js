import React from 'react';
import { Row, Col, Radio } from 'antd';
import moment from 'moment';
import RequestCard from '../../components/RequestCard';
import RequestForm from '../../components/RequestForm';
import RequestDetails from '../../components/RequestDetails';
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
  constructor(props) {
    super(props);
    this.state = {
      showModal: false,
      request: null,
    };

    this.showModal = this.showModal.bind(this);
    this.onCancel = this.onCancel.bind(this);
  }

  onCancel() {
    this.setState({
      showModal: false,
      request: null,
    });
  }

  showModal(request, modal) {
    this.setState({
      showModal: modal,
      request,
    });
  }

  renderRequests(requests) {
    return (
      requests.map(request => (
        <div className={styles.requestWrapper}>
          <RequestCard
            request={request}
            isOrganization={this.props.isOrganization}
            onClick={() => this.showModal(request, this.props.isOrganization ? 'editForm' : 'details')}
          />
        </div>
      ))
    );
  }

  render() {
    return (
      <div className={styles.requests}>
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
            {this.props.isOrganization &&
              <div className={styles.actionWrapper}>
                <RadioGroup defaultValue="Ativas">
                  <RadioButton value="Ativas">ATIVAS</RadioButton>
                  <RadioButton value="Inativas">INATIVAS</RadioButton>
                </RadioGroup>
                <button
                  className={styles.button}
                  onClick={() => this.showModal(null, 'editForm')}
                >
                  NOVA SOLICITAÇÃO
                </button>
              </div>
            }
            {this.renderRequests(allRequests)}
            {this.props.isOrganization &&
              <RequestForm
                visible={this.state.showModal === 'editForm'}
                onCancel={() => this.onCancel()}
                request={this.state.request}
              />
            }
            {!this.props.isOrganization &&
              <RequestDetails
                visible={this.state.showModal === 'details'}
                onCancel={() => this.onCancel()}
                request={this.state.request}
              />
            }
          </Col>
        </Row>
      </div>
    );
  }
}

export default Requests;
