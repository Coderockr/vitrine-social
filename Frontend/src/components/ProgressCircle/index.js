import React from 'react';
import './style.css';

const ProgressCircle = ({ progress }) => (
  <div className={`progress-circle progress-${progress}`} />
)

export default ProgressCircle
