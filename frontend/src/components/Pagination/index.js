import React from 'react';
import {
  Nav,
  NavItem,
  NavGroup,
} from 're-bulma';

import './style.css';

const Pagination = () => (
  <Nav className="pagination">
    <NavGroup className="wrapper">
      <NavItem>
        <a href="#">1</a>
      </NavItem>
      <NavItem>
        <a href="#">2</a>
      </NavItem>
      <NavItem>
        <a href="#">3</a>
      </NavItem>
      <NavItem>
        <a href="#">4</a>
      </NavItem>
      <NavItem>
        <a href="#">{'>'}</a>
      </NavItem>
    </NavGroup>
  </Nav>
);

export default Pagination;
