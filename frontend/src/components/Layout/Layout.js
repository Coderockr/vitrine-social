import React from 'react';
import { Row, Col, Layout as AntLayout } from 'antd';
import Header from '../Header';
import Search from '../Search';

const { Content } = AntLayout;

const Layout = ({ children }) => (
  <AntLayout>
    <Header />
    <Search />
    <Content>
      <Row>
        <Col
          xxl={{ span: 16, offset: 4 }}
          xl={{ span: 22, offset: 1 }}
        >
          { children }
        </Col>
      </Row>
    </Content>
  </AntLayout>
);

export default Layout;
