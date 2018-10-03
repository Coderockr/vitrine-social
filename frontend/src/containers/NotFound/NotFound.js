import React from 'react';
import { Row, Col } from 'antd';
import { Link } from 'react-router-dom';
import Layout from '../../components/Layout';
import styles from './styles.module.scss';

const NotFound = () => (
  <Layout>
    <Row>
      <Col
        xl={{ span: 14, offset: 5 }}
        lg={{ span: 16, offset: 4 }}
        md={{ span: 20, offset: 2 }}
        sm={{ span: 22, offset: 1 }}
        xs={{ span: 22, offset: 1 }}
      >
        <div className={styles.sectionWrapper}>
          <h1 className={styles.title}>404</h1><br />
          <h2>
            A página que voce esta procurando não pode ser encontrada.<br />
            <Link to="/" className={styles.link}>Clique aqui</Link> para retornar para a página inicial.
          </h2>
        </div>
      </Col>
    </Row>
  </Layout>
);

export default NotFound;
