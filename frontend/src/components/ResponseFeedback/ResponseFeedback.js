import React from 'react';
import { Row, Col, Icon } from 'antd';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

class ResponseFeedback extends React.Component {
  state = {
  }

  renderSuccess() {
    return (
      <div className={styles.contentWrapper}>
        <Icon type="check-circle-o" style={{ fontSize: 150, color: colors.green_300 }} />
        <p className={styles.title}>Sucesso!</p>
        <p className={styles.message}>Mensagem de Sucesso</p>
      </div>
    );
  }

  renderError() {
    return (
      <div className={styles.contentWrapper}>
        <Icon type="close-circle-o" style={{ fontSize: 150, color: colors.red_400 }} />
        <p className={styles.title}>Erro!</p>
        <p className={styles.message}>Mensagem de Erro</p>
      </div>
    );
  }

  renderLoading() {
    return (
      <div className={styles.contentWrapper}>
        <Icon type="loading" style={{ fontSize: 110, color: colors.purple_400 }} />
      </div>
    );
  }

  renderContent() {
    if (this.props.type === 'error') {
      return this.renderError();
    } if (this.props.type === 'success') {
      return this.renderSuccess();
    } if (this.props.type === 'loading') {
      return this.renderLoading();
    }
    return null;
  }

  render() {
    return (
      <div className={styles.wrapper} hidden={!this.props.type}>
        <Row type="flex" align="bottom" justify="center" className={styles.row}>
          <Col>
            {this.renderContent()}
          </Col>
        </Row>
      </div>
    );
  }
}

export default ResponseFeedback;
