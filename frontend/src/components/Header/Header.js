import React from 'react';
import { Row, Col, Modal, Layout, Menu, Dropdown, Avatar } from 'antd';
import { Link } from 'react-router-dom';
import cx from 'classnames';
import { getUser, deauthorizeUser } from '../../utils/auth';
import Icon from '../../components/Icons';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

const mediaQuery = window.matchMedia('(max-width: 685px)');

const userMenu = user => (
  <Menu>
    <Menu.Item key="0" className={styles.userTitleItem}>
      <p>{user.name}</p>
    </Menu.Item>
    <Menu.Divider />
    <Menu.Item key="1" className={styles.userMenuItem}>
      <Link to={`/entidade/${user.id}`}>Meu Perfil</Link>
    </Menu.Item>
    <Menu.Item key="2" className={styles.userMenuItem}>
      <Link onClick={() => deauthorizeUser()} to="/login">Log Out</Link>
    </Menu.Item>
  </Menu>
);

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
      <button className={styles.menuButton} onClick={() => this.showMenu()}>
        <Icon
          icon="menu"
          size={40}
          color={colors.white}
        />
      </button>
    );
  }

  renderUserMenuItem(collapsed) {
    const user = getUser();
    if (!user) {
      return (
        <Menu.Item key="/login">
          <Link to="/login">Login</Link>
        </Menu.Item>
      );
    }

    if (collapsed) {
      return null;
    }

    return (
      <Menu.Item key={`/organization/${user.id}`}>
        <Dropdown overlay={userMenu(user)} placement="bottomCenter">
          <Avatar
            icon="user"
            style={{
              fontSize: 33,
              color: colors.white,
              backgroundColor: colors.ambar_200,
              textShadow: '2px 1px 1px #FF974A',
            }}
            src={user.logo}
            size="small"
          />
        </Dropdown>
      </Menu.Item>
    );
  }

  renderMenu(collapsed) {
    const user = getUser();
    return (
      <Menu
        className={collapsed ? styles.menuModal : styles.appHeader}
        mode={collapsed ? 'inline' : 'horizontal'}
        defaultSelectedKeys={['1']}
        theme={collapsed ? 'light' : 'dark'}
        selectedKeys={[window.location.pathname]}
      >
        <Menu.Item key="/">
          <Link to="/">Início</Link>
        </Menu.Item>
        <Menu.Item key="/sobre">
          <Link to="/sobre">Sobre o Projeto</Link>
        </Menu.Item>
        <Menu.Item key="/contato">
          <Link to="/contato">Contato</Link>
        </Menu.Item>
        {collapsed && user &&
          <Menu.Item>
            <Link to={`/entidade/${user.id}`}>Meu Perfil</Link>
          </Menu.Item>
        }
        {collapsed && user &&
          <Menu.Item>
            <Link onClick={() => deauthorizeUser()} to="/login">Log Out</Link>
          </Menu.Item>
        }
        {this.renderUserMenuItem(collapsed)}
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
            <Link to="/">
              <img className={styles.logo} src={`${process.env.REACT_APP_HOST}assets/images/vitrinesocial.svg`} alt="logo" />
            </Link>
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
