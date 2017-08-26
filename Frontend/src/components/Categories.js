import React from 'react'
import CategoriesItem from './CategoriesItem'
import {
  Container,
  Column,
  Columns
} from 're-bulma'

const Categories = () => {
  return (
    <Container>
      <Columns>
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
      <Columns>
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
