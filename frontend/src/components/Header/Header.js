import React from 'react';
import { Row, Col, Modal, Layout, Menu } from 'antd';
import { Link } from 'react-router-dom';
import cx from 'classnames';
import Icon from '../../components/Icons';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

const mediaQuery = window.matchMedia('(max-width: 685px)');

class Header extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      collapsed: mediaQuery.matches,
      visible: false,
    };

    mediaQuery.addListener(this.widthChange.bind(this));
  }

  componentWillUnmount() {
    mediaQuery.removeListener(this.widthChange);
  }

  widthChange() {
    this.setState({
      collapsed: mediaQuery.matches,
      visible: mediaQuery.matches ? this.state.visible : false,
    });
  }

  showMenu() {
    this.setState({
      visible: true,
    });
  }

  hideMenu() {
    this.setState({
      visible: false,
    });
  }

  renderMenuButton() {
    return (
      <button className={styles.button} onClick={() => this.showMenu()}>
        <Icon
          icon="menu"
          size={40}
          color={colors.white}
        />
      </button>
    );
  }

  renderMenu(collapsed) {
    return (
      <Menu
        className={collapsed ? styles.menuModal : styles.appHeader}
        mode={collapsed ? 'inline' : 'horizontal'}
        defaultSelectedKeys={['1']}
        theme={collapsed ? 'light' : 'dark'}
        selectedKeys={[window.location.pathname]}
      >
        <Menu.Item key="/">
          <Link to="/">Sobre o Projeto</Link>
        </Menu.Item>
        <Menu.Item key="/organization">
          <Link to="/organization">Quero Participar</Link>
        </Menu.Item>
        <Menu.Item key="/contact">
          <Link to="/contact">Contato</Link>
        </Menu.Item>
        <Menu.Item key="/login">
          <Link to="/login">Login</Link>
        </Menu.Item>
      </Menu>
    );
  }

  render() {
    return (
      <Layout.Header
        className={cx(styles.appHeader, this.props.className)}
      >
        <Row>
          <Col
            xxl={{ span: 14, offset: 5 }}
            xl={{ span: 20, offset: 2 }}
            lg={{ span: 22, offset: 1 }}
            md={{ span: 24, offset: 0 }}
          >
            <div className={styles.logo} />
            {this.state.collapsed ? this.renderMenuButton() : this.renderMenu()}
          </Col>
        </Row>
        <Modal
          visible={this.state.visible}
          footer={null}
          onCancel={() => this.hideMenu()}
          bodyStyle={{ height: '100vh', backgroundColor: 'rgba(255,255,255,.8)' }}
          className={cx(styles.modal, 'menuModal')}
        >
          {this.renderMenu(true)}
        </Modal>
      </Layout.Header>
    );
  }
}

export default Header;
