import React from 'react';
import cx from 'classnames';
import { Pagination as AntPagination } from 'antd';

import './style.css';

const Pagination = () => (
  <div className={cx('wrapper')}>
    <AntPagination className="mypagination" total={50} />
  </div>
);

export default Pagination;
