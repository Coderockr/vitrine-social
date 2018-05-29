import React from 'react';
import { storiesOf } from '@kadira/storybook';

import Categories from '../../components/Categories';
import Header from '../../components/Header';
import Pagination from '../../components/Pagination';
import Search from '../../components/Search';
import Requests from '../../components/Requests';
import Dialog from '../../components/Dialog';

storiesOf('Home', module)
  .add('Default View', () => (
    <div>
      <Header />
      <Search />
      <Categories />
      <Requests />
      <Pagination />
      <Dialog />
    </div>
  ));
