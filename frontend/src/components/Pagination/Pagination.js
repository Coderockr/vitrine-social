import React from 'react';
import { Pagination as AntPagination } from 'antd';

import styles from './styles.module.scss';

const Pagination = () => (
  <div className={styles.wrapper}>
    <AntPagination total={50} />
  </div>
);

export default Pagination;
