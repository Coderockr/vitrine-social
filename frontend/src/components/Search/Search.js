import React from 'react';
import { Row, Col } from 'antd';
import Icon from '../Icons';
import colors from '../../utils/styles/colors';
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

  handleSearch = (e) => {
    e.preventDefault();
    this.props.search(this.input.value);
  }

  render() {
    return (
      <Row className={styles.search}>
        <Col
          xxl={{ span: 10, offset: 7 }}
          xl={{ span: 14, offset: 5 }}
          sm={{ span: 16, offset: 4 }}
          xs={{ span: 22, offset: 1 }}
        >
          <form className={styles.wrapper} onSubmit={this.handleSearch}>
            <input
              type="text"
              placeholder={this.state.placeholder}
              ref={(ref) => { this.input = ref; }}
              defaultValue={this.props.text}
            />
            <button>
              <Icon icon="lupa" size={32} color={colors.grey_400} />
            </button>
          </form>
        </Col>
      </Row>
    );
  }
}

export default Search;
