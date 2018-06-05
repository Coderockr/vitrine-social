import React from 'react';
import { Layout, Row, Col, Form, Icon, Input } from 'antd';
import Header from '../../components/Header';
import BottomNotification from '../../components/BottomNotification';
import api from '../../utils/api';
import { authorizeUser } from '../../utils/auth';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { Content } = Layout;

class Login extends React.Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        this.loginUser(values);
      }
      return err;
    });
  }

  loginUser(params) {
    api.post('auth/login', params).then(
      (response) => {
        if (response.data) {
          authorizeUser(response.data);
        }
        return null;
      }, (error) => {
        if (!error.response) {
          BottomNotification('Problema de conexão com a API.');
        } else if (error.response.status === 401) {
          BottomNotification('Usuário e/ou senha incorretos.');
        } else if (error.response.data.message) {
          BottomNotification(error.response.data.message);
        }
      },
    );
  }

  render() {
    const { getFieldDecorator } = this.props.form;

    return (
      <Layout className={styles.layout}>
        <Header className={styles.header} />
        <Content className={styles.content}>
          <Row className={styles.row}>
            <Col
              xxl={{ span: 6, offset: 9 }}
              lg={{ span: 8, offset: 8 }}
              md={{ span: 10, offset: 7 }}
              sm={{ span: 12, offset: 6 }}
              xs={{ span: 20, offset: 2 }}
            >
              <h1>Login da Organização</h1>
              <Form onSubmit={this.handleSubmit}>
                <FormItem>
                  {getFieldDecorator('email', {
                    rules: [{ required: true, message: 'Informe seu usuário!' }],
                  })(
                    <Input prefix={<Icon type="user" />} placeholder="Usuário" size="large" />,
                  )}
                </FormItem>
                <FormItem>
                  {getFieldDecorator('password', {
                    rules: [{ required: true, message: 'Informe sua senha!' }],
                  })(
                    <Input prefix={<Icon type="lock" />} type="password" placeholder="Senha" size="large" />,
                  )}
                </FormItem>
                <FormItem>
                  <a
                    className={styles.forgotPassword}
                    href=""
                  >
                    Esqueci a senha
                  </a>
                </FormItem>
                <FormItem>
                  <div className={styles.buttonWrapper}>
                    <button type="primary" htmlType="submit" className={styles.button}>
                      LOG IN
                    </button>
                  </div>
                </FormItem>
              </Form>
            </Col>
          </Row>
        </Content>
      </Layout>
    );
  }
}

const WrappedLoginForm = Form.create()(Login);

export default WrappedLoginForm;
