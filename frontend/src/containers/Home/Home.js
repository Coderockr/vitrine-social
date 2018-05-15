import React from 'react';
import { Layout } from 'antd';
import Categories from '../../components/Categories';
import Pagination from '../../components/Pagination';
import Search from '../../components/Search';
import Requests from '../../components/Requests';
import RequestDetails from '../RequestDetails';

const { Content } = Layout;

const Home = () => (
  <Content>
    <Search />
    <Categories />
    <Requests />
    <Pagination />
    <RequestDetails visible />
  </Content>
);

export default Home;
