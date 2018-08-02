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

    const {
      organization,
      requiredQuantity,
      reachedQuantity,
      title,
      unit,
    } = request;
    return (
      <Row>
        <Col>
          <div className={styles.requestCard}>
            <ItemIndicator request={request} />
            <div className={styles.organizationContent}>
              <div>
                <p className={styles.receivedTop}>{`Recebidos: ${reachedQuantity} de ${requiredQuantity} ${unit}`}</p>
                <h2>{`${title} (${requiredQuantity} ${unit})`}</h2>
                <Link to={`/organization/${organization.id}`}>{organization.name}</Link>
              </div>
              <p className={styles.date}>{dateText}</p>
              <p className={styles.receivedBottom}>{`Recebidos: ${reachedQuantity} de ${requiredQuantity} ${unit}`}</p>
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
