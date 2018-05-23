import React from 'react';
import { Row, Col, Layout, Menu } from 'antd';
import { Link } from 'react-router-dom';
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
          <Menu.Item key="1">
            <Link to="/hello">Sobre o Projeto</Link>
          </Menu.Item>
          <Menu.Item key="2">
            <Link to="/hello">Quero Participar</Link>
          </Menu.Item>
          <Menu.Item key="3">
            <Link to="/hello">Contato</Link>
          </Menu.Item>
          <Menu.Item key="4">
            <Link to="/">Home</Link>
          </Menu.Item>
          <Menu.Item key="5">
            <Link to="/organization">Organization Profile</Link>
          </Menu.Item>
          <Menu.Item key="6">
            <Link to="/login">Login</Link>
          </Menu.Item>
        </Menu>
      </Col>
    </Row>
  </Layout.Header>
);

export default Header;
