import React from 'react';
import { Row, Col } from 'antd';
import CategoriesItem from '../../components/CategoriesItem';
import styles from './styles.module.scss';

const allCategories = [
  { title: 'Alimentos', image: 'alimentos' },
  { title: 'Roupas', image: 'roupas' },
  { title: 'Brinquedos', image: 'brinquedos' },
  { title: 'Saúde', image: 'saude' },
  { title: 'Higiene Pessoal', image: 'higiene' },
  { title: 'Serviços', image: 'servicos', active: true },
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
  <div className={styles.categories}>
    <Row>
      <Col span={20} offset={2}>
        <h2 className={styles.containerTitle}>
          <span>Doações por tipo</span>
        </h2>
      </Col>
    </Row>
    <Row>
      <Col span={22} offset={1}>
        <div className={styles.categoriesWrapper}>
          {renderCategories(allCategories)}
        </div>
      </Col>
    </Row>
  </div>
);

const renderCategories = categories => (
  categories.map(category => (
    <div className={styles.categoryWrapper}>
      <CategoriesItem
        title={category.title}
        image={category.image}
        active={category.active}
      />
    </div>
  ))
);

export default Categories;
