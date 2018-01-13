import React from 'react';
import {
  Container,
  Column,
  Columns,
  Section,
  Title
} from 're-bulma';
import CategoriesItem from '../../components/CategoriesItem';

import './style.css';

const Categories = () => {
  return (
    <Section className="categories">
      <Container>
        <Columns>
          <Column>
            <h2 className="containerTitle">
              <span>Doações por tipo</span>
            </h2>
          </Column>
        </Columns>

        <Columns className="row">
          <Column>
            <CategoriesItem title="Alimentos" image="alimentos" />
          </Column>
          <Column>
            <CategoriesItem title="Roupas" image="roupas" />
          </Column>
          <Column>
            <CategoriesItem title="Brinquedos" image="brinquedos" active />
          </Column>
          <Column>
            <CategoriesItem title="Saúde" image="saude" />
          </Column>
          <Column>
            <CategoriesItem title="Higiene Pessoal" image="higiene" />
          </Column>
        </Columns>

        <Columns className="row">
          <Column>
            <CategoriesItem title="Serviços" image="servicos" />
          </Column>
          <Column>
            <CategoriesItem title="Materiais de Construção" image="construcao" />
          </Column>
          <Column>
            <CategoriesItem title="Voluntários" image="voluntarios" />
          </Column>
          <Column>
            <CategoriesItem title="Móveis" image="moveis" />
          </Column>
          <Column>
            <CategoriesItem title="Equipamentos" image="equipamentos" />
          </Column>
        </Columns>

        <Columns className="row">
          <Column>
            <CategoriesItem title="Artigos Domésticos" image="domesticos" />
          </Column>
          <Column>
            <CategoriesItem title="Livros" image="livros" />
          </Column>
          <Column>
            <CategoriesItem title="Papelaria" image="papelaria" />
          </Column>
          <Column>
            <CategoriesItem title="Pets" image="pet" />
          </Column>
          <Column>
            <CategoriesItem title="Outros" image="outros" />
          </Column>
        </Columns>
      </Container>
    </Section>
  )
}

export default Categories;
