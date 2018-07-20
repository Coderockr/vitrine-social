import React from 'react';
import { Layout, Row, Col, Form, Input, Icon, Button } from 'antd';
import cx from 'classnames';
import api from '../../utils/api';
import Header from '../../components/Header';
import Footer from '../../components/Footer';
import BottomNotification from '../../components/BottomNotification';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { Content } = Layout;

class ForgotPassword extends React.Component {
  state = {
    loading: false,
  }

  componentDidMount() {
    document.title = 'Vitrine Social - Esqueci a Senha';
  }

  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        this.forgotPassword(values);
      }
      return err;
    });
  }

  forgotPassword(params) {
    this.setState({ loading: true });
    api.post('auth/forgot-password', params).then(
      () => {
        this.setState({ loading: false });
        return BottomNotification({ message: 'Verifique seu e-mail e siga as instruções para resetar a sua senha!', success: true });
      }, (error) => {
        this.setState({ loading: false });
        if (!error.response) {
          return BottomNotification({ message: 'Problema de conexão com a API.', success: false });
        } if (error.response.status === 404) {
          return BottomNotification({ message: 'Verifique seu e-mail e siga as instruções para resetar a sua senha!', success: false });
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
              <h1>RECUPERAR A SENHA</h1>
              <Form>
                <FormItem>
                  {getFieldDecorator('email', {
                    rules: [{
                      type: 'email', message: 'E-mail inválido',
                    }, {
                        required: true, message: 'Informe seu email!',
                    }],
                  })(
                    <Input prefix={<Icon type="user" />} placeholder="Usuário" size="large" />,
                  )}
                </FormItem>
                <FormItem>
                  <div className={styles.buttonWrapper}>
                    <Button
                      className={cx(styles.button, styles.sendButton)}
                      loading={this.state.loading}
                      onClick={this.handleSubmit}
                    >
                      ENVIAR
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

const WrappedForgotPasswordForm = Form.create()(ForgotPassword);

export default WrappedForgotPasswordForm;
