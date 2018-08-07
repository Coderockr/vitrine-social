import React from 'react';
import cx from 'classnames';
import { Row, Col, Carousel, Avatar, Icon } from 'antd';
import Img from 'react-image';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import Arrow from '../../components/Arrow';
import ChangePassword from '../../components/ChangePassword';
import OrganizationProfileForm from '../../components/OrganizationProfileForm';
import { maskPhone } from '../../utils/mask';
import colors from '../../utils/styles/colors';
import { api } from '../../utils/api';
import { getUser, updateUser } from '../../utils/auth';
import styles from './styles.module.scss';
import Loading from '../../components/Loading';
import ErrorCard from '../../components/ErrorCard';

const carouselSettings = {
  slidesToShow: 1,
};

const mediaQuery = window.matchMedia('(min-width: 700px)');

class OrganizationProfile extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      arrowSize: mediaQuery.matches ? 60 : 32,
      isOrganization: false,
      editProfileVisible: false,
      changePasswordVisible: false,
      saveEnabled: false,
    };

    mediaQuery.addListener(this.widthChange.bind(this));
  }

  componentWillMount() {
    this.fetchData();
  }

  componentDidMount() {
    document.title = 'Vitrine Social - Perfil da Organização';
  }

  componentWillUnmount() {
    mediaQuery.removeListener(this.widthChange);
  }

  widthChange() {
    this.setState({
      arrowSize: mediaQuery.matches ? 60 : 32,
    });
  }

  activeStatusFilter(request) {
    return request.status === 'ACTIVE';
  }

  inactiveStatusFilter(request) {
    return request.status === 'INACTIVE';
  }

  filterRequestsByStatus(requests, active) {
    if (active) {
      return requests.filter(this.activeStatusFilter);
    }
    return requests.filter(this.inactiveStatusFilter);
  }

  fetchData(onSaveOrganization) {
    const user = getUser();
    const { match: { params } } = this.props;
    api.get(`organization/${params.organizationId}`).then(
      (response) => {
        this.setState({
          organization: response.data,
          activeRequests: this.filterRequestsByStatus(response.data.needs, true),
          inactiveRequests: this.filterRequestsByStatus(response.data.needs, false),
          isOrganization: user ? user.id === response.data.id : false,
          loading: false,
        });
        if (onSaveOrganization) {
          updateUser(response.data);
        }
      }, (error) => {
        this.setState({
          loading: false,
          error,
        });
      },
    );
  }

  renderImages(images) {
    return (
      images.map(image => (
        <Img
          key={image.id}
          src={image.url}
          alt={image.name}
          loader={<Icon type="loading" style={{ fontSize: 60, color: colors.teal_400 }} />}
        />
      ),
      )
    );
  }

  renderOrganizationInfo() {
    const {
      loading,
      error,
      isOrganization,
      organization,
      editProfileVisible,
      saveEnabled,
      changePasswordVisible,
      arrowSize,
    } = this.state;

    if (loading) {
      return <Loading />;
    }
    if (error) {
      return <ErrorCard text="Não foi possível carregar os dados da Organização!" />;
    }

    const { address } = organization;
    const addressString = `${address.street} ${address.number}, ${address.complement ? `${address.complement},` : ''} Bairro ${address.neighborhood}, ${address.city} - ${address.state}`;

    return (
      <div>
        {isOrganization &&
          <div className={styles.buttonWrapper}>
            <button
              className={styles.editButton}
              onClick={() => this.setState({ editProfileVisible: true })}
            >
              EDITAR
            </button>
            <OrganizationProfileForm
              visible={editProfileVisible}
              onCancel={() => this.setState({ editProfileVisible: false, saveEnabled: false })}
              onSave={() => this.fetchData(true)}
              saveEnabled={saveEnabled}
              enableSave={enable => this.setState({ saveEnabled: enable })}
              organization={organization}
            />
            <button
              className={styles.editButton}
              onClick={() => this.setState({ changePasswordVisible: true })}
            >
              ALTERAR SENHA
            </button>
            <ChangePassword
              modal
              user
              visible={changePasswordVisible}
              onCancel={() => this.setState({ changePasswordVisible: false })}
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
          {(organization.about || organization.website) &&
            <div className={cx(styles.border, styles.aboutBorder)}>
              <h1>Sobre</h1>
              {organization.about &&
                <p style={{ 'white-space': 'pre-line' }}>{organization.about}</p>
              }
              {organization.website &&
              <a target="_blank" rel="me" href={organization.website.includes('http') ? organization.website : `//${organization.website}`}>{organization.website}</a>
              }
            </div>
          }
          <div className={cx(styles.border, styles.phoneBorder)}>
            <h1>Telefone</h1>
            <a href={`tel:${organization.phone}`}>{maskPhone(organization.phone)}</a>
          </div>
          <div className={cx(styles.border, styles.addressBorder)}>
            <h1>Endereço</h1>
            <a target="_blank" rel="me" href={`https://maps.google.com/?q=${addressString}`}>{addressString}</a>
          </div>
          {organization.images.length > 0 &&
            <div>
              <div className={cx(styles.border, styles.imagesBorder)}>
                <h1>Imagens da Organização</h1>
              </div>
              <div className={styles.arrowWrapper}>
                <Arrow
                  size={arrowSize}
                  color={colors.teal_400}
                  onClick={() => this.carousel.prev()}
                  left
                  over
                  hidden={organization.images.length < 2}
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
                  size={arrowSize}
                  color={colors.teal_400}
                  onClick={() => this.carousel.next()}
                  over
                  hidden={organization.images.length < 2}
                />
              </div>
            </div>
          }
        </Col>
      </div>
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
              {this.renderOrganizationInfo()}
            </div>
          </Col>
        </Row>
        <Requests
          isOrganization={this.state.isOrganization}
          loading={this.state.loading}
          activeRequests={this.state.loading ? null : this.state.activeRequests}
          inactiveRequests={this.state.loading ? null : this.state.inactiveRequests}
          onSave={() => this.fetchData()}
          error={this.state.error}
        />
      </Layout>
    );
  }
}

export default OrganizationProfile;
