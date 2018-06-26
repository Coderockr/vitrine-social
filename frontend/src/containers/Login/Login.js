import React from 'react';
import { Layout, Row, Col, Form, Icon, Input, Button } from 'antd';
import cx from 'classnames';
import Header from '../../components/Header';
import Footer from '../../components/Footer';
import BottomNotification from '../../components/BottomNotification';
import api from '../../utils/api';
import { authorizeUser } from '../../utils/auth';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { Content } = Layout;

class Login extends React.Component {
  state = {
    loading: false,
  }

  componentDidMount() {
    document.title = 'Vitrine Social - Login';
  }

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
    const { history } = this.props;
    this.setState({ loading: true });
    api.post('auth/login', params).then(
      (response) => {
        if (response.data) {
          authorizeUser(response.data);
          history.push(`/organization/${response.data.organization.id}`);
          return BottomNotification({ message: 'Login realizado com sucesso!', success: true });
        }
        this.setState({ loading: false });
        return null;
      }, (error) => {
        this.setState({ loading: false });
        if (!error.response) {
          return BottomNotification({ message: 'Problema de conexão com a API.', success: false });
        } if (error.response.status === 401) {
          return BottomNotification({ message: 'Usuário e/ou senha incorretos.', success: false });
        } if (error.response.data.message) {
          return BottomNotification({ message: error.response.data.message, success: false });
        }
        return null;
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
              <h1>LOGIN DA ORGANIZAÇÃO</h1>
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
                    <Button type="primary" htmlType="submit" className={cx(styles.button, styles.loginButton)} loading={this.state.loading}>
                      LOG IN
                    </Button>
                  </div>
                </FormItem>
              </Form>
            </Col>
          </Row>
        </Content>
        <Footer className={styles.footer} />
      </Layout>
    );
  }
}

const WrappedLoginForm = Form.create()(Login);

export default WrappedLoginForm;
