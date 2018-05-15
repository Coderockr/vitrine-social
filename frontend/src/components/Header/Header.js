import React from 'react';
import { Row, Col, Layout, Menu } from 'antd';
import styles from './styles.module.scss';

const Header = () => (
  <Layout.Header className={styles.appHeader}>
    <Row>
      <Col span={12} offset={6}>
        <div className={styles.logo} />
        <Menu
          className={styles.appHeader}
          mode="horizontal"
          defaultSelectedKeys={['1']}
          theme="dark"
        >
          <Menu.Item key="1">Sobre o Projeto</Menu.Item>
          <Menu.Item key="2">Quero Participar</Menu.Item>
          <Menu.Item key="3">Contato</Menu.Item>
        </Menu>
      </Col>
    </Row>
  </Layout.Header>
);

export default Header;
