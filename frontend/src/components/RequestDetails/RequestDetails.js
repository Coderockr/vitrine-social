import React from 'react';
import moment from 'moment';
import { Modal, Carousel } from 'antd';
import styles from './styles.module.scss';
import ItemIndicator from '../../components/ItemIndicator';
import Arrow from '../../components/Arrow';
import ContactForm from '../../components/ContactForm';
import colors from '../../utils/styles/colors';

const mediaQuery = window.matchMedia('(min-width: 700px)');

const request = {
  organization: {
    name: 'Lar Abdon batista',
    link: 'http://coderockr.com/',
  },
  category: 'voluntarios',
  data: moment().subtract(2, 'days'),
  item: '10 voluntários para ler para criancinhas felizes',
  description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
};

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

const imagesArray = [
  { title: 'Leitura Infantil', src: 'assets/images/leitura-infantil.jpg' },
  { title: 'Leitura Infantil 2', src: 'assets/images/leitura-infantil 2.jpg' },
  { title: 'Leitura Infantil 3', src: 'assets/images/leitura-infantil 3.jpg' },
  { title: 'Leitura Infantil 4', src: 'assets/images/leitura-infantil.jpg' },
  { title: 'Leitura Infantil 5', src: 'assets/images/leitura-infantil 2.jpg' },
  { title: 'Leitura Infantil 6', src: 'assets/images/leitura-infantil 3.jpg' },
];

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
          onClick={() => this.showPreview(imageObj.src)}
          onKeyPress={() => this.showPreview(imageObj.src)}
          role="link"
          tabIndex={0}
        >
          <img
            src={imageObj.src}
            alt="Leitura Infantil"
          />
        </a>
      ))
    );
  }

  renderContent() {
    const { previewVisible, previewImage } = this.state;

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
            <h1>{request.item}</h1>
            <p className={styles.updatedAt}>
              Atualizado em {
                moment(request.data).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
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
              {this.renderImages(imagesArray)}
            </Carousel>
          </div>
          <Arrow size={32} color={colors.purple_400} onClick={() => this.carousel.next()} />
        </div>
        <button
          onClick={() => this.showContactForm()}
        >
          QUERO AJUDAR!
        </button>
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
