import React from 'react';
import { Row, Col } from 'antd';
import CategoriesItem from '../../components/CategoriesItem';
import styles from './styles.module.scss';
import Loading from '../Loading/Loading';

const Categories = ({ loading, categories, hasSearch }) => (
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
          {loading ? <Loading /> : renderCategories(categories)}
        </div>
      </Col>
    </Row>
  </div>
);

const renderCategories = categories => (
  categories.map(category => (
    <div key={category.id} className={styles.categoryWrapper}>
      <CategoriesItem
        title={category.name}
        image={category.slug}
        active={category.active}
      />
    </div>
  ))
);

export default Categories;
