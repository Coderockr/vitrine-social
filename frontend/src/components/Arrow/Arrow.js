import React from 'react';
import cx from 'classnames';
import Icon from '../../components/Icons';
import styles from './styles.module.scss';

const getIcon = left => (left ? 'left' : 'right');

const Arrow = ({
  size, color, onClick, hidden, over, left,
}) => (
  <div
    role="button"
    tabIndex={0}
    onClick={onClick}
    onKeyPress={onClick}
    className={styles.arrowButton}
    hidden={hidden}
  >
    <div
      className={cx(styles.arrowButton, { [styles.arrowButtonOver]: over })}
    >
      <Icon
        icon={`arrow-${getIcon(left)}-drop-circle-outline`}
        size={size}
        color={color}
        className={cx(styles.imgBottom, { [styles.imgOver]: over })}
      />
      <Icon
        icon={`arrow-${getIcon(left)}-drop-circle`}
        size={size}
        color={color}
        className={cx(styles.imgTop, { [styles.imgOver]: over })}
      />
    </div>
  </div>
);

export default Arrow;
