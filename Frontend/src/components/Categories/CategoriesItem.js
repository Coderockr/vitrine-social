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

const CategoriesItem = ({ image, title }) => {
  return (
    <Card isFullwidth className="categoriesItem">
      <CardImage>
        <Image src={ image } ratio="isSquare"/>
      </CardImage>
      <CardContent>
        <Content className="categoriesContainer">
          <p className="categoriesTitle">{ title }</p>
        </Content>
      </CardContent>
    </Card>
  )
}

export default CategoriesItem;
