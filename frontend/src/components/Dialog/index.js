import React from 'react';
import cx from 'classnames';
import Icon from '../Icons';
import './style.css';

const Dialog = ({ children, active }) => (
  <div className={cx('backdrop', { 'backdrop-active': active })}>
    <div className="dialog-card">
      <a className="close-button">
        <Icon
          icon="close"
          size={32}
          color="#FA578A"
        />
      </a>
      {children}
    </div>
  </div>
);

export default Dialog;
