import React from 'react';
import moment from 'moment';
import { Modal, Carousel } from 'antd';
import './style.css';
import ItemIndicator from '../../components/ItemIndicator';
import Arrow from '../../components/Arrow';

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

class RequestDetails extends React.Component {
  state = {
    visible: this.props.visible,
  }

  render() {
    return (
      <Modal
        visible={this.state.visible}
        footer={null}
        width={800}
        className="modal"
        destroyOnClose
        onCancel={() => this.setState({ visible: false })}
      >
        <div className="contentWrapper">
          <div className="itemDetails">
            <ItemIndicator className="itemIndicator" request={request} size="lg" />
            <div>
              <h1 className="title">{request.item}</h1>
              <p className="updatedAt">
                Atualizado em {
                  moment(request.data).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
                }
              </p>
            </div>
          </div>
          <div className="organizationBox">
            <div className="organizationBorder">
              <p className="organization">{request.organization.name}</p>
              <p className="description">{request.description}</p>
            </div>
          </div>
          <div className="arrowWrapper">
            <Arrow size={32} color="#948CF9" onClick={() => this.carousel.prev()} left />
            <div className="carouselWrapper">
              <Carousel
                ref={(ref) => { this.carousel = ref; }}
                infinite={false}
                slidesToShow={3}
              >
                <div>
                  <img src="assets/images/leitura-infantil.jpg" alt="Leitura Infantil" />
                </div>
                <div>
                  <img src="assets/images/leitura-infantil 2.jpg" alt="Leitura Infantil 2" />
                </div>
                <div>
                  <img src="assets/images/leitura-infantil 3.jpg" alt="Leitura Infantil 3" />
                </div>
                <div>
                  <img src="assets/images/leitura-infantil 2.jpg" alt="Leitura Infantil 2" />
                </div>
                <div>
                  <img src="assets/images/leitura-infantil 3.jpg" alt="Leitura Infantil 3" />
                </div>
              </Carousel>
            </div>
            <Arrow size={32} color="#948CF9" onClick={() => this.carousel.next()} />
          </div>
          <button className="button">QUERO AJUDAR!</button>
        </div>
      </Modal>
    );
  }
}

export default RequestDetails;
