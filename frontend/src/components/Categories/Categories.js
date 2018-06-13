import React from 'react';
import { Row, Col } from 'antd';
import CategoriesItem from '../../components/CategoriesItem';
import styles from './styles.module.scss';
import Loading from '../Loading/Loading';
import ErrorCard from '../../components/ErrorCard';

const Categories = ({
  loading,
  error,
  categories,
  hasSearch,
  onClick,
}) => (
  <div className={hasSearch ? styles.searchMargin : styles.categories}>
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
          {render(loading, error, categories, onClick)}
        </div>
      </Col>
    </Row>
  </div>
);

const render = (loading, error, categories, onClick) => {
  if (loading) {
    return <Loading />;
  }
  if (error) {
    return <ErrorCard text="Não foi possível listar as categorias de doações!" />;
  }

  return renderCategories(categories, onClick);
};

const renderCategories = (categories, onClick) => (
  categories.map(category => (
    <div key={category.id} className={styles.categoryWrapper}>
      <CategoriesItem
        title={category.name}
        image={category.slug}
        active={category.active}
        onClick={() => onClick(category.id)}
      />
    </div>
  ))
);

export default Categories;
