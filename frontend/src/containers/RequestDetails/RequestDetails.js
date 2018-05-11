import React from 'react';
import moment from 'moment';
import { Modal } from 'antd';
import './style.css';
import ItemIndicator from '../../components/ItemIndicator';

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
          <button className="button">QUERO AJUDAR!</button>
        </div>
      </Modal>
    );
  }
}

export default RequestDetails;
