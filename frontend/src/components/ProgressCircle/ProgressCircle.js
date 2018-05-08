import React from 'react';
import './style.css';

const ProgressCircle = ({ progress, size }) => (
  <div className={`progress-circle progress-${progress} ${size}`} />
);

export default ProgressCircle;
