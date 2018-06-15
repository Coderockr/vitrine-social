import React from 'react';
import { Row, Col, Modal, Form, Icon, Input } from 'antd';
import cx from 'classnames';
import BottomNotification from '../../components/BottomNotification';
import api, { headers } from '../../utils/api';
import styles from './styles.module.scss';

const FormItem = Form.Item;

class ChangePassword extends React.Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (values.newPassword !== values.confirmPassword) {
        return BottomNotification({ message: 'Senhas não conferem! Verifique e tente novamente!', success: false });
      }
      if (!err) {
        this.changePassword(values.newPassword);
      }
      return err;
    });
  }

  changePassword(newPassword) {
    const { history } = this.props;
    api.defaults.headers = { ...headers, Authorization: this.props.token };
    api.post('auth/reset-password', { newPassword }).then(
      () => {
        history.push('/login');
        BottomNotification({ message: 'Senha alterada com sucesso!', success: true });
      }, () => {
        BottomNotification({ message: 'Não foi possível alterar a sua senha!', success: false });
      },
    );
  }

  renderModal() {
    return (
      <Modal
        visible={this.props.visible}
        footer={null}
        width={800}
        className={styles.modal}
        destroyOnClose
        onCancel={this.props.onCancel}
        // wrapClassName={this.state.responseFeedback && styles.modalFixed}
      >
        {this.renderForm()}
      </Modal>
    );
  }

  renderForm() {
    const { getFieldDecorator } = this.props.form;

    return (
      <Row className={this.props.modal ? styles.rowModal : styles.row}>
        <Col
          xxl={{ span: 6, offset: 9 }}
          lg={{ span: 8, offset: 8 }}
          md={{ span: 10, offset: 7 }}
          sm={{ span: 12, offset: 6 }}
          xs={{ span: 20, offset: 2 }}
        >
          <h1>Alterar a Senha</h1>
          <Form onSubmit={this.handleSubmit}>
            <FormItem>
              {getFieldDecorator('newPassword', {
                rules: [{ required: true, message: 'Informe sua nova senha!' }],
              })(
                <Input prefix={<Icon type="lock" />} type="password" placeholder="Nova senha" size="large" />,
              )}
            </FormItem>
            <FormItem>
              {getFieldDecorator('confirmPassword', {
                rules: [{ required: true, message: 'Informe sua nova senha novamente!' }],
              })(
                <Input prefix={<Icon type="lock" />} type="password" placeholder="Confirme a senha" size="large" />,
              )}
            </FormItem>
            <FormItem>
              <div className={styles.buttonWrapper}>
                <button type="primary" htmlType="submit" className={cx(styles.button, styles.sendButton)}>
                  ENVIAR
                </button>
              </div>
            </FormItem>
          </Form>
        </Col>
      </Row>
    );
  }

  render() {
    if (this.props.modal) {
      return this.renderModal();
    }
    return this.renderForm();
  }
}

const WrappedChangePasswordForm = Form.create()(ChangePassword);

export default WrappedChangePasswordForm;
