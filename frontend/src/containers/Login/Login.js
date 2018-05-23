import React from 'react';
import { Layout, Row, Col, Form, Icon, Input } from 'antd';
import styles from './styles.module.scss';
import Header from '../../components/Header';

const FormItem = Form.Item;
const { Content } = Layout;

class Login extends React.Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        return values;
      }
      return err;
    });
  }

  render() {
    const { getFieldDecorator } = this.props.form;

    return (
      <Layout>
        <Header />
        <Content>
          <Row type="flex" className={styles.row} style={{ height: '100vh' }}>
            <Col
              lg={{ span: 8, offset: 8 }}
              md={{ span: 10, offset: 7 }}
              sm={{ span: 12, offset: 6 }}
              xs={{ span: 20, offset: 2 }}
            >
              <h1 className={styles.title}>Login da Organização</h1>
              <Form onSubmit={this.handleSubmit}>
                <FormItem>
                  {getFieldDecorator('userName', {
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
                  <div className={styles.buttonWrapper}>
                    <button type="primary" htmlType="submit" className={styles.button}>
                      LOG IN
                    </button>
                    <a className="login-form-forgot" href="">Esqueci a senha</a>
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
