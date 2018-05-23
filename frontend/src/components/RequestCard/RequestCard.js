import React from 'react';
import moment from 'moment';
import { Row, Col } from 'antd';
import RequestDetails from '../RequestDetails';
import RequestDetailsEdit from '../RequestDetailsEdit';
import ItemIndicator from '../ItemIndicator';
import styles from './styles.module.scss';

class RequestCard extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      showDetails: false,
      editRequest: false,
    };

    this.openModal = this.openModal.bind(this);
    this.onCancel = this.onCancel.bind(this);
  }

  onCancel() {
    this.setState({
      showDetails: false,
      editRequest: false,
    });
  }

  openModal(modal) {
    if (modal === 'edit') {
      return this.setState({ editRequest: true });
    }

    return this.setState({ showDetails: true });
  }

  renderButton(isOrganization) {
    return (
      <button
        className={styles.button}
        onClick={() => this.openModal(isOrganization ? 'edit' : 'details')}
      >
        {isOrganization ? 'EDITAR' : 'MAIS DETALHES'}
      </button>
    );
  }

  render() {
    const { request, isOrganization } = this.props;

    if (!request) {
      return (<div />);
    }

    return (
      <Row>
        <Col>
          {!isOrganization &&
            <RequestDetails
              visible={this.state.showDetails}
              onCancel={this.onCancel}
            />
          }
          {isOrganization &&
            <RequestDetailsEdit
              visible={this.state.editRequest}
              onCancel={this.onCancel}
              request={request}
            />
          }
          <div className={styles.requestCard}>
            <ItemIndicator request={request} />
            <div className={styles.organizationContent}>
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
            <div className={styles.interestedContent}>
              {this.renderButton(isOrganization)}
            </div>
          </div>
        </Col>
      </Row>
    );
  }
}

export default RequestCard;
