import React from 'react';
import { Row, Col, Radio } from 'antd';
import cx from 'classnames';
import ReactGA from 'react-ga';
import RequestCard from '../../components/RequestCard';
import RequestForm from '../../components/RequestForm';
import RequestFormCreate from '../../components/RequestFormCreate';
import RequestDetails from '../../components/RequestDetails';
import styles from './styles.module.scss';
import Loading from '../Loading/Loading';
import ErrorCard from '../../components/ErrorCard';

const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;

class Requests extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      showModal: false,
      request: null,
      status: 'ACTIVE',
      saveEnabled: false,
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

    if (modal === 'editForm') {
      const action = request ? 'Editar' : 'Nova';
      ReactGA.event({
        category: 'Entidade',
        action: `${action} Solicitação`,
      });
    }
  }

  statusChanged() {
    let status = 'ACTIVE';
    if (this.state.status === 'ACTIVE') {
      status = 'INACTIVE';
    }
    this.setState({ status });
    ReactGA.event({
      category: 'Entidade',
      action: `Switch para Status ${status}`,
    });
  }

  renderRequests() {
    if (this.props.loading) {
      return <Loading />;
    }

    if (this.props.error) {
      return <ErrorCard text="Não foi possível listar as solicitações!" />;
    }

    const requests = (this.state.status === 'INACTIVE' ? this.props.inactiveRequests : this.props.activeRequests);
    if (requests.length === 0) {
      const status = this.state.status === 'ACTIVE' ? 'ativa' : 'inativa';
      return (
        <div className={styles.emptyWrapper}>
          <p className={styles.emptyText}>Não há nenhuma solicitação {status}!</p>
        </div>
      );
    }

    return (
      requests.map((request) => {
        if (request.status && request.status !== this.state.status) {
          return null;
        }
        return (
          <div className={styles.requestWrapper} key={request.id}>
            <RequestCard
              request={request}
              isOrganization={this.props.isOrganization}
              onClick={() => this.showModal(request, this.props.isOrganization ? 'editForm' : 'details')}
            />
          </div>
        );
      }));
  }

  render() {
    return (
      <div className={cx(styles.requests, { [styles.requestsSearch]: this.props.search })} >
        <Row>
          <Col span={20} offset={2}>
            <h2 className={styles.containerTitle}>
              <span>{this.props.search ? 'RESULTADO DA BUSCA' : 'SOLICITAÇÕES RECENTES'}</span>
            </h2>
          </Col>
        </Row>
        <Row>
          <Col
            lg={{ span: 18, offset: 3 }}
            md={{ span: 20, offset: 2 }}
            sm={{ span: 22, offset: 1 }}
            xs={{ span: 24, offset: 0 }}
            className={styles.row}
          >
            {this.props.isOrganization &&
              <div className={styles.organizationActionWrapper}>
                <RadioGroup defaultValue="ACTIVE" onChange={() => this.statusChanged()}>
                  <RadioButton value="ACTIVE">ATIVAS</RadioButton>
                  <RadioButton value="INACTIVE">INATIVAS</RadioButton>
                </RadioGroup>
                <button
                  className={styles.newButton}
                  onClick={() => this.showModal(null, 'newForm')}
                >
                  NOVA SOLICITAÇÃO
                </button>
              </div>
            }
            {!this.props.isOrganization &&
              <div className={styles.actionWrapper}>
                <RadioGroup className="purpleRadio" defaultValue="NEWEST" onChange={this.props.orderChanged}>
                  <RadioButton value="NEWEST">MAIS RECENTES</RadioButton>
                  <RadioButton value="OLDEST">MAIS ANTIGAS</RadioButton>
                </RadioGroup>
              </div>
            }
            {this.renderRequests()}
            {this.props.isOrganization && this.state.showModal === 'editForm' &&
              <RequestForm
                visible
                onCancel={() => this.onCancel()}
                onSave={this.props.onSave}
                saveEnabled={this.state.saveEnabled}
                enableSave={enable => this.setState({ saveEnabled: enable })}
                request={this.state.request}
              />
            }
            {this.props.isOrganization && this.state.showModal === 'newForm' &&
              <RequestFormCreate
                visible
                onCancel={() => this.onCancel()}
                onSave={this.props.onSave}
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
