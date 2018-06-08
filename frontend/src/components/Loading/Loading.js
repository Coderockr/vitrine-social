import React from 'react';
import { Row, Col, Icon } from 'antd';
import styles from './styles.module.scss';
import colors from '../../utils/styles/colors';

const Loading = () => (
  <Row type="flex" align="center" justify="center">
    <Col>
      <div className={styles.contentWrapper}>
        <Icon type="loading" style={{ fontSize: 110, color: colors.purple_400 }} />
      </div>
    </Col>
  </Row>
);


export default Loading;
