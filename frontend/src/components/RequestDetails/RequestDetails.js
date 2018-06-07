import React from 'react';
import moment from 'moment';
import { Modal, Carousel } from 'antd';
import cx from 'classnames';
import styles from './styles.module.scss';
import ItemIndicator from '../../components/ItemIndicator';
import Arrow from '../../components/Arrow';
import ContactForm from '../../components/ContactForm';
import colors from '../../utils/styles/colors';
import Loading from '../Loading/Loading';

const mediaQuery = window.matchMedia('(min-width: 700px)');

const carouselSettings = {
  slidesToShow: 3,
  responsive: [
    {
      breakpoint: 801,
      settings: {
        slidesToShow: 2,
      },
    },
    {
      breakpoint: 696,
      settings: {
        slidesToShow: 1,
      },
    },
  ],
};

class RequestDetails extends React.Component {
  state = {
    contactFormVisible: false,
    previewVisible: false,
    previewImage: '',
  }

  showContactForm() {
    this.setState({
      contactFormVisible: true,
    });
  }

  cancelPreview = () => this.setState({ previewVisible: false })

  showPreview = (imgSource) => {
    this.setState({
      previewImage: imgSource,
      previewVisible: mediaQuery.matches,
    });
  }

  renderImages(images) {
    return (
      images.map(imageObj => (
        <a
          key={imageObj.id}
          onClick={() => this.showPreview(imageObj.url)}
          onKeyPress={() => this.showPreview(imageObj.url)}
          role="link"
          tabIndex={0}
        >
          <img
            src={imageObj.url}
            alt={imageObj.name}
          />
        </a>
      ))
    );
  }

  renderContent() {
    const { previewVisible, previewImage } = this.state;
    const { request } = this.props;

    if (!request) {
      return <Loading />;
    }

    if (this.state.contactFormVisible) {
      return (
        <ContactForm
          visible={this.state.contactFormVisible}
          onClick={() => this.setState({ contactFormVisible: false })}
        />
      );
    }

    return (
      <div className={styles.contentWrapper}>
        <div className={styles.itemDetails}>
          <ItemIndicator className={styles.itemIndicator} request={request} size="lg" />
          <div>
            <h1>{request.title}</h1>
            <p className={styles.updatedAt}>
              Atualizado em {
                moment(request.updatedAt).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
              }
            </p>
          </div>
        </div>
        <div className={styles.organizationBox}>
          <div className={styles.organizationBorder}>
            <p className={styles.organization}>{request.organization.name}</p>
            <p className={styles.description}>{request.description}</p>
          </div>
        </div>
        <div className={styles.arrowWrapper}>
          <Arrow size={32} color={colors.purple_400} onClick={() => this.carousel.prev()} left />
          <div className={styles.carouselWrapper}>
            <Carousel
              ref={(ref) => { this.carousel = ref; }}
              infinite={false}
              {...carouselSettings}
            >
              {this.renderImages(request.images)}
            </Carousel>
          </div>
          <Arrow size={32} color={colors.purple_400} onClick={() => this.carousel.next()} />
        </div>
        <div className={styles.buttonWrapper}>
          <button
            className={cx(styles.button, styles.helpButton)}
            onClick={() => this.showContactForm()}
          >
            QUERO AJUDAR!
          </button>
        </div>
        <Modal
          visible={previewVisible}
          footer={null}
          onCancel={this.cancelPreview}
        >
          <img className={styles.modalImage} alt="example" style={{ width: '100%' }} src={previewImage} />
        </Modal>
      </div>
    );
  }

  render() {
    return (
      <Modal
        visible={this.props.visible}
        footer={null}
        width={800}
        className={styles.modal}
        destroyOnClose
        onCancel={this.props.onCancel}
        closable={!this.state.contactFormVisible}
      >
        {this.renderContent()}
      </Modal>
    );
  }
}

export default RequestDetails;
