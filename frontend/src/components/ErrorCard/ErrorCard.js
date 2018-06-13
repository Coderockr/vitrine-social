import React from 'react';
import styles from './styles.module.scss';

const ErrorCard = ({ text }) => (
  <div className={styles.errorWrapper}>
    <p className={styles.errorText}>{text}</p>
  </div>
);

export default ErrorCard;
