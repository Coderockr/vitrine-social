import React from 'react';
import cx from 'classnames';
import { Row, Col, Carousel, Avatar } from 'antd';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import Arrow from '../../components/Arrow';
import OrganizationProfileEdit from '../OrganizationProfileEdit';
import RequestDetailsEdit from '../../components/RequestDetailsEdit';
import RequestDetails from '../../components/RequestDetails';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

const organization = {
  name: 'Lar Abdon Batista',
  about: 'O Lar Abdon Batista foi criado em 1911, pelo então Prefeito Abdon Batista. Inicialmente foi chamado de Sociedade de Caridade e Asylo de Órfãos e Desvalidos e funcionava no prédio que hoje abriga a Secretaria Municipal de Assistência Social, na rua Procópio Gomes, no bairro Bucarein.',
  link: 'http://coderockr.com/',
  phoneNumber: '(47) 3227-6359',
  address: 'Rua Presidente Affonso Penna, 680 - Bucarein, Joinville - SC',
  images: [
    { title: 'Leitura Infantil', image: <img src="assets/images/leitura-infantil.jpg" alt="Leitura Infantil" /> },
    { title: 'Leitura Infantil 2', image: <img src="assets/images/leitura-infantil 2.jpg" alt="Leitura Infantil 2" /> },
    { title: 'Leitura Infantil 3', image: <img src="assets/images/leitura-infantil 3.jpg" alt="Leitura Infantil 3" /> },
    { title: 'Leitura Infantil 4', image: <img src="assets/images/leitura-infantil.jpg" alt="Leitura Infantil 4" /> },
    { title: 'Leitura Infantil 5', image: <img src="assets/images/leitura-infantil 2.jpg" alt="Leitura Infantil 5" /> },
    { title: 'Leitura Infantil 6', image: <img src="assets/images/leitura-infantil 3.jpg" alt="Leitura Infantil 6" /> },
  ],
};

const carouselSettings = {
  slidesToShow: 1,
};

const mediaQuery = window.matchMedia('(min-width: 700px)');

class OrganizationProfile extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      arrowSize: mediaQuery.matches ? 60 : 32,
      isOrganization: true,
      editProfileVisible: false,
      editRequestVisible: false,
      requestDetailsVisible: false,
    };

    mediaQuery.addListener(this.widthChange.bind(this));
  }

  componentWillUnmount() {
    mediaQuery.removeListener(this.widthChange);
  }

  widthChange() {
    this.setState({
      arrowSize: mediaQuery.matches ? 60 : 32,
    });
  }

  renderImages(images) {
    return (
      images.map(imageObj => imageObj.image)
    );
  }

  render() {
    return (
      <Layout>
        <Row>
          <Col
            xl={{ span: 20, offset: 2 }}
            xs={{ span: 22, offset: 1 }}
          >
            <div className={styles.profileWrapper}>
              <h2 className={styles.containerTitle}>
                <span>PERFIL DA ORGANIZAÇÃO</span>
              </h2>
              <div className={styles.buttonWrapper} hidden={!this.state.isOrganization}>
                <button
                  className={styles.button}
                  onClick={() => this.setState({ editProfileVisible: true })}
                >
                  EDITAR
                </button>
              </div>
              <Avatar
                src="assets/images/leitura-infantil 3.jpg"
                size={'large'}
                style={{ marginTop: 20 }}
              />
              <h1 className={styles.organizationName}>
                <span>{organization.name}</span>
              </h1>
              <Col
                sm={{ span: 18, offset: 3 }}
                xs={{ span: 24, offset: 0 }}
              >
                <div className={cx(styles.border, styles.aboutBorder)}>
                  <h1>Sobre</h1>
                  <p>{organization.about}</p>
                </div>
                <div className={cx(styles.border, styles.phoneBorder)}>
                  <h1>Telefone</h1>
                  <a>{organization.phoneNumber}</a>
                </div>
                <div className={cx(styles.border, styles.addressBorder)}>
                  <h1>Endereço</h1>
                  <a>{organization.address}</a>
                </div>
                <div className={cx(styles.border, styles.imagesBorder)}>
                  <h1>Imagens da Organização</h1>
                </div>
                <div className={styles.arrowWrapper}>
                  <Arrow
                    size={this.state.arrowSize}
                    color={colors.teal_400}
                    onClick={() => this.carousel.prev()}
                    left
                    over
                  />
                  <div className={styles.carouselWrapper}>
                    <Carousel
                      ref={(ref) => { this.carousel = ref; }}
                      infinite={false}
                      {...carouselSettings}
                    >
                      {this.renderImages(organization.images)}
                    </Carousel>
                  </div>
                  <Arrow
                    size={this.state.arrowSize}
                    color={colors.teal_400}
                    onClick={() => this.carousel.next()}
                    over
                  />
                </div>
              </Col>
            </div>
          </Col>
        </Row>
        <Requests
          isOrganization={this.state.isOrganization}
          onEdit={request => this.setState({ editRequestVisible: true, request })}
          onClick={request => this.setState({ requestDetailsVisible: true, request })}
        />
        <Pagination />
        <OrganizationProfileEdit
          visible={this.state.editProfileVisible}
          onCancel={() => this.setState({ editProfileVisible: false })}
        />
        <RequestDetails
          visible={this.state.requestDetailsVisible}
          request={this.state.request}
          onCancel={() => this.setState({ requestDetailsVisible: false })}
        />
        <RequestDetailsEdit
          visible={this.state.editRequestVisible}
          request={this.state.request}
          onCancel={() => this.setState({ editRequestVisible: false })}
        />
      </Layout>
    );
  }
}

export default OrganizationProfile;
