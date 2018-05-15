import React from 'react';
import { Layout as AntLayout } from 'antd';
import Header from '../../components/Header';
import Home from '../Home';
import OrganizationProfile from '../OrganizationProfile';

const Layout = () => (
  <AntLayout>
    <Header />
    <Home />
    <OrganizationProfile />
  </AntLayout>
);

export default Layout;
