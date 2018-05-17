import React from 'react';
import { Row, Col, Modal, Avatar, Form, Input } from 'antd';
import cx from 'classnames';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;

class OrganizationProfileEdit extends React.Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        return values;
      }
      return null;
    });
  }

  render() {
    const { getFieldDecorator } = this.props.form;
    const formItemLayout = {
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 14, offset: 5 },
        lg: { span: 16, offset: 4 },
      },
    };

    return (
      <Modal
        visible={this.props.visible}
        footer={null}
        width={800}
        className={styles.modal}
        destroyOnClose
        onCancel={this.props.onCancel}
        closable={false}
        maskClosable={false}
      >
        <Row>
          <Col span={20} offset={2}>
            <h1 className={styles.title}>Editar Perfil da Organização</h1>
            <div className={styles.avatarWrapper}>
              <Avatar
                icon="user"
                style={{ fontSize: 140, backgroundColor: '#FF974B' }}
              />
            </div>
            <Form onSubmit={this.handleSubmit}>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('name', {
                  rules: [{ required: true, message: 'Preencha o nome da organização' }],
                })(
                  <Input size="large" placeholder="Nome da Organização" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('email', {
                  rules: [{
                    type: 'email', message: 'E-mail inválido',
                  }, {
                    required: true, message: 'Preencha o e-mail',
                  }],
                })(
                  <Input size="large" placeholder="E-mail" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('phone', {
                  rules: [{
                    required: true, message: 'Preencha o telefone',
                  }, {
                    pattern: /^1\d\d(\d\d)?$|^0800 ?\d{3} ?\d{4}$|^(\(0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d\) ?|0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d[ .-]?)?(9|9[ .-])?[2-9]\d{3}[ .-]?\d{4}$/gm, message: 'Telefone Inválido',
                  }],
                })(
                  <Input size="large" placeholder="Telefone" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('address', {
                  rules: [{ required: true, message: 'Preencha o endereço' }],
                })(
                  <Input size="large" placeholder="Endereço" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <TextArea rows={5} placeholder="Sobre a Organização" />
              </FormItem>
              <FormItem>
                <div className={styles.buttonWrapper}>
                  <button className={cx(styles.button, styles.saveButton)} htmlType="submit">
                    SALVAR
                  </button>
                  <button
                    className={cx(styles.button, styles.cancelButton)}
                    onClick={this.props.onCancel}
                  >
                    CANCELAR
                  </button>
                </div>
              </FormItem>
            </Form>
          </Col>
        </Row>
      </Modal>
    );
  }
}

const WrappedEditProfileForm = Form.create()(OrganizationProfileEdit);

export default WrappedEditProfileForm;
