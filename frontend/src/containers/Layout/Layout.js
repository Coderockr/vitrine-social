import React from 'react';
import { Row, Col, Layout as AntLayout } from 'antd';
import Header from '../../components/Header';
import Search from '../../components/Search';
import OrganizationProfile from '../OrganizationProfile';

const { Content } = AntLayout;

const Layout = () => (
  <AntLayout>
    <Header />
    <Search />
    <Content>
      <Row>
        <Col
          xxl={{ span: 16, offset: 4 }}
          xl={{ span: 22, offset: 1 }}
        >
          <OrganizationProfile />
        </Col>
      </Row>
    </Content>
  </AntLayout>
);

export default Layout;
