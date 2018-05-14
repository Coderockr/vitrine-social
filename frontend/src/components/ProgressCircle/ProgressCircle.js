import React from 'react';
import cx from 'classnames';
import styles from './styles.module.scss';

const ProgressCircle = ({ progress, size }) => (
  <div
    className={cx(
      styles.progress_circle,
      styles[`progress_${progress}`],
      styles[`${size}`],
    )}
  />
);

export default ProgressCircle;
