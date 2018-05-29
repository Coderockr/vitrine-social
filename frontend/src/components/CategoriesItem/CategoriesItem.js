import React from 'react';
import cx from 'classnames';
import styles from './styles.module.scss';
import Icon from '../Icons';
import { ambar_400, white } from '../../utils/styles/colors.js';

const CategoriesItem = ({ image, title, active }) => (
  <div className={cx(styles.categoriesItem, { [styles.active]: active })}>
    <div className={styles.categoriesImage}>
      <Icon
        icon={image}
        size={70}
        color={active ? white : ambar_400}
      />
    </div>
    <div className={styles.categoriesCard}>
      <p className={styles.categoriesTitle}>{ title }</p>
    </div>
  </div>
);

export default CategoriesItem;
