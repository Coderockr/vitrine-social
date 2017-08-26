import React from 'react'

import { map } from './map.json'

const Icon = (icon) => {
  <svg viewBox={map[icon].viewBox}>
    <path d={map[icon].paths} />
  </svg>
}
