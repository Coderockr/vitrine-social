import React from 'react'

import map from './map.json'

const defaultProps = {
  size: 25,
  className: false,
  styles: {}
}

const Icon = ({ icon, size, styles, className, color, ...props }) => (
  <svg
    {...props}
    width={size}
    height={size}
    viewBox={map[icon] ? map[icon].viewBox : '0 0 50 50'}
    style={{ ...styles, width: size, height: size }}
    fill={color}
    className={className ? className + ' ' : ''}
    fill={color}
  >
    <path d={map[icon] ? map[icon].paths : ''} />
  </svg>
)

Icon.defaultProps = defaultProps;

export default Icon;
