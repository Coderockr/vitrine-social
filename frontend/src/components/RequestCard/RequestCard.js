import React from 'react';
import moment from 'moment';
import { Row, Col } from 'antd';
import { Link } from 'react-router-dom';
import ItemIndicator from '../ItemIndicator';
import styles from './styles.module.scss';

class RequestCard extends React.PureComponent {
  render() {
    const { request, isOrganization } = this.props;

    if (!request) {
      return (<div />);
    }

    let dateText = `Criado em: ${moment(request.createdAt).format('LLL')}`;
    if (request.updatedAt) {
      dateText = `Atualizado em: ${moment(request.updatedAt).format('LLL')}`;
    }

    return (
      <Row>
        <Col>
          <div className={styles.requestCard}>
            <ItemIndicator request={request} />
            <div className={styles.organizationContent}>
              <h2>{request.title}</h2>
              <Link to={`/organization/${request.organization.id}`}>{request.organization.name}</Link>
              <p>{dateText}</p>
            </div>
            <div className={styles.interestedContent}>
              <button
                className={styles.detailsButton}
                onClick={() => this.props.onClick()}
              >
                {isOrganization ? 'EDITAR' : 'MAIS DETALHES'}
              </button>
            </div>
          </div>
        </Col>
      </Row>
    );
  }
}

export default RequestCard;
