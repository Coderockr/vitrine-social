import React from 'react';
import { Row, Col } from 'antd';
import cx from 'classnames';
import { Link } from 'react-router-dom';
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
      <Col span={20} offset={2} className={styles.sectionWrapper}>
        <h2 className={styles.containerTitle}>
          <span>{'COMO FUNCIONA?'}</span>
        </h2>
        {renderItemSections()}
        {renderLastSection()}
        <img className={styles.image} src="./assets/images/community.svg" alt="community" />
      </Col>
    </Row>
    <Row>
      <Col span={20} offset={2} className={styles.lastSection}>
        <h2 className={styles.containerTitle}>
          <span>{'NOSSA HISTÓRIA'}</span>
        </h2>
        <a target="_blank" rel="me" href="http://www.coderockr.com">
          <img className={styles.logo} src="./assets/images/coderockr.svg" alt="coderockr" />
        </a>
        <div className={styles.mainText}>
          <p>
          O projeto Vitrine Social foi desenvolvido pela equipe da Coderockr durante os
          Coderockr Jams, eventos internos da Coderockr realizados com o propósito de aprender sobre
          novas tecnologias e metodologias, e também onde os colaboradores podem compartilhar
          seus conhecimentos com o restante da equipe.
          </p>
          <p>
          Em um desses eventos decidimos trabalhar em uma ideia para ajudar a comunidade, e então
          surgiu o Vitrine Social. Começamos a análise e discussão do projeto testando a
          metodologia <a target="_blank" rel="me" href="https://blog.coderockr.com/modelando-sistemas-usando-event-storming-1e18e6563eaa">Event Storming</a> e
          continuamos com o desenvolvimento a partir do que definimos neste processo.
          </p>
          <p>
          Pretendemos dar continuidade ao projeto e criar novas funcionalidades de acordo com as
          sugestões das entidades e dos doadores. Se você tiver alguma sugestão <Link to="/contact">entre em contato!</Link>
          </p>
        </div>
      </Col>
    </Row>
  </Layout>
);

const renderItemSections = () => (
  itemSections.map(item => (
    <div className={styles.itemSection}>
      <div className={cx(styles.item, item.style)}>
        <span className={cx(styles.itemNumber, item.style)}>{item.number}</span>
      </div>
      <p>{item.text}</p>
    </div>
  ))
);

const renderLastSection = () => (
  <div className={styles.itemSection}>
    <div className={cx(styles.itemPink, styles.item)}>
      <img className={styles.itemIcon} src="./assets/images/heart.svg" alt="heart" />
    </div>
    <p>
    Pronto! Doando desta forma, além poupar o seu tempo buscando entidades confiáveis, você ainda
    tem a certeza de que os recursos doados serão recebidos e utilizados de forma eficiente pelas
    entidades que mais os necessitam! Seja uma entidade grande e de renome, ou aquela pequena
    entidade que você ainda não conhecia.
    </p>
  </div>
);

export default About;
