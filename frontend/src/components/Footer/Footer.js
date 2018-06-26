import React from 'react';
import { Row, Col, Layout, Icon } from 'antd';
import cx from 'classnames';
import styles from './styles.module.scss';

const iconStyle = {
  fontSize: 30,
  margin: 10,
  color: 'white',
};

const Footer = ({ className }) => (
  <Layout.Footer className={cx(styles.appFooter, className)}>
    <Row>
      <Col
        xxl={{ span: 14, offset: 5 }}
        xl={{ span: 20, offset: 2 }}
        lg={{ span: 22, offset: 1 }}
        md={{ span: 24, offset: 0 }}
      >
        <div className={styles.footerWrapper}>
          <p className={styles.copyright}>Vitrine Social ©2018 Created by  <a target="_blank" rel="me" href="http://www.coderockr.com">Coderockr</a></p>
          <div>
            <a target="_blank" rel="me" href="https://www.facebook.com/avitrinesocial/"><Icon type="facebook" style={iconStyle} /></a>
            <a target="_blank" rel="me" href="https://www.instagram.com/avitrine.social/"><Icon type="instagram" style={iconStyle} /></a>
            <a target="_blank" rel="me" href="https://twitter.com/@avitrinesocial"><Icon type="twitter" style={iconStyle} /></a>
          </div>
        </div>
      </Col>
    </Row>
  </Layout.Footer>
);

export default Footer;
