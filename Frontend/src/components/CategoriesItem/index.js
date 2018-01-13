import React from 'react'
import {
  Card,
  CardImage,
  CardContent,
  Content ,
  Image,
  Title
} from 're-bulma'

import style from  './style.css';
import Icon from '../Icons';

const CategoriesItem = ({ image, title }) => {
  return (
    <Card isFullwidth className="categoriesItem">
      <CardImage className="categoriesImage">
        <Icon
          icon={ image }
          size={70}
          color='#FF974B'
        />
      </CardImage>
      <CardContent className="categoriesCard">
        <Content className="categoriesContainer">
          <p className="categoriesTitle">{ title }</p>
        </Content>
      </CardContent>
    </Card>
  )
}

export default CategoriesItem;
