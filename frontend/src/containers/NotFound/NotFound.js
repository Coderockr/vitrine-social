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
            O link que você clicou parece estar quebrado ou a página pode ter sido removida.
          </h2>
          <h3>
            Visite nossa <Link to="/" className={styles.link}>Página Inicial</Link> ou
            &nbsp;<Link to="/contato" className={styles.link}>entre em contato</Link> conosco
            para relatar um problema.
          </h3>
        </div>
      </Col>
    </Row>
  </Layout>
);

export default NotFound;
