import React from 'react';
import moment from 'moment';
import { Row, Col } from 'antd';
import ItemIndicator from '../ItemIndicator';
import styles from './styles.module.scss';

class RequestCard extends React.PureComponent {
  render() {
    const { request, isOrganization } = this.props;

    if (!request) {
      return (<div />);
    }

    return (
      <Row>
        <Col>
          <div className={styles.requestCard}>
            <ItemIndicator request={request} />
            <div className={styles.organizationContent}>
              <h2>{request.title}</h2>
              <a target="_blank">
                {request.organization.name}
              </a>
              <p>
                Atualizado em: {
                  moment(request.updatedAt)
                    .format('DD / MMMM / YYYY')
                    .replace(/(\/)/g, 'de')
                }
              </p>
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
