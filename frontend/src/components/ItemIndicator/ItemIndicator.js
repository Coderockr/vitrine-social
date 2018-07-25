import React from 'react';
import cx from 'classnames';
import Icon from '../Icons';
import ProgressCircle from '../ProgressCircle';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

const calculateProgress = (receivedQty, requiredQty) => (
  Math.ceil((receivedQty / requiredQty) * 100)
);

const ItemIndicator = ({ request, size, className }) => (
  <div className={
    cx(
      styles.requestCircle,
      styles[size],
      className,
    )}
  >
    <div className={cx(styles.requestIcon, styles[size])}>
      <Icon
        icon={request.category.slug}
        size={size === 'lg' ? 85 : 60}
        color={colors.ambar_400}
      />
    </div>
    <div className={cx(styles.progressCircleContainer, styles[size])}>
      <ProgressCircle
        progress={calculateProgress(request.reachedQuantity, request.requiredQuantity)}
        size={size}
      />
      <div className={cx(styles.lasteQtd, styles[size])}>
        <p>Faltam {request.requiredQuantity - request.reachedQuantity} {request.unit}</p>
      </div>
    </div>
  </div>
);

export default ItemIndicator;
