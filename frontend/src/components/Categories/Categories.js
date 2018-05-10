import React from 'react';
import { Row, Col } from 'antd';
import CategoriesItem from '../../components/CategoriesItem';

import './style.css';

const allCategories = [
  { title: 'Alimentos', image: 'alimentos' },
  { title: 'Roupas', image: 'roupas' },
  { title: 'Brinquedos', image: 'brinquedos' },
  { title: 'Saúde', image: 'saude' },
  { title: 'Higiene Pessoal', image: 'higiene' },
  { title: 'Serviços', image: 'servicos' },
  { title: 'Materiais de Construção', image: 'construcao' },
  { title: 'Voluntários', image: 'voluntarios' },
  { title: 'Móveis', image: 'moveis' },
  { title: 'Equipamentos', image: 'equipamentos' },
  { title: 'Artigos Domésticos', image: 'domesticos' },
  { title: 'Livros', image: 'livros' },
  { title: 'Papelaria', image: 'papelaria' },
  { title: 'Pets', image: 'pet' },
  { title: 'Outros', image: 'outros' },
];

const Categories = () => (
  <div className="categories">
    <Row>
      <Col span={20} offset={2}>
        <h2 className="containerTitle">
          <span>Doações por tipo</span>
        </h2>
      </Col>
    </Row>
    <Row className="row">
      <Col span={20} offset={2}>
        <div className="categoriesWrapper">
          {renderCategories(allCategories)}
        </div>
      </Col>
    </Row>
  </div>
);

const renderCategories = categories => (
  categories.map(category => (
    <div className="categoryWrapper">
      <CategoriesItem title={category.title} image={category.image} />
    </div>
  ))
);

export default Categories;
