import React from 'react';
import moment from 'moment';
import './style.css';
import Dialog from '../../components/Dialog';
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

const RequestDetails = () => (
  <Dialog active>
    <div className="itemDetails">
      <ItemIndicator className="itemIndicator" request={request} size="lg" />
      <h1 className="title">{request.item}</h1>
    </div>
    <p className="updatedAt">
      Atualizado em {
        moment(request.data).format('DD / MMMM / YYYY').replace(/(\/)/g, 'de')
      }
    </p>
    <div className="organizationBox">
      <div className="organizationBorder">
        <p className="organization">{request.organization.name}</p>
        <p className="description">{request.description}</p>
      </div>
    </div>
    <button className="button">QUERO AJUDAR!</button>
  </Dialog>
);

export default RequestDetails;
