import React from 'react';
import { Row, Col, Modal, Form, Icon, Input } from 'antd';
import cx from 'classnames';
import BottomNotification from '../../components/BottomNotification';
import ResponseFeedback from '../ResponseFeedback';
import api, { headers } from '../../utils/api';
import styles from './styles.module.scss';

const FormItem = Form.Item;

class ChangePassword extends React.Component {
  state = {
    responseFeedback: '',
    responseFeedbackMessage: '',
  }

  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (values.newPassword !== values.confirmPassword) {
        return BottomNotification({ message: 'Senhas não conferem! Verifique e tente novamente!', success: false });
      }
      if (!err) {
        if (this.props.user) {
          return this.changePassword(values.newPassword, values.currentPassword);
        }
        return this.resetPassword(values.newPassword);
      }
      return err;
    });
  }

  changePassword(newPassword, currentPassword) {
    api.post('auth/update-password', { newPassword, currentPassword }).then(
      () => {
        this.setState({
          responseFeedback: 'success',
          responseFeedbackMessage: 'Senha alterada com sucesso!',
        });
      }, () => {
        this.setState({
          responseFeedback: 'error',
          responseFeedbackMessage: 'Não foi possível alterar a sua senha! Verifique sua senha atual e tente novamente',
        });
      },
    );
  }

  resetPassword(newPassword) {
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

  closeModal() {
    this.props.onCancel();
    setTimeout(() => {
      this.props.form.resetFields();
      this.setState({
        responseFeedback: '',
        responseFeedbackMessage: '',
      });
    }, 100);
  }

  renderModal() {
    return (
      <Modal
        visible={this.props.visible}
        footer={null}
        width={800}
        className={styles.modal}
        destroyOnClose
        onCancel={() => this.closeModal()}
        wrapClassName={this.state.responseFeedback && styles.modalFixed}
      >
        {this.renderForm()}
        <ResponseFeedback
          small
          type={this.state.responseFeedback}
          message={this.state.responseFeedbackMessage}
          onClick={this.state.responseFeedback === 'error' ?
            () => this.setState({ responseFeedback: '', responseFeedbackMessage: '' }) :
            () => this.closeModal()
          }
        />
      </Modal>
    );
  }

  renderForm() {
    const { getFieldDecorator } = this.props.form;

    return (
      <Row className={cx(
        this.props.modal ? styles.rowModal : styles.row,
        { [styles.blurRow]: this.state.responseFeedback },
      )}
      >
        <Col
          xxl={{ span: 6, offset: 9 }}
          lg={{ span: 8, offset: 8 }}
          md={{ span: 10, offset: 7 }}
          sm={{ span: 12, offset: 6 }}
          xs={{ span: 20, offset: 2 }}
        >
          <h1>Alterar a Senha</h1>
          <Form onSubmit={this.handleSubmit}>
            {this.props.user && (
              <FormItem>
                {getFieldDecorator('currentPassword', {
                  rules: [{ required: true, message: 'Informe sua senha atual!' }],
                })(
                  <Input prefix={<Icon type="lock" />} type="password" placeholder="Senha atual" size="large" />,
                )}
              </FormItem>
            )}
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
