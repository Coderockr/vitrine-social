import React from 'react'
import {
  Card,
  CardImage,
  CardContent,
  Content ,
  Image,
  Title
} from 're-bulma'

const CategoriesItem = ({ image, title }) => {
  return (
    <Card isFullwidth>
      <CardImage>
        <Image src={ image } ratio="isSquare"/>
      </CardImage>
      <CardContent>
        <Content>
          <Title size="is5">{ title }</Title>
        </Content>
      </CardContent>
    </Card>
  )
}

export default CategoriesItem;
