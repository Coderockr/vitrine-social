import React from 'react';
import Categories from '../../components/Categories';
import Pagination from '../../components/Pagination';
import Requests from '../../components/Requests';
import RequestDetails from '../RequestDetails';

const Home = () => (
  <div>
    <Categories />
    <Requests />
    <Pagination />
    <RequestDetails visible />
  </div>
);

export default Home;
