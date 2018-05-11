import React from 'react';
import Categories from '../../components/Categories';
import Header from '../../components/Header';
import Pagination from '../../components/Pagination';
import Search from '../../components/Search';
import Requests from '../../components/Requests';
import RequestDetails from '../RequestDetails';

const App = () => (
  <div>
    <Header />
    <Search />
    <Categories />
    <Requests />
    <Pagination />
    <RequestDetails visible />
  </div>
);

export default App;
