import React from 'react';
import moment from 'moment';
import { Modal, Carousel, Icon } from 'antd';
import Img from 'react-image';
import cx from 'classnames';
import ReactGA from 'react-ga';
import styles from './styles.module.scss';
import ItemIndicator from '../../components/ItemIndicator';
import Arrow from '../../components/Arrow';
import ContactForm from '../../components/ContactForm';
import colors from '../../utils/styles/colors';
import Loading from '../Loading/Loading';

const mediaQuery = window.matchMedia('(min-width: 700px)');
const mediaQueryTwoImages = window.matchMedia('(max-width: 801px)');
const mediaQueryOneImage = window.matchMedia('(max-width: 696px)');

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

const mediaQueryImagesCount = () => {
  if (mediaQueryOneImage.matches) {
    return 2;
  }
  if (mediaQueryTwoImages.matches) {
    return 3;
  }
  return 4;
};

class RequestDetails extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      contactFormVisible: false,
      previewVisible: false,
      previewImage: '',
      responseFeedback: false,
      hideArrowCount: mediaQueryImagesCount(),
    };

    mediaQueryOneImage.addListener(this.hideArrowCount.bind(this));
    mediaQueryTwoImages.addListener(this.hideArrowCount.bind(this));
  }

  componentWillMount() {
    ReactGA.modalview('/request-details', null, 'Detalhes da Solicitação');
  }

  componentWillUnmount() {
    mediaQueryOneImage.removeListener(this.hideArrowCount);
    mediaQueryTwoImages.removeListener(this.hideArrowCount);
  }

  showContactForm() {
    ReactGA.event({
      category: 'Doador',
      action: 'Click em Quero Ajudar',
    });
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

  hideArrowCount() {
    this.setState({ hideArrowCount: mediaQueryImagesCount() });
  }

  closeModal() {
    this.props.onCancel();
    this.setState({ contactFormVisible: false });
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
          <Img
            src={imageObj.url}
            alt={imageObj.name}
            loader={<Icon type="loading" style={{ fontSize: 40, color: colors.cian_400 }} />}
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
          onFeedback={visible => this.setState({ responseFeedback: visible })}
          request={request}
        />
      );
    }

    const {
      createdAt,
      updatedAt,
      title,
      description,
      organization,
      requiredQuantity,
      reachedQuantity,
      unit,
      images,
    } = request;

    let dateText = `Criado em: ${moment(createdAt).format('LLL')}`;
    if (updatedAt) {
      dateText = `Atualizado em: ${moment(updatedAt).format('LLL')}`;
    }

    return (
      <div className={styles.contentWrapper}>
        <div className={styles.itemDetails}>
          <ItemIndicator className={styles.itemIndicator} request={request} size="lg" />
          <div>
            <h1>{title}</h1>
            <p className={styles.received}>{`Recebidos: ${reachedQuantity} de ${requiredQuantity} ${unit}`}</p>
            <p className={styles.updatedAt}>{dateText}</p>
          </div>
        </div>
        <div className={styles.organizationBox}>
          <div className={styles.organizationBorder}>
            <p className={styles.organization}>{organization.name}</p>
            <p style={{ 'white-space': 'pre-line' }} className={styles.description}>{description}</p>
          </div>
        </div>
        {images.length > 0 &&
          <div className={styles.arrowWrapper}>
            <Arrow
              size={32}
              color={colors.purple_400}
              onClick={() => this.carousel.prev()}
              left
              hidden={images.length < this.state.hideArrowCount}
            />
            <div className={styles.carouselWrapper}>
              <Carousel
                ref={(ref) => { this.carousel = ref; }}
                infinite={false}
                {...carouselSettings}
              >
                {this.renderImages(images)}
              </Carousel>
            </div>
            <Arrow
              size={32}
              color={colors.purple_400}
              onClick={() => this.carousel.next()}
              hidden={images.length < this.state.hideArrowCount}
            />
          </div>
        }
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
        onCancel={() => this.closeModal()}
        closable={!this.state.contactFormVisible}
        wrapClassName={this.state.responseFeedback && styles.modalFixed}
      >
        {this.renderContent()}
      </Modal>
    );
  }
}

export default RequestDetails;
