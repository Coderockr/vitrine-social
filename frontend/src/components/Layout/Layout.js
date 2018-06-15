import React from 'react';
import { Row, Col, Layout as AntLayout } from 'antd';
import Header from '../Header';
import Footer from '../Footer';

const { Content } = AntLayout;

const Layout = ({ children }) => (
  <AntLayout style={{ height: "100vh" }}>
    <Header />
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
    <Footer />
  </AntLayout>
);

export default Layout;
