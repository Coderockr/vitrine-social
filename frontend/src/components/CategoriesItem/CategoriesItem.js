import React from 'react';
import cx from 'classnames';
import styles from './styles.module.scss';
import Icon from '../Icons';

const CategoriesItem = ({
  image,
  title,
  active,
  onClick,
}) => (
  <div className={cx(styles.categoriesItem, { [styles.active]: active })}>
    <button className={styles.categoriesImage} onClick={onClick}>
      <Icon
        className={styles.categoriesIcon}
        icon={image}
        size={70}
        color="currentColor"
      />
    </button>
    <div className={styles.categoriesCard}>
      <p className={styles.categoriesTitle}>{ title }</p>
    </div>
  </div>
);

export default CategoriesItem;
