import React from 'react';
import { Layout } from 'antd';
import Header from '../../components/Header';
import Footer from '../../components/Footer';
import ChangePassword from '../../components/ChangePassword';
import styles from './styles.module.scss';

const { Content } = Layout;

const ResetPassword = ({ history, match }) => (
  <Layout className={styles.layout}>
    <Header className={styles.header} />
    <Content className={styles.content}>
      <ChangePassword token={match.params.token} history={history} />
    </Content>
    <Footer className={styles.footer} />
  </Layout>
);

export default ResetPassword;
