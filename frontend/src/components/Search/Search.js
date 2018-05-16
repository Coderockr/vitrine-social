import React from 'react';
import { Row, Col } from 'antd';
import Icon from '../Icons';
import styles from './styles.module.scss';

const mediaQuery = window.matchMedia('(min-width: 720px)');

class Search extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      placeholder: mediaQuery.matches ? 'Como você gostaria de ajudar?' : 'Como quer ajudar?',
    };

    mediaQuery.addListener(this.widthChange.bind(this));
  }

  componentWillUnmount() {
    mediaQuery.removeListener(this.widthChange);
  }

  widthChange() {
    this.setState({
      placeholder: mediaQuery.matches ? 'Como você gostaria de ajudar?' : 'Como quer ajudar?',
    });
  }

  render() {
    return (
      <Row className={styles.search}>
        <Col
          xxl={{ span: 10, offset: 7 }}
          xl={{ span: 14, offset: 5 }}
          sm={{ span: 16, offset: 4 }}
          xs={{ span: 20, offset: 2 }}
        >
          <div className={styles.wrapper}>
            <input type="text" placeholder={this.state.placeholder} />
            <button>
              <Icon icon="lupa" size={32} color="#444F60" />
            </button>
          </div>
        </Col>
      </Row>
    );
  }
}

export default Search;
