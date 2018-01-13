import React from 'react'
import {
  Card,
  CardImage,
  CardContent,
  Content ,
  Image,
  Title
} from 're-bulma';
import cx from 'classnames';
import style from  './style.css';
import Icon from '../Icons';

const CategoriesItem = ({ image, title, active }) => {
  return (
    <Card isFullwidth className="categoriesItem">
      <CardImage className={cx('categoriesImage', { active })}>
        <Icon
          icon={ image }
          size={70}
          color={active ? '#FFFFFF' : '#FF974B'}
        />
      </CardImage>
      <CardContent className={cx('categoriesCard', { active })}>
        <Content className="categoriesContainer">
          <p className="categoriesTitle">{ title }</p>
        </Content>
      </CardContent>
    </Card>
  )
}

export default CategoriesItem;
