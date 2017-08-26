import React from 'react'
import {
  Container,
  Column,
  Columns,
  Title
} from 're-bulma'
import CategoriesItem from '../../components/CategoriesItem'

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
          <CategoriesItem title="Artigos Domésticos" image="alimentos" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="brinquedos" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="construcao" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="domesticos" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="equipamentos" />
        </Column>
      </Columns>
      <Columns style={style}>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="higiene" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="livros" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="moveis" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="outros" />
        </Column>
        <Column>
          <CategoriesItem title="Artigos Domésticos" image="papelaria" />
        </Column>
      </Columns>
    </Container>
  )
}

export default Categories;
