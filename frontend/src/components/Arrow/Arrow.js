import React from 'react';
import Icon from '../../components/Icons';
import './style.css';

class Arrow extends React.PureComponent {
  state = {}

  renderIcon() {
    if (this.props.left) {
      return (
        <Icon icon="arrow-left-drop-circle-outline" size={this.props.size} color={this.props.color} />
      );
    }
    return (
      <Icon icon="arrow-right-drop-circle-outline" size={this.props.size} color={this.props.color} />
    );
  }

  render() {
    return (
      <div
        role="button"
        tabIndex={0}
        onClick={this.props.onClick}
        onKeyPress={this.props.onClick}
        className="arrowButton"
      >
        {this.renderIcon()}
      </div>
    );
  }
}

export default Arrow;
