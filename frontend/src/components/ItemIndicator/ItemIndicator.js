import React from 'react';
import cx from 'classnames';
import Icon from '../Icons';
import ProgressCircle from '../ProgressCircle';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

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
        icon={request.category}
        size={size === 'lg' ? 85 : 60}
        color={colors.ambar_400}
      />
      <div className={cx(styles.progressCircleContainer, styles[size])}>
        <ProgressCircle progress={60} size={size} />
        <div className={cx(styles.lasteQtd, styles[size])}>
          <p>Faltam 4</p>
        </div>
      </div>
    </div>
  </div>
);

export default ItemIndicator;
