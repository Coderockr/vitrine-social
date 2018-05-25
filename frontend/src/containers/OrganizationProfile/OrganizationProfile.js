import React from 'react';
import cx from 'classnames';
import { Row, Col, Carousel, Avatar } from 'antd';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import Arrow from '../../components/Arrow';
import OrganizationProfileForm from '../../components/OrganizationProfileForm';
import { maskPhone } from '../../utils/mask';
import colors from '../../utils/styles/colors';
import api from '../../utils/api';
import styles from './styles.module.scss';

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
    };

    mediaQuery.addListener(this.widthChange.bind(this));
  }

  componentWillMount() {
    this.fetchData();
  }

  componentWillUnmount() {
    mediaQuery.removeListener(this.widthChange);
  }

  widthChange() {
    this.setState({
      arrowSize: mediaQuery.matches ? 60 : 32,
    });
  }

  fetchData() {
    api.get('organization/1').then(
      (response) => {
        this.setState({ organization: response.data });
      },
    );
  }

  renderImages(images) {
    return (
      images.map(image => <img key={image.id} src={image.url} alt={image.name} />)
    );
  }

  render() {
    const { organization } = this.state;
    if (!organization) {
      return null;
    }

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
              {this.state.isOrganization &&
                <div className={styles.buttonWrapper}>
                  <button
                    className={styles.button}
                    onClick={() => this.setState({ editProfileVisible: true })}
                  >
                    EDITAR
                  </button>
                  <OrganizationProfileForm
                    visible={this.state.editProfileVisible}
                    onCancel={() => this.setState({ editProfileVisible: false })}
                    organization={organization}
                  />
                </div>
              }
              <Avatar
                src={organization.logo}
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
                  <p>{organization.resume}</p>
                </div>
                <div className={cx(styles.border, styles.phoneBorder)}>
                  <h1>Telefone</h1>
                  <a>{maskPhone(organization.phone)}</a>
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
        <Requests isOrganization={this.state.isOrganization} />
        <Pagination />
      </Layout>
    );
  }
}

export default OrganizationProfile;
