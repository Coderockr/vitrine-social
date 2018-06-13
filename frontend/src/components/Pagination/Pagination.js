import React from 'react';
import { Pagination as AntPagination } from 'antd';

import styles from './styles.module.scss';

const Pagination = ({ current, total, onChange }) => (
  <div className={styles.wrapper} hidden={total <= 10}>
    <AntPagination
      current={current}
      total={total}
      onChange={onChange}
    />
  </div>
);

export default Pagination;
