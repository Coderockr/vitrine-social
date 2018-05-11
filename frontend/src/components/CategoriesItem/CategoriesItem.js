import React from 'react';
import cx from 'classnames';
import './style.css';
import Icon from '../Icons';

const CategoriesItem = ({ image, title, active }) => (
  <div className="categoriesItem">
    <div className={cx('categoriesImage', { active })}>
      <Icon
        icon={image}
        size={70}
        color={active ? '#FFFFFF' : '#FF974B'}
      />
    </div>
    <div className={cx('categoriesCard', { active })}>
      <p className="categoriesTitle">{ title }</p>
    </div>
  </div>
);

export default CategoriesItem;
