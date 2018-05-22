import React from 'react';
import Categories from '../../components/Categories';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';

const Home = () => (
  <Layout>
    <Categories />
    <Requests />
    <Pagination />
  </Layout>
);

export default Home;
