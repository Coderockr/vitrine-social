import React from 'react';
import cx from 'classnames';
import Icon from '../../components/Icons';
import styles from './styles.module.scss';

class Arrow extends React.PureComponent {
  renderIcon() {
    if (this.props.left) {
      return (
        <div
          className={this.props.over ?
            cx(styles.arrowButton, styles.arrowButtonOver) :
            styles.arrowButton}
        >
          <Icon
            icon="arrow-left-drop-circle-outline"
            size={this.props.size}
            color={this.props.color}
            className={this.props.over ?
              cx(styles.imgBottom, styles.imgOver) :
              styles.imgBottom}
          />
          <Icon
            icon="arrow-left-drop-circle"
            size={this.props.size}
            color={this.props.color}
            className={this.props.over ?
              cx(styles.imgTop, styles.imgOver) :
              styles.imgTop}
          />
        </div>
      );
    }
    return (
      <div
        className={this.props.over ?
          cx(styles.arrowButton, styles.arrowButtonOver) :
          styles.arrowButton}
      >
        <Icon
          icon="arrow-right-drop-circle-outline"
          size={this.props.size}
          color={this.props.color}
          className={this.props.over ?
            cx(styles.imgBottom, styles.imgOver) :
            styles.imgBottom}
        />
        <Icon
          icon="arrow-right-drop-circle"
          size={this.props.size}
          color={this.props.color}
          className={this.props.over ?
            cx(styles.imgTop, styles.imgOver) :
            styles.imgTop}
        />
      </div>
    );
  }

  render() {
    return (
      <div
        role="button"
        tabIndex={0}
        onClick={this.props.onClick}
        onKeyPress={this.props.onClick}
        className={styles.arrowButton}
      >
        {this.renderIcon()}
      </div>
    );
  }
}

export default Arrow;
