import React from 'react';
import Icon from '../Icons';
import ProgressCircle from '../ProgressCircle';

import './style.css';

const ItemIndicator = ({ request, size, className }) => (
  <div className={`requestCircle ${size} ${className}`}>
    <div className={`requestIcon ${size}`}>
      <Icon icon={request.category} size={size === 'lg' ? 85 : 60} color="#FF974B" />
      <div className={`progress-circle-container ${size}`}>
        <ProgressCircle progress={60} size={size} />
        <div className={`laste-qtd ${size}`}>
          <p>
            Faltam 4
          </p>
        </div>
      </div>
    </div>
  </div>
);

export default ItemIndicator;
