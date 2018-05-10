import React from 'react';
import Categories from '../../components/Categories';
import Header from '../../components/Header';
import Pagination from '../../components/Pagination';
import Search from '../../components/Search';
import Requests from '../../components/Requests';

const App = () => (
  <div>
    <Header />
    <Search />
    <Categories />
    <Requests />
    <Pagination />
  </div>
);

export default App;
