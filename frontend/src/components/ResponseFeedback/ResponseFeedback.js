import React from 'react';
import { Row, Col, Icon } from 'antd';
import cx from 'classnames';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

class ResponseFeedback extends React.Component {
  renderFeedback() {
    if (this.props.type === 'loading') {
      return (
        <div className={styles.contentWrapper}>
          <Icon type="loading" style={{ fontSize: 110, color: colors.purple_400 }} />
        </div>
      );
    }

    let feedback = {
      icon: { type: 'close-circle-o', color: colors.red_400 },
      title: 'Erro!',
      buttonClassName: styles.errorButton,
      buttonTitle: 'VOLTAR',
    };

    if (this.props.type === 'success') {
      feedback = {
        icon: { type: 'check-circle-o', color: colors.green_300 },
        title: 'Sucesso!',
        buttonClassName: styles.successButton,
        buttonTitle: 'FECHAR',
      };
    }

    return (
      <div className={styles.contentWrapper}>
        <Icon type={feedback.icon.type} style={{ fontSize: 150, color: feedback.icon.color }} />
        <p className={styles.title}>{feedback.title}</p>
        <p className={styles.message}>{this.props.message}</p>
        <button
          className={cx(styles.button, feedback.buttonClassName)}
          onClick={this.props.onClick}
        >
          {feedback.buttonTitle}
        </button>
      </div>
    );
  }

  render() {
    return (
      <div className={styles.wrapper} hidden={!this.props.type}>
        <Row type="flex" align="bottom" justify="center" className={styles.row}>
          <Col>
            {this.renderFeedback()}
          </Col>
        </Row>
      </div>
    );
  }
}

export default ResponseFeedback;
