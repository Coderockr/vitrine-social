import React from 'react';
import Categories from '../../components/Categories';
import Header from '../../components/Header';
import Pagination from '../../components/Pagination';
import Search from '../../components/Search';
import Requests from '../../components/Requests';
import Dialog from '../../components/Dialog';

const App = () => (
  <div>
    <Header />
    <Search />
    <Categories />
    <Requests />
    <Pagination />
    <Dialog active>
      <h1>Teste</h1>
    </Dialog>
  </div>
);

export default App;
