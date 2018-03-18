import React, { Component } from 'react';
import cx from 'classnames';
import Icon from '../Icons';
import './style.css';


class Dialog extends Component {
  constructor(props) {
    super(props);

    this.state = {
      active: this.props.active,
    };
  }

  closeModal() {
    this.setState({ active: false });
  }

  render() {
    const { children } = this.props;

    return (
      <div className={cx('backdrop', { 'backdrop-active': this.state.active })}>
        <div className="dialog-card">
          <a
            className="close-button"
            onClick={() => this.closeModal()}
            onKeyUp={() => this.closeModal()}
            role="button"
            tabIndex={0}
          >
            <Icon icon="close" size={32} color="#FA578A" />
          </a>
          {children}
        </div>
      </div>
    );
  }
}

export default Dialog;
