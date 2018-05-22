import React from 'react';
import moment from 'moment';
import { Row, Col } from 'antd';
import ItemIndicator from '../ItemIndicator';
import styles from './styles.module.scss';

class RequestCard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <Row>
        <Col>
          <div className={styles.requestCard}>
            <ItemIndicator request={this.props.request} />
            <div className={styles.organizationContent}>
              <h2>{this.props.request.item}</h2>
              <a href={this.props.request.organization.link} target="_blank">
                {this.props.request.organization.name}
              </a>
              <p>
                Atualizado em: {
                  moment(this.props.request.date).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
                }
              </p>
            </div>
            <div className={styles.interestedContent}>
              <button
                className={styles.button}
                onClick={this.props.isOrganization ?
                  () => this.props.onEdit(this.props.request) :
                  () => this.props.onClick(this.props.request)}
              >
                {this.props.isOrganization ? 'EDITAR' : 'MAIS DETALHES'}
              </button>
            </div>
          </div>
        </Col>
      </Row>
    );
  }
}

export default RequestCard;
