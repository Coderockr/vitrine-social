import React from 'react';
import { Row, Col } from 'antd';
import Icon from '../Icons';

import './style.css';

const Search = () => {
  const screenWidth = window.innerWidth;
  return (
    <Row className="search">
      <Col md={{ span: 14, offset: 5 }} sm={{ span: 16, offset: 4 }} xs={{ span: 20, offset: 2 }}>
        <div className="col">
          <input type="text" placeholder={screenWidth < 720 ? 'Como quer ajudar?' : 'Como você gostaria de ajudar?'} />
          <button className="searchButton">
            <Icon icon="lupa" size={32} color="#444F60" />
          </button>
        </div>
      </Col>
    </Row>
  );
};

export default Search;
