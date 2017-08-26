import React from 'react'
import CategoriesItem from './CategoriesItem'
import {
  Container,
  Column,
  Columns,
  Title
} from 're-bulma'

const Categories = () => {

  const style = { padding: '10px 30px' };

  return (
    <Container>
      <Columns>
        <Column>
          <h2 className="containerTitle">
            <span>Doações por tipo</span>
          </h2>
        </Column>
      </Columns>
      <Columns style={style}>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
      </Columns>
      <Columns style={style}>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
        <Column>
          <CategoriesItem title="Title Categorie" image="http://placehold.it/200x200" />
        </Column>
      </Columns>
    </Container>
  )
}

export default Categories;
