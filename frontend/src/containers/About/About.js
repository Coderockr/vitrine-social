import React from 'react';
import { Row, Col } from 'antd';
import cx from 'classnames';
import Layout from '../../components/Layout';
import styles from './styles.module.scss';

const itemSections = [
  {
    number: 1,
    text: 'Ao acessar a Vitrine Social o doador visualiza as solicitações de doações de diversas entidades beneficentes em um só lugar, podendo ele percorrer pelas mais recentes, filtar por categorias, ou mesmo buscar por solicitações ou entidades específicas.',
    style: styles.itemAmbar,
  },
  {
    number: 2,
    text: 'Ao clicar no botão “mais detalhes”, são exibidas todas as demais informações sobre a solicitação de doação,  para que o doador saiba mais sobre a natureza da necessidade, e como os recursos doados serão utilizados.',
    style: styles.itemCian,
  },
  {
    number: 3,
    text: 'Ao  clicar em “quero ajudar” o doador pode optar por entar em contato diretamente com a entidade beneficente por telefone, ou então deixar uma mensagem através da plataforma, para posteriormente tratarem sobre os detalhes logisticos da doação.',
    style: styles.itemPurple,
  },
];

const About = () => (
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
          <h1 className={styles.title}>CADA DOAÇÃO FAZ TODA A DIFERENÇA</h1>
          <p className={styles.mainText}>
          Ajudar o próximo nem sempre é algo fácil. Muitas vezes, apesar de querermos colaborar
          com alguma causa social, somos desestimulados a ajudar, seja por desconhecer as entidades
          beneficentes mais necessitadas, ou também pela dificuldade de entrar em contato e informar
          sobre qual é a melhor forma de ajudar. Para simplificar este processo,
          criamos a Vitrine Social.
          </p>
          <img src="./assets/images/brand.svg" alt="brand" />
          <img src="./assets/images/hands.svg" alt="giving" />
        </div>
      </Col>
    </Row>
    <Row>
      <Col span={20} offset={2}>
        <h2 className={cx(styles.containerTitle, styles.sectionWrapper)}>
          <span>{'COMO FUNCIONA?'}</span>
        </h2>
        {renderItemSections()}
      </Col>
    </Row>
  </Layout>
);

const renderItemSections = () => (
  itemSections.map(item => (
    <div className={styles.itemSection}>
      <div className={cx(styles.item, item.style)}>
        <p className={cx(styles.itemNumber, item.style)}>{item.number}</p>
      </div>
      <p className={styles.mainText}>{item.text}</p>
    </div>
  ))
);

export default About;
